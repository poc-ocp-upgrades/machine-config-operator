package daemon

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"path/filepath"
	"reflect"
	"strconv"
	"syscall"
	"time"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	"github.com/coreos/ignition/config/validate"
	"github.com/golang/glog"
	"github.com/google/renameio"
	drain "github.com/openshift/kubernetes-drain"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"github.com/openshift/machine-config-operator/pkg/daemon/constants"
	errors "github.com/pkg/errors"
	"github.com/vincent-petithory/dataurl"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	defaultDirectoryPermissions	os.FileMode	= 0755
	defaultFilePermissions		os.FileMode	= 0644
	coreUserName					= "core"
	coreUserSSHPath					= "/home/core/.ssh/"
)

func writeFileAtomicallyWithDefaults(fpath string, b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return writeFileAtomically(fpath, b, defaultDirectoryPermissions, defaultFilePermissions, -1, -1)
}
func writeFileAtomically(fpath string, b []byte, dirMode, fileMode os.FileMode, uid, gid int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := os.MkdirAll(filepath.Dir(fpath), dirMode); err != nil {
		return fmt.Errorf("failed to create directory %q: %v", filepath.Dir(fpath), err)
	}
	t, err := renameio.TempFile("", fpath)
	if err != nil {
		return err
	}
	defer t.Cleanup()
	if err := t.Chmod(fileMode); err != nil {
		return err
	}
	w := bufio.NewWriter(t)
	if _, err := w.Write(b); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}
	if uid != -1 && gid != -1 {
		if err := t.Chown(uid, gid); err != nil {
			return err
		}
	}
	return t.CloseAtomicallyReplace()
}
func (dn *Daemon) writePendingState(desiredConfig *mcfgv1.MachineConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := &pendingConfigState{PendingConfig: desiredConfig.GetName(), BootID: dn.bootID}
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return writeFileAtomicallyWithDefaults(pathStateJSON, b)
}
func getNodeRef(node *corev1.Node) *corev1.ObjectReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &corev1.ObjectReference{Kind: "Node", Name: node.GetName(), UID: types.UID(node.GetUID())}
}
func (dn *Daemon) updateOSAndReboot(newConfig *mcfgv1.MachineConfig) (retErr error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := dn.updateOS(newConfig); err != nil {
		return err
	}
	if err := dn.writePendingState(newConfig); err != nil {
		return errors.Wrapf(err, "writing pending state")
	}
	defer func() {
		if retErr != nil {
			if err := os.Remove(pathStateJSON); err != nil {
				retErr = errors.Wrapf(retErr, "error removing pending config file %v", err)
				return
			}
		}
	}()
	if dn.recorder != nil {
		dn.recorder.Eventf(getNodeRef(dn.node), corev1.EventTypeNormal, "PendingConfig", fmt.Sprintf("Written pending config %s", newConfig.GetName()))
	}
	if dn.onceFrom == "" {
		glog.Info("Update prepared; draining the node")
		dn.recorder.Eventf(getNodeRef(dn.node), corev1.EventTypeNormal, "Drain", "Draining node to update config.")
		backoff := wait.Backoff{Steps: 5, Duration: 10 * time.Second, Factor: 2}
		var lastErr error
		if err := wait.ExponentialBackoff(backoff, func() (bool, error) {
			err := drain.Drain(dn.kubeClient, []*corev1.Node{dn.node}, &drain.DrainOptions{DeleteLocalData: true, Force: true, GracePeriodSeconds: 600, IgnoreDaemonsets: true})
			if err == nil {
				return true, nil
			}
			lastErr = err
			glog.Infof("Draining failed with: %v, retrying", err)
			return false, nil
		}); err != nil {
			if err == wait.ErrWaitTimeout {
				return errors.Wrapf(lastErr, "failed to drain node (%d tries): %v", backoff.Steps, err)
			}
			return errors.Wrap(err, "failed to drain node")
		}
		glog.Info("Node successfully drained")
	}
	return dn.reboot(fmt.Sprintf("Node will reboot into config %v", newConfig.GetName()), defaultRebootTimeout, exec.Command(defaultRebootCommand))
}
func (dn *Daemon) catchIgnoreSIGTERM() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if dn.installedSigterm {
		return
	}
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM)
	dn.installedSigterm = true
	go func() {
		for {
			<-termChan
			glog.Info("Got SIGTERM, but actively updating")
		}
	}()
}

var errUnreconcilable = errors.New("unreconcilable")

func (dn *Daemon) update(oldConfig, newConfig *mcfgv1.MachineConfig) (retErr error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if dn.nodeWriter != nil {
		state, err := getNodeAnnotationExt(dn.node, constants.MachineConfigDaemonStateAnnotationKey, true)
		if err != nil {
			return err
		}
		if state != constants.MachineConfigDaemonStateDegraded && state != constants.MachineConfigDaemonStateUnreconcilable {
			if err := dn.nodeWriter.SetWorking(dn.kubeClient.CoreV1().Nodes(), dn.nodeLister, dn.name); err != nil {
				return errors.Wrap(err, "error setting node's state to Working")
			}
		}
	}
	dn.catchIgnoreSIGTERM()
	defer func() {
		if retErr != nil {
			dn.cancelSIGTERM()
		}
	}()
	oldConfigName := oldConfig.GetName()
	newConfigName := newConfig.GetName()
	glog.Infof("Checking reconcilable for config %v to %v", oldConfigName, newConfigName)
	reconcilableError := dn.reconcilable(oldConfig, newConfig)
	if reconcilableError != nil {
		wrappedErr := fmt.Errorf("can't reconcile config %s with %s: %v", oldConfigName, newConfigName, reconcilableError)
		if dn.recorder != nil {
			mcRef := &corev1.ObjectReference{Kind: "MachineConfig", Name: newConfig.GetName(), UID: newConfig.GetUID()}
			dn.recorder.Eventf(mcRef, corev1.EventTypeWarning, "FailedToReconcile", wrappedErr.Error())
		}
		dn.logSystem(wrappedErr.Error())
		return errors.Wrapf(errUnreconcilable, "%v", wrappedErr)
	}
	if err := dn.updateFiles(oldConfig, newConfig); err != nil {
		return err
	}
	defer func() {
		if retErr != nil {
			if err := dn.updateFiles(newConfig, oldConfig); err != nil {
				retErr = errors.Wrapf(retErr, "error rolling back files writes %v", err)
				return
			}
		}
	}()
	if err := dn.updateSSHKeys(newConfig.Spec.Config.Passwd.Users); err != nil {
		return err
	}
	defer func() {
		if retErr != nil {
			if err := dn.updateSSHKeys(oldConfig.Spec.Config.Passwd.Users); err != nil {
				retErr = errors.Wrapf(retErr, "error rolling back SSH keys updates %v", err)
				return
			}
		}
	}()
	return dn.updateOSAndReboot(newConfig)
}
func (dn *Daemon) reconcilable(oldConfig, newConfig *mcfgv1.MachineConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	glog.Info("Checking if configs are reconcilable")
	if oldConfig.Kind == "" && dn.onceFrom != "" {
		glog.Info("Missing kind in old config. Assuming no prior state.")
		return nil
	}
	oldIgn := oldConfig.Spec.Config
	newIgn := newConfig.Spec.Config
	rpt := validate.ValidateWithoutSource(reflect.ValueOf(newIgn))
	if rpt.IsFatal() {
		return errors.Errorf("invalid Ignition config found: %v", rpt)
	}
	if oldIgn.Ignition.Version != newIgn.Ignition.Version {
		return fmt.Errorf("ignition version mismatch between old and new config: old: %s new: %s", oldIgn.Ignition.Version, newIgn.Ignition.Version)
	}
	if !reflect.DeepEqual(oldIgn.Networkd, newIgn.Networkd) {
		return errors.New("ignition networkd section contains changes")
	}
	if !reflect.DeepEqual(oldIgn.Passwd, newIgn.Passwd) {
		if !reflect.DeepEqual(oldIgn.Passwd.Groups, newIgn.Passwd.Groups) {
			return errors.New("ignition Passwd Groups section contains changes")
		}
		if !reflect.DeepEqual(oldIgn.Passwd.Users, newIgn.Passwd.Users) {
			if len(oldIgn.Passwd.Users) >= 0 && len(newIgn.Passwd.Users) >= 1 {
				for _, user := range newIgn.Passwd.Users {
					if user.Name != coreUserName {
						return errors.New("ignition passwd user section contains unsupported changes: non-core user")
					}
				}
				glog.Infof("user data to be verified before ssh update: %v", newIgn.Passwd.Users[len(newIgn.Passwd.Users)-1])
				if err := verifyUserFields(newIgn.Passwd.Users[len(newIgn.Passwd.Users)-1]); err != nil {
					return err
				}
			}
		}
	}
	if !reflect.DeepEqual(oldIgn.Storage.Disks, newIgn.Storage.Disks) {
		return errors.New("ignition disks section contains changes")
	}
	if !reflect.DeepEqual(oldIgn.Storage.Filesystems, newIgn.Storage.Filesystems) {
		return errors.New("ignition filesystems section contains changes")
	}
	if !reflect.DeepEqual(oldIgn.Storage.Raid, newIgn.Storage.Raid) {
		return errors.New("ignition raid section contains changes")
	}
	if !reflect.DeepEqual(oldIgn.Storage.Directories, newIgn.Storage.Directories) {
		return errors.New("ignition directories section contains changes")
	}
	if !reflect.DeepEqual(oldIgn.Storage.Links, newIgn.Storage.Links) {
		if len(newIgn.Storage.Links) != 0 {
			return errors.New("ignition links section contains changes")
		}
	}
	for _, f := range newIgn.Storage.Files {
		if f.Append {
			return fmt.Errorf("ignition file %v includes append", f.Path)
		}
	}
	glog.V(2).Info("Configs are reconcilable")
	return nil
}
func verifyUserFields(pwdUser ignv2_2types.PasswdUser) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	emptyUser := ignv2_2types.PasswdUser{}
	tempUser := pwdUser
	if tempUser.Name == coreUserName && len(tempUser.SSHAuthorizedKeys) >= 1 {
		tempUser.Name = ""
		tempUser.SSHAuthorizedKeys = nil
		if !reflect.DeepEqual(emptyUser, tempUser) {
			return errors.New("ignition passwd user section contains unsupported changes: non-sshKey changes")
		}
		glog.Info("SSH Keys reconcilable")
	} else {
		return errors.New("ignition passwd user section contains unsupported changes: user must be core and have 1 or more sshKeys")
	}
	return nil
}
func (dn *Daemon) updateFiles(oldConfig, newConfig *mcfgv1.MachineConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	glog.Info("Updating files")
	if err := dn.writeFiles(newConfig.Spec.Config.Storage.Files); err != nil {
		return err
	}
	if err := dn.writeUnits(newConfig.Spec.Config.Systemd.Units); err != nil {
		return err
	}
	if err := dn.deleteStaleData(oldConfig, newConfig); err != nil {
		return err
	}
	return nil
}
func (dn *Daemon) deleteStaleData(oldConfig, newConfig *mcfgv1.MachineConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	glog.Info("Deleting stale data")
	newFileSet := make(map[string]struct{})
	for _, f := range newConfig.Spec.Config.Storage.Files {
		newFileSet[f.Path] = struct{}{}
	}
	for _, f := range oldConfig.Spec.Config.Storage.Files {
		if _, ok := newFileSet[f.Path]; !ok {
			glog.V(2).Infof("Deleting stale config file: %s", f.Path)
			if err := os.Remove(f.Path); err != nil {
				newErr := fmt.Errorf("unable to delete %s: %s", f.Path, err)
				if !os.IsNotExist(err) {
					return newErr
				}
				glog.Warningf("%v", newErr)
			}
			glog.Infof("Removed stale file %q", f.Path)
		}
	}
	newUnitSet := make(map[string]struct{})
	newDropinSet := make(map[string]struct{})
	for _, u := range newConfig.Spec.Config.Systemd.Units {
		for j := range u.Dropins {
			path := filepath.Join(pathSystemd, u.Name+".d", u.Dropins[j].Name)
			newDropinSet[path] = struct{}{}
		}
		path := filepath.Join(pathSystemd, u.Name)
		newUnitSet[path] = struct{}{}
	}
	for _, u := range oldConfig.Spec.Config.Systemd.Units {
		for j := range u.Dropins {
			path := filepath.Join(pathSystemd, u.Name+".d", u.Dropins[j].Name)
			if _, ok := newDropinSet[path]; !ok {
				glog.V(2).Infof("Deleting stale systemd dropin file: %s", path)
				if err := os.Remove(path); err != nil {
					newErr := fmt.Errorf("unable to delete %s: %s", path, err)
					if !os.IsNotExist(err) {
						return newErr
					}
					glog.Warningf("%v", newErr)
				}
				glog.Infof("Removed stale systemd dropin %q", path)
			}
		}
		path := filepath.Join(pathSystemd, u.Name)
		if _, ok := newUnitSet[path]; !ok {
			if err := dn.disableUnit(u); err != nil {
				glog.Warningf("Unable to disable %s: %s", u.Name, err)
			}
			glog.V(2).Infof("Deleting stale systemd unit file: %s", path)
			if err := os.Remove(path); err != nil {
				newErr := fmt.Errorf("unable to delete %s: %s", path, err)
				if !os.IsNotExist(err) {
					return newErr
				}
				glog.Warningf("%v", newErr)
			}
			glog.Infof("Removed stale systemd unit %q", path)
		}
	}
	return nil
}
func (dn *Daemon) enableUnit(unit ignv2_2types.Unit) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wantsPath := filepath.Join(wantsPathSystemd, unit.Name)
	if _, err := os.Stat(wantsPath); err == nil {
		glog.Infof("%s already exists. Not making a new symlink", wantsPath)
		return nil
	}
	servicePath := filepath.Join(pathSystemd, unit.Name)
	err := renameio.Symlink(servicePath, wantsPath)
	if err != nil {
		glog.Warningf("Cannot enable unit %s: %s", unit.Name, err)
	} else {
		glog.Infof("Enabled %s", unit.Name)
		glog.V(2).Infof("Symlinked %s to %s", servicePath, wantsPath)
	}
	return err
}
func (dn *Daemon) disableUnit(unit ignv2_2types.Unit) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wantsPath := filepath.Join(wantsPathSystemd, unit.Name)
	if _, err := os.Stat(wantsPath); err != nil {
		glog.Infof("%s was not present. No need to remove", wantsPath)
		return nil
	}
	glog.V(2).Infof("Disabling unit at %s", wantsPath)
	return os.Remove(wantsPath)
}
func (dn *Daemon) writeUnits(units []ignv2_2types.Unit) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, u := range units {
		for i := range u.Dropins {
			glog.Infof("Writing systemd unit dropin %q", u.Dropins[i].Name)
			dpath := filepath.Join(pathSystemd, u.Name+".d", u.Dropins[i].Name)
			if err := writeFileAtomicallyWithDefaults(dpath, []byte(u.Dropins[i].Contents)); err != nil {
				return fmt.Errorf("failed to write systemd unit dropin %q: %v", u.Dropins[i].Name, err)
			}
			glog.V(2).Infof("Wrote systemd unit dropin at %s", dpath)
		}
		if u.Contents == "" {
			continue
		}
		glog.Infof("Writing systemd unit %q", u.Name)
		fpath := filepath.Join(pathSystemd, u.Name)
		if u.Mask {
			glog.V(2).Info("Systemd unit masked")
			if err := os.RemoveAll(fpath); err != nil {
				return fmt.Errorf("failed to remove unit %q: %v", u.Name, err)
			}
			glog.V(2).Infof("Removed unit %q", u.Name)
			if err := renameio.Symlink(pathDevNull, fpath); err != nil {
				return fmt.Errorf("failed to symlink unit %q to %s: %v", u.Name, pathDevNull, err)
			}
			glog.V(2).Infof("Created symlink unit %q to %s", u.Name, pathDevNull)
			continue
		}
		if err := writeFileAtomicallyWithDefaults(fpath, []byte(u.Contents)); err != nil {
			return fmt.Errorf("failed to write systemd unit %q: %v", u.Name, err)
		}
		glog.V(2).Infof("Successfully wrote systemd unit %q: ", u.Name)
		glog.Infof("Enabling systemd unit %q", u.Name)
		if u.Enable {
			if err := dn.enableUnit(u); err != nil {
				return err
			}
			glog.V(2).Infof("Enabled systemd unit %q", u.Name)
		}
		if u.Enabled != nil {
			if *u.Enabled {
				if err := dn.enableUnit(u); err != nil {
					return err
				}
				glog.V(2).Infof("Enabled systemd unit %q", u.Name)
			} else {
				if err := dn.disableUnit(u); err != nil {
					return err
				}
				glog.V(2).Infof("Disabled systemd unit %q", u.Name)
			}
		}
	}
	return nil
}
func (dn *Daemon) writeFiles(files []ignv2_2types.File) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, file := range files {
		glog.Infof("Writing file %q", file.Path)
		contents, err := dataurl.DecodeString(file.Contents.Source)
		if err != nil {
			return err
		}
		mode := defaultFilePermissions
		if file.Mode != nil {
			mode = os.FileMode(*file.Mode)
		}
		var (
			uid, gid = -1, -1
		)
		if file.User != nil || file.Group != nil {
			uid, gid, err = getFileOwnership(file)
			if err != nil {
				return fmt.Errorf("failed to retrieve file ownership for file %q: %v", file.Path, err)
			}
		}
		if err := writeFileAtomically(file.Path, contents.Data, defaultDirectoryPermissions, mode, uid, gid); err != nil {
			return err
		}
	}
	return nil
}
func getFileOwnership(file ignv2_2types.File) (int, int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	uid, gid := 0, 0
	if file.User != nil {
		if file.User.ID != nil {
			uid = *file.User.ID
		} else if file.User.Name != "" {
			osUser, err := user.Lookup(file.User.Name)
			if err != nil {
				return uid, gid, fmt.Errorf("failed to retrieve UserID for username: %s", file.User.Name)
			}
			glog.V(2).Infof("Retrieved UserId: %s for username: %s", osUser.Uid, file.User.Name)
			uid, _ = strconv.Atoi(osUser.Uid)
		}
	}
	if file.Group != nil {
		if file.Group.ID != nil {
			gid = *file.Group.ID
		} else if file.Group.Name != "" {
			osGroup, err := user.LookupGroup(file.Group.Name)
			if err != nil {
				return uid, gid, fmt.Errorf("failed to retrieve GroupID for group: %s", file.Group.Name)
			}
			glog.V(2).Infof("Retrieved GroupID: %s for group: %s", osGroup.Gid, file.Group.Name)
			gid, _ = strconv.Atoi(osGroup.Gid)
		}
	}
	return uid, gid, nil
}
func (dn *Daemon) atomicallyWriteSSHKey(newUser ignv2_2types.PasswdUser, keys string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	authKeyPath := filepath.Join(coreUserSSHPath, "authorized_keys")
	glog.Infof("Writing SSHKeys at %q", authKeyPath)
	if err := writeFileAtomicallyWithDefaults(authKeyPath, []byte(keys)); err != nil {
		return err
	}
	glog.V(2).Infof("Wrote SSHKeys at %s", authKeyPath)
	return nil
}
func (dn *Daemon) updateSSHKeys(newUsers []ignv2_2types.PasswdUser) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(newUsers) == 0 {
		return nil
	}
	var concatSSHKeys string
	for _, k := range newUsers[len(newUsers)-1].SSHAuthorizedKeys {
		concatSSHKeys = concatSSHKeys + string(k) + "\n"
	}
	if err := dn.atomicSSHKeysWriter(newUsers[0], concatSSHKeys); err != nil {
		return err
	}
	return nil
}
func (dn *Daemon) updateOS(config *mcfgv1.MachineConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if dn.OperatingSystem != machineConfigDaemonOSRHCOS {
		glog.V(2).Info("Updating of non RHCOS nodes are not supported")
		return nil
	}
	newURL := config.Spec.OSImageURL
	osMatch, err := compareOSImageURL(dn.bootedOSImageURL, newURL)
	if err != nil {
		return err
	}
	if osMatch {
		return nil
	}
	if dn.recorder != nil {
		dn.recorder.Eventf(getNodeRef(dn.node), corev1.EventTypeNormal, "InClusterUpgrade", fmt.Sprintf("In cluster upgrade to %s", newURL))
	}
	glog.Infof("Updating OS to %s", newURL)
	if err := dn.NodeUpdaterClient.RunPivot(newURL); err != nil {
		return fmt.Errorf("failed to run pivot: %v", err)
	}
	return nil
}
func (dn *Daemon) logSystem(format string, a ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	message := fmt.Sprintf(format, a...)
	glog.Info(message)
	logger := exec.Command("logger")
	stdin, err := logger.StdinPipe()
	if err != nil {
		glog.Errorf("failed to get stdin pipe: %v", err)
		return
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, message)
	}()
	err = logger.Run()
	if err != nil {
		glog.Errorf("failed to invoke logger: %v", err)
		return
	}
}
func (dn *Daemon) cancelSIGTERM() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if dn.installedSigterm {
		signal.Reset(syscall.SIGTERM)
		dn.installedSigterm = false
	}
}
func (dn *Daemon) reboot(rationale string, timeout time.Duration, rebootCmd *exec.Cmd) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if dn.recorder != nil {
		dn.recorder.Eventf(getNodeRef(dn.node), corev1.EventTypeNormal, "Reboot", rationale)
	}
	dn.logSystem("machine-config-daemon initiating reboot: %s", rationale)
	dn.cancelSIGTERM()
	dn.Close()
	if dn.skipReboot && dn.onceFrom != "" {
		glog.Info("MCD is not rebooting in onceFrom with --skip-reboot")
		return nil
	}
	err := rebootCmd.Run()
	if err != nil {
		return errors.Wrapf(err, "failed to reboot")
	}
	time.Sleep(timeout)
	return fmt.Errorf("reboot failed; this error should be unreachable, something is seriously wrong")
}

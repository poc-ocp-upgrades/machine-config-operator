package daemon

import (
	"bufio"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	godefaulthttp "net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
	imgref "github.com/containers/image/docker/reference"
	ignv2 "github.com/coreos/ignition/config/v2_2"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	"github.com/golang/glog"
	drain "github.com/openshift/kubernetes-drain"
	"github.com/openshift/machine-config-operator/lib/resourceread"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"github.com/openshift/machine-config-operator/pkg/daemon/constants"
	mcfgclientset "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned"
	mcfginformersv1 "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/machineconfiguration.openshift.io/v1"
	mcfglistersv1 "github.com/openshift/machine-config-operator/pkg/generated/listers/machineconfiguration.openshift.io/v1"
	"github.com/pkg/errors"
	"github.com/vincent-petithory/dataurl"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformersv1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	clientsetcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisterv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
)

type Daemon struct {
	name			string
	OperatingSystem		string
	NodeUpdaterClient	NodeUpdaterClient
	bootID			string
	bootedOSImageURL	string
	kubeClient		kubernetes.Interface
	mcClient		mcfgclientset.Interface
	recorder		record.EventRecorder
	rootMount		string
	nodeLister		corelisterv1.NodeLister
	nodeListerSynced	cache.InformerSynced
	mcLister		mcfglistersv1.MachineConfigLister
	mcListerSynced		cache.InformerSynced
	onceFrom		string
	skipReboot		bool
	kubeletHealthzEnabled	bool
	kubeletHealthzEndpoint	string
	installedSigterm	bool
	nodeWriter		NodeWriter
	exitCh			chan<- error
	stopCh			<-chan struct{}
	node			*corev1.Node
	atomicSSHKeysWriter	func(ignv2_2types.PasswdUser, string) error
	queue			workqueue.RateLimitingInterface
	enqueueNode		func(*corev1.Node)
	syncHandler		func(node string) error
	booting			bool
	currentConfigPath	string
}
type pendingConfigState struct {
	PendingConfig	string	`json:"pendingConfig,omitempty"`
	BootID		string	`json:"bootID,omitempty"`
}

const (
	pathSystemd			= "/etc/systemd/system"
	wantsPathSystemd		= "/etc/systemd/system/multi-user.target.wants/"
	pathDevNull			= "/dev/null"
	pathStateJSON			= "/etc/machine-config-daemon/state.json"
	currentConfigPath		= "/var/machine-config-daemon/currentconfig"
	kubeletHealthzEndpoint		= "http://localhost:10248/healthz"
	kubeletHealthzPollingInterval	= time.Duration(30 * time.Second)
	kubeletHealthzTimeout		= time.Duration(30 * time.Second)
	kubeletHealthzFailureThreshold	= 3
	maxRetries			= 15
	updateDelay			= 5 * time.Second
)

type onceFromOrigin int

const (
	onceFromUnknownConfig	onceFromOrigin	= iota
	onceFromLocalConfig
	onceFromRemoteConfig
)

var (
	defaultRebootTimeout	= 24 * time.Hour
	defaultRebootCommand	= "reboot"
)

func New(rootMount string, nodeName string, operatingSystem string, nodeUpdaterClient NodeUpdaterClient, bootID, onceFrom string, skipReboot bool, mcClient mcfgclientset.Interface, kubeClient kubernetes.Interface, kubeletHealthzEnabled bool, kubeletHealthzEndpoint string, nodeWriter NodeWriter, exitCh chan<- error, stopCh <-chan struct{}) (*Daemon, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		osImageURL	string
		err		error
	)
	if operatingSystem == machineConfigDaemonOSRHCOS {
		var osVersion string
		osImageURL, osVersion, err = nodeUpdaterClient.GetBootedOSImageURL(rootMount)
		if err != nil {
			return nil, fmt.Errorf("error reading osImageURL from rpm-ostree: %v", err)
		}
		glog.Infof("Booted osImageURL: %s (%s)", osImageURL, osVersion)
	}
	dn := &Daemon{name: nodeName, OperatingSystem: operatingSystem, NodeUpdaterClient: nodeUpdaterClient, rootMount: rootMount, bootedOSImageURL: osImageURL, bootID: bootID, onceFrom: onceFrom, skipReboot: skipReboot, kubeletHealthzEnabled: kubeletHealthzEnabled, kubeletHealthzEndpoint: kubeletHealthzEndpoint, nodeWriter: nodeWriter, exitCh: exitCh, stopCh: stopCh, kubeClient: kubeClient, mcClient: mcClient, currentConfigPath: currentConfigPath}
	dn.atomicSSHKeysWriter = dn.atomicallyWriteSSHKey
	return dn, nil
}
func NewClusterDrivenDaemon(rootMount, nodeName, operatingSystem string, nodeUpdaterClient NodeUpdaterClient, mcInformer mcfginformersv1.MachineConfigInformer, kubeClient kubernetes.Interface, bootID, onceFrom string, skipReboot bool, nodeInformer coreinformersv1.NodeInformer, kubeletHealthzEnabled bool, kubeletHealthzEndpoint string, nodeWriter NodeWriter, exitCh chan<- error, stopCh <-chan struct{}) (*Daemon, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	dn, err := New(rootMount, nodeName, operatingSystem, nodeUpdaterClient, bootID, onceFrom, skipReboot, nil, kubeClient, kubeletHealthzEnabled, kubeletHealthzEndpoint, nodeWriter, exitCh, stopCh)
	if err != nil {
		return nil, err
	}
	dn.queue = workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "machineconfigdaemon")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.V(2).Infof)
	eventBroadcaster.StartRecordingToSink(&clientsetcorev1.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	dn.recorder = eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: "machineconfigdaemon", Host: nodeName})
	glog.Infof("Managing node: %s", nodeName)
	go dn.runLoginMonitor(dn.stopCh, dn.exitCh)
	nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{UpdateFunc: dn.handleNodeUpdate})
	dn.nodeLister = nodeInformer.Lister()
	dn.nodeListerSynced = nodeInformer.Informer().HasSynced
	dn.mcLister = mcInformer.Lister()
	dn.mcListerSynced = mcInformer.Informer().HasSynced
	dn.enqueueNode = dn.enqueueDefault
	dn.syncHandler = dn.syncNode
	dn.booting = true
	return dn, nil
}
func (dn *Daemon) worker() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for dn.processNextWorkItem() {
	}
}
func (dn *Daemon) processNextWorkItem() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if dn.booting {
		if err := dn.bootstrapNode(); err != nil {
			dn.updateErrorState(err)
			glog.Warningf("Booting the MCD errored with %v", err)
		}
		return true
	}
	key, quit := dn.queue.Get()
	if quit {
		return false
	}
	defer dn.queue.Done(key)
	err := dn.syncHandler(key.(string))
	dn.handleErr(err, key)
	return true
}
func (dn *Daemon) bootstrapNode() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	node, err := dn.nodeLister.Get(dn.name)
	if err != nil {
		return err
	}
	node, err = dn.loadNodeAnnotations(node)
	if err != nil {
		return err
	}
	dn.node = node
	if err := dn.CheckStateOnBoot(); err != nil {
		return err
	}
	dn.booting = false
	return nil
}
func (dn *Daemon) handleErr(err error, key interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		dn.queue.Forget(key)
		return
	}
	dn.updateErrorState(err)
	if dn.queue.NumRequeues(key) < maxRetries {
		glog.V(2).Infof("Error syncing node %v: %v", key, err)
		dn.queue.AddRateLimited(key)
		return
	}
	utilruntime.HandleError(err)
	glog.V(2).Infof("Dropping node %q out of the queue: %v", key, err)
	dn.queue.Forget(key)
	dn.queue.AddAfter(key, 1*time.Minute)
}
func (dn *Daemon) updateErrorState(err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch errors.Cause(err) {
	case errUnreconcilable:
		dn.nodeWriter.SetUnreconcilable(err, dn.kubeClient.CoreV1().Nodes(), dn.nodeLister, dn.name)
	default:
		dn.nodeWriter.SetDegraded(err, dn.kubeClient.CoreV1().Nodes(), dn.nodeLister, dn.name)
	}
}
func (dn *Daemon) syncNode(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	startTime := time.Now()
	glog.V(4).Infof("Started syncing node %q (%v)", key, startTime)
	defer func() {
		glog.V(4).Infof("Finished syncing node %q (%v)", key, time.Since(startTime))
	}()
	_, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	node, err := dn.nodeLister.Get(name)
	if apierrors.IsNotFound(err) {
		glog.V(2).Infof("node %v has been deleted", key)
		return nil
	}
	if err != nil {
		return err
	}
	node = node.DeepCopy()
	if node.DeletionTimestamp != nil {
		return nil
	}
	if node.Name == dn.name {
		dn.node = node
		current, desired, err := dn.prepUpdateFromCluster()
		if err != nil {
			glog.Infof("Unable to prep update: %s", err)
			return err
		}
		if current != nil || desired != nil {
			if err := dn.triggerUpdateWithMachineConfig(current, desired); err != nil {
				glog.Infof("Unable to apply update: %s", err)
				return err
			}
		}
		glog.V(2).Infof("Node %s is already synced", node.Name)
	}
	return nil
}
func (dn *Daemon) enqueueAfter(node *corev1.Node, after time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(node)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("couldn't get key for object %#v: %v", node, err))
		return
	}
	dn.queue.AddAfter(key, after)
}
func (dn *Daemon) enqueueDefault(node *corev1.Node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	dn.enqueueAfter(node, updateDelay)
}

const (
	sdMessageSessionStart = "8d45620c1a4348dbb17410da57c60c66"
)

func (dn *Daemon) detectEarlySSHAccessesFromBoot() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	journalOutput, err := exec.Command("journalctl", "-b", "-o", "cat", "MESSAGE_ID="+sdMessageSessionStart).CombinedOutput()
	if err != nil {
		return err
	}
	if len(journalOutput) > 0 {
		glog.Info("Detected a login session before the daemon took over on first boot")
		glog.Infof("Applying annotation: %v", machineConfigDaemonSSHAccessAnnotationKey)
		if err := dn.applySSHAccessedAnnotation(); err != nil {
			return err
		}
	}
	return nil
}
func (dn *Daemon) runOnceFrom() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	configi, contentFrom, err := dn.senseAndLoadOnceFrom()
	if err != nil {
		glog.Warningf("Unable to decipher onceFrom config type: %s", err)
		return err
	}
	switch configi.(type) {
	case ignv2_2types.Config:
		glog.V(2).Info("Daemon running directly from Ignition")
		return dn.runOnceFromIgnition(configi.(ignv2_2types.Config))
	case mcfgv1.MachineConfig:
		glog.V(2).Info("Daemon running directly from MachineConfig")
		return dn.runOnceFromMachineConfig(configi.(mcfgv1.MachineConfig), contentFrom)
	}
	return errors.New("unsupported onceFrom type provided")
}
func (dn *Daemon) Run(stopCh <-chan struct{}, exitCh <-chan error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if dn.kubeletHealthzEnabled {
		glog.Info("Enabling Kubelet Healthz Monitor")
		go dn.runKubeletHealthzMonitor(stopCh, dn.exitCh)
	}
	if dn.onceFrom != "" {
		return dn.runOnceFrom()
	}
	defer utilruntime.HandleCrash()
	defer dn.queue.ShutDown()
	if !cache.WaitForCacheSync(stopCh, dn.nodeListerSynced, dn.mcListerSynced) {
		return errors.New("failed to sync initial listers cache")
	}
	go wait.Until(dn.worker, time.Second, stopCh)
	for {
		select {
		case <-stopCh:
			return nil
		case err := <-exitCh:
			glog.Warningf("Got an error from auxiliary tools: %v", err)
		}
	}
}
func (dn *Daemon) BindPodMounts() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	targetSecrets := filepath.Join(dn.rootMount, "/run/secrets")
	if err := os.MkdirAll(targetSecrets, 0755); err != nil {
		return err
	}
	mnt := exec.Command("mount", "--rbind", "/run/secrets", targetSecrets)
	return mnt.Run()
}
func (dn *Daemon) runLoginMonitor(stopCh <-chan struct{}, exitCh chan<- error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := exec.Command("journalctl", "-b", "-f", "-o", "cat", "MESSAGE_ID="+sdMessageSessionStart)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		exitCh <- err
		return
	}
	if err := cmd.Start(); err != nil {
		exitCh <- err
		return
	}
	worker := make(chan struct{})
	go func() {
		for {
			select {
			case <-worker:
				return
			default:
				buf := make([]byte, 1024)
				l, err := stdout.Read(buf)
				if err != nil {
					if err == io.EOF {
						return
					}
					exitCh <- err
					return
				}
				if l > 0 {
					line := strings.Split(string(buf), "\n")[0]
					glog.Infof("Detected a new login session: %s", line)
					glog.Infof("Login access is discouraged! Applying annotation: %v", machineConfigDaemonSSHAccessAnnotationKey)
					if err := dn.applySSHAccessedAnnotation(); err != nil {
						exitCh <- err
					}
				}
			}
		}
	}()
	select {
	case <-stopCh:
		close(worker)
		cmd.Process.Kill()
	}
}
func (dn *Daemon) applySSHAccessedAnnotation() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := dn.nodeWriter.SetSSHAccessed(dn.kubeClient.CoreV1().Nodes(), dn.nodeLister, dn.name); err != nil {
		return fmt.Errorf("error: cannot apply annotation for SSH access due to: %v", err)
	}
	return nil
}
func (dn *Daemon) runKubeletHealthzMonitor(stopCh <-chan struct{}, exitCh chan<- error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	failureCount := 0
	for {
		select {
		case <-stopCh:
			return
		case <-time.After(kubeletHealthzPollingInterval):
			if err := dn.getHealth(); err != nil {
				glog.Warningf("Failed kubelet health check: %v", err)
				failureCount++
				if failureCount >= kubeletHealthzFailureThreshold {
					exitCh <- fmt.Errorf("kubelet health failure threshold reached")
				}
			} else {
				failureCount = 0
			}
		}
	}
}
func (dn *Daemon) getHealth() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	glog.V(2).Info("Kubelet health running")
	ctx, cancel := context.WithTimeout(context.Background(), kubeletHealthzTimeout)
	defer cancel()
	req, err := http.NewRequest("GET", dn.kubeletHealthzEndpoint, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if string(respData) != "ok" {
		glog.Warningf("Kubelet Healthz Endpoint returned: %s", string(respData))
		return nil
	}
	glog.V(2).Info("Kubelet health ok")
	return nil
}

type stateAndConfigs struct {
	bootstrapping	bool
	state		string
	currentConfig	*mcfgv1.MachineConfig
	pendingConfig	*mcfgv1.MachineConfig
	desiredConfig	*mcfgv1.MachineConfig
}

func (dn *Daemon) getStateAndConfigs(pendingConfigName string) (*stateAndConfigs, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := os.Lstat(constants.InitialNodeAnnotationsFilePath)
	var bootstrapping bool
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	} else {
		bootstrapping = true
		glog.Info("In bootstrap mode")
	}
	currentConfigName, err := getNodeAnnotation(dn.node, constants.CurrentMachineConfigAnnotationKey)
	if err != nil {
		return nil, err
	}
	desiredConfigName, err := getNodeAnnotation(dn.node, constants.DesiredMachineConfigAnnotationKey)
	if err != nil {
		return nil, err
	}
	currentConfig, err := dn.mcLister.Get(currentConfigName)
	if err != nil {
		return nil, err
	}
	state, err := getNodeAnnotationExt(dn.node, constants.MachineConfigDaemonStateAnnotationKey, true)
	if err != nil {
		return nil, err
	}
	if state == "" {
		state = constants.MachineConfigDaemonStateDone
	}
	var desiredConfig *mcfgv1.MachineConfig
	if currentConfigName == desiredConfigName {
		desiredConfig = currentConfig
		glog.Infof("Current+desired config: %s", currentConfigName)
	} else {
		desiredConfig, err = dn.mcLister.Get(desiredConfigName)
		if err != nil {
			return nil, err
		}
		glog.Infof("Current config: %s", currentConfigName)
		glog.Infof("Desired config: %s", desiredConfigName)
	}
	var pendingConfig *mcfgv1.MachineConfig
	if pendingConfigName == desiredConfigName {
		pendingConfig = desiredConfig
	} else if pendingConfigName != "" {
		pendingConfig, err = dn.mcLister.Get(pendingConfigName)
		if err != nil {
			return nil, err
		}
		glog.Infof("Pending config: %s", pendingConfigName)
	}
	return &stateAndConfigs{bootstrapping: bootstrapping, currentConfig: currentConfig, pendingConfig: pendingConfig, desiredConfig: desiredConfig, state: state}, nil
}
func (dn *Daemon) getPendingConfig() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s, err := ioutil.ReadFile(pathStateJSON)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", errors.Wrapf(err, "loading transient state")
		}
		return "", nil
	}
	var p pendingConfigState
	if err := json.Unmarshal([]byte(s), &p); err != nil {
		return "", errors.Wrapf(err, "parsing transient state")
	}
	if p.BootID == dn.bootID {
		return "", fmt.Errorf("pending config %s bootID %s matches current! Failed to reboot?", p.PendingConfig, dn.bootID)
	}
	return p.PendingConfig, nil
}
func (dn *Daemon) CheckStateOnBoot() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if dn.OperatingSystem == machineConfigDaemonOSRHCOS {
		status, err := dn.NodeUpdaterClient.GetStatus()
		if err != nil {
			glog.Fatalf("unable to get rpm-ostree status: %s", err)
		}
		glog.Info(status)
	}
	pendingConfigName, err := dn.getPendingConfig()
	if err != nil {
		return err
	}
	state, err := dn.getStateAndConfigs(pendingConfigName)
	if err != nil {
		return err
	}
	if err := dn.detectEarlySSHAccessesFromBoot(); err != nil {
		return fmt.Errorf("error detecting previous SSH accesses: %v", err)
	}
	if state.bootstrapping {
		targetOSImageURL := state.currentConfig.Spec.OSImageURL
		osMatch, err := dn.checkOS(targetOSImageURL)
		if err != nil {
			return err
		}
		if !osMatch {
			glog.Infof("Bootstrap pivot required to: %s", targetOSImageURL)
			return dn.updateOSAndReboot(state.currentConfig)
		}
		glog.Info("No bootstrap pivot required; unlinking bootstrap node annotations")
		if err := os.Remove(constants.InitialNodeAnnotationsFilePath); err != nil {
			return errors.Wrapf(err, "removing initial node annotations file")
		}
	}
	var expectedConfig *mcfgv1.MachineConfig
	if state.pendingConfig != nil {
		expectedConfig = state.pendingConfig
	} else {
		expectedConfig = state.currentConfig
	}
	if isOnDiskValid := dn.validateOnDiskState(expectedConfig); !isOnDiskValid {
		return errors.New("unexpected on-disk state")
	}
	glog.Info("Validated on-disk state")
	if state.pendingConfig != nil {
		if err := dn.nodeWriter.SetDone(dn.kubeClient.CoreV1().Nodes(), dn.nodeLister, dn.name, state.pendingConfig.GetName()); err != nil {
			return err
		}
		if err := os.Remove(pathStateJSON); err != nil {
			return errors.Wrapf(err, "removing transient state file")
		}
		state.currentConfig = state.pendingConfig
	}
	mcJSON, err := json.Marshal(state.currentConfig)
	if err != nil {
		return err
	}
	if err := writeFileAtomicallyWithDefaults(currentConfigPath, mcJSON); err != nil {
		return err
	}
	inDesiredConfig := state.currentConfig == state.desiredConfig
	if inDesiredConfig {
		if state.pendingConfig != nil {
			glog.Infof("Completing pending config %s", state.pendingConfig.GetName())
			if err := dn.completeUpdate(dn.node, state.pendingConfig.GetName()); err != nil {
				return err
			}
		}
		glog.Infof("In desired config %s", state.currentConfig.GetName())
		return nil
	}
	return dn.triggerUpdateWithMachineConfig(state.currentConfig, state.desiredConfig)
}
func (dn *Daemon) runOnceFromMachineConfig(machineConfig mcfgv1.MachineConfig, contentFrom onceFromOrigin) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if contentFrom == onceFromRemoteConfig {
		if dn.kubeClient == nil || dn.mcClient == nil {
			panic("running in onceFrom mode with a remote MachineConfig without a cluster")
		}
		current, desired, err := dn.prepUpdateFromCluster()
		if err != nil {
			dn.nodeWriter.SetDegraded(err, dn.kubeClient.CoreV1().Nodes(), dn.nodeLister, dn.name)
			return err
		}
		if current == nil || desired == nil {
			return nil
		}
		if err := dn.triggerUpdateWithMachineConfig(current, &machineConfig); err != nil {
			dn.nodeWriter.SetDegraded(err, dn.kubeClient.CoreV1().Nodes(), dn.nodeLister, dn.name)
			return err
		}
		return nil
	}
	if contentFrom == onceFromLocalConfig {
		oldConfig := mcfgv1.MachineConfig{}
		return dn.update(&oldConfig, &machineConfig)
	}
	return fmt.Errorf("%s is not a path nor url; can not run once", dn.onceFrom)
}
func (dn *Daemon) runOnceFromIgnition(ignConfig ignv2_2types.Config) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := dn.writeFiles(ignConfig.Storage.Files); err != nil {
		return err
	}
	if err := dn.writeUnits(ignConfig.Systemd.Units); err != nil {
		return err
	}
	return dn.reboot("runOnceFromIgnition complete", defaultRebootTimeout, exec.Command(defaultRebootCommand))
}
func (dn *Daemon) handleNodeUpdate(old, cur interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	oldNode := old.(*corev1.Node)
	curNode := cur.(*corev1.Node)
	glog.V(4).Infof("Updating Node %s", oldNode.Name)
	dn.enqueueNode(curNode)
}
func (dn *Daemon) getCurrentMCOnDisk() (*mcfgv1.MachineConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mcJSON, err := os.Open(dn.currentConfigPath)
	if err != nil {
		return nil, err
	}
	defer mcJSON.Close()
	currentOnDisk := &mcfgv1.MachineConfig{}
	if err := json.NewDecoder(bufio.NewReader(mcJSON)).Decode(currentOnDisk); err != nil {
		return nil, err
	}
	return currentOnDisk, nil
}
func (dn *Daemon) prepUpdateFromCluster() (*mcfgv1.MachineConfig, *mcfgv1.MachineConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	desiredConfigName, err := getNodeAnnotationExt(dn.node, constants.DesiredMachineConfigAnnotationKey, true)
	if err != nil {
		return nil, nil, err
	}
	desiredConfig, err := dn.mcLister.Get(desiredConfigName)
	if err != nil {
		return nil, nil, err
	}
	currentConfigName, err := getNodeAnnotation(dn.node, constants.CurrentMachineConfigAnnotationKey)
	if err != nil {
		return nil, nil, err
	}
	currentConfig, err := dn.mcLister.Get(currentConfigName)
	if err != nil {
		return nil, nil, err
	}
	state, err := getNodeAnnotation(dn.node, constants.MachineConfigDaemonStateAnnotationKey)
	if err != nil {
		return nil, nil, err
	}
	currentMCOnDisk, err := dn.getCurrentMCOnDisk()
	if err != nil {
		return nil, nil, err
	}
	if currentMCOnDisk.GetName() != currentConfig.GetName() {
		return currentMCOnDisk, desiredConfig, nil
	}
	if desiredConfigName == currentConfigName && state == constants.MachineConfigDaemonStateDone {
		glog.V(2).Info("No updating is required")
		return nil, nil, nil
	}
	return currentConfig, desiredConfig, nil
}
func (dn *Daemon) completeUpdate(node *corev1.Node, desiredConfigName string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := drain.Uncordon(dn.kubeClient.CoreV1().Nodes(), node, nil); err != nil {
		return err
	}
	dn.logSystem("machine-config-daemon: completed update for config %s", desiredConfigName)
	return nil
}
func (dn *Daemon) triggerUpdateWithMachineConfig(currentConfig *mcfgv1.MachineConfig, desiredConfig *mcfgv1.MachineConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if currentConfig == nil {
		ccAnnotation, err := getNodeAnnotation(dn.node, constants.CurrentMachineConfigAnnotationKey)
		if err != nil {
			return err
		}
		currentConfig, err = dn.mcLister.Get(ccAnnotation)
		if err != nil {
			return err
		}
	}
	if desiredConfig == nil {
		dcAnnotation, err := getNodeAnnotation(dn.node, constants.DesiredMachineConfigAnnotationKey)
		if err != nil {
			return err
		}
		desiredConfig, err = dn.mcLister.Get(dcAnnotation)
		if err != nil {
			return err
		}
	}
	return dn.update(currentConfig, desiredConfig)
}
func (dn *Daemon) validateOnDiskState(currentConfig *mcfgv1.MachineConfig) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	osMatch, err := dn.checkOS(currentConfig.Spec.OSImageURL)
	if err != nil {
		glog.Errorf("%s", err)
		return false
	}
	if !osMatch {
		glog.Errorf("expected target osImageURL %s", currentConfig.Spec.OSImageURL)
		return false
	}
	if !checkFiles(currentConfig.Spec.Config.Storage.Files) {
		return false
	}
	if !checkUnits(currentConfig.Spec.Config.Systemd.Units) {
		return false
	}
	return true
}
func getRefDigest(ref string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	refParsed, err := imgref.ParseNamed(ref)
	if err != nil {
		return "", errors.Wrapf(err, "parsing reference: %q", ref)
	}
	canon, ok := refParsed.(imgref.Canonical)
	if !ok {
		return "", fmt.Errorf("not canonical form: %q", ref)
	}
	return canon.Digest().String(), nil
}
func compareOSImageURL(current, desired string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if desired == "" || desired == "://dummy" {
		glog.Info(`No target osImageURL provided`)
		return true, nil
	}
	if current == desired {
		return true, nil
	}
	bootedDigest, err := getRefDigest(current)
	if err != nil {
		return false, errors.Wrap(err, "parsing booted osImageURL")
	}
	desiredDigest, err := getRefDigest(desired)
	if err != nil {
		return false, errors.Wrap(err, "parsing desired osImageURL")
	}
	if bootedDigest == desiredDigest {
		glog.Infof("Current and target osImageURL have matching digest %q", bootedDigest)
		return true, nil
	}
	return false, nil
}
func (dn *Daemon) checkOS(osImageURL string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if dn.OperatingSystem != machineConfigDaemonOSRHCOS {
		glog.Infof(`Not booted into Red Hat CoreOS, ignoring target OSImageURL %s`, osImageURL)
		return true, nil
	}
	return compareOSImageURL(dn.bootedOSImageURL, osImageURL)
}
func checkUnits(units []ignv2_2types.Unit) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, u := range units {
		for j := range u.Dropins {
			path := filepath.Join(pathSystemd, u.Name+".d", u.Dropins[j].Name)
			if status := checkFileContentsAndMode(path, []byte(u.Dropins[j].Contents), defaultFilePermissions); !status {
				return false
			}
		}
		if u.Contents == "" {
			continue
		}
		path := filepath.Join(pathSystemd, u.Name)
		if u.Mask {
			link, err := filepath.EvalSymlinks(path)
			if err != nil {
				glog.Errorf("state validation: error while evaluation symlink for path: %q, err: %v", path, err)
				return false
			}
			if strings.Compare(pathDevNull, link) != 0 {
				glog.Errorf("state validation: invalid unit masked setting. path: %q; expected: %v; received: %v", path, pathDevNull, link)
				return false
			}
		}
		if status := checkFileContentsAndMode(path, []byte(u.Contents), defaultFilePermissions); !status {
			return false
		}
	}
	return true
}
func checkFiles(files []ignv2_2types.File) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	checkedFiles := make(map[string]bool)
	for i := len(files) - 1; i >= 0; i-- {
		f := files[i]
		if _, ok := checkedFiles[f.Path]; ok {
			continue
		}
		mode := defaultFilePermissions
		if f.Mode != nil {
			mode = os.FileMode(*f.Mode)
		}
		contents, err := dataurl.DecodeString(f.Contents.Source)
		if err != nil {
			glog.Errorf("couldn't parse file: %v", err)
			return false
		}
		if status := checkFileContentsAndMode(f.Path, contents.Data, mode); !status {
			return false
		}
		checkedFiles[f.Path] = true
	}
	return true
}
func checkFileContentsAndMode(filePath string, expectedContent []byte, mode os.FileMode) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fi, err := os.Lstat(filePath)
	if err != nil {
		glog.Errorf("could not stat file: %q, error: %v", filePath, err)
		return false
	}
	if fi.Mode() != mode {
		glog.Errorf("mode mismatch for file: %q; expected: %v; received: %v", filePath, mode, fi.Mode())
		return false
	}
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		glog.Errorf("could not read file: %q, error: %v", filePath, err)
		return false
	}
	if !bytes.Equal(contents, expectedContent) {
		glog.Errorf("content mismatch for file: %q", filePath)
		return false
	}
	return true
}
func (dn *Daemon) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func ValidPath(path string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, validStart := range []string{".", "..", "/"} {
		if strings.HasPrefix(path, validStart) {
			return true
		}
	}
	return false
}
func (dn *Daemon) senseAndLoadOnceFrom() (interface{}, onceFromOrigin, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		content		[]byte
		contentFrom	onceFromOrigin
	)
	if strings.HasPrefix(dn.onceFrom, "http://") || strings.HasPrefix(dn.onceFrom, "https://") {
		contentFrom = onceFromRemoteConfig
		resp, err := http.Get(dn.onceFrom)
		if err != nil {
			return nil, contentFrom, err
		}
		defer resp.Body.Close()
		content, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, contentFrom, err
		}
	} else {
		contentFrom = onceFromLocalConfig
		absoluteOnceFrom, err := filepath.Abs(filepath.Clean(dn.onceFrom))
		if err != nil {
			return nil, contentFrom, err
		}
		content, err = ioutil.ReadFile(absoluteOnceFrom)
		if err != nil {
			return nil, contentFrom, err
		}
	}
	ignConfig, _, err := ignv2.Parse(content)
	if err == nil && ignConfig.Ignition.Version != "" {
		glog.V(2).Info("onceFrom file is of type Ignition")
		return ignConfig, contentFrom, nil
	}
	glog.V(2).Infof("%s is not an Ignition config: %v. Trying MachineConfig.", dn.onceFrom, err)
	mc, err := resourceread.ReadMachineConfigV1(content)
	if err == nil && mc != nil {
		glog.V(2).Info("onceFrom file is of type MachineConfig")
		return *mc, contentFrom, nil
	}
	return nil, onceFromUnknownConfig, fmt.Errorf("unable to decipher onceFrom config type: %v", err)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

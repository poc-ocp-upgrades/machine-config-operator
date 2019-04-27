package daemon

import (
	"fmt"
	"os/exec"
	"testing"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	k8sfake "k8s.io/client-go/kubernetes/fake"
)

func TestUpdateOS(t *testing.T) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	expectedError := fmt.Errorf("broken")
	testClient := RpmOstreeClientMock{GetBootedOSImageURLReturns: []GetBootedOSImageURLReturn{}, RunPivotReturns: []error{nil, expectedError}}
	d := Daemon{name: "nodeName", OperatingSystem: machineConfigDaemonOSRHCOS, NodeUpdaterClient: testClient, kubeClient: k8sfake.NewSimpleClientset(), rootMount: "/", bootedOSImageURL: "test"}
	mcfg := &mcfgv1.MachineConfig{}
	differentMcfg := &mcfgv1.MachineConfig{Spec: mcfgv1.MachineConfigSpec{OSImageURL: "somethingDifferent"}}
	if err := d.updateOS(mcfg); err != nil {
		t.Errorf("Expected no error. Got %s.", err)
	}
	if err := d.updateOS(differentMcfg); err == expectedError {
		t.Error("Expected an error. Got none.")
	}
}
func TestReconcilable(t *testing.T) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := Daemon{name: "nodeName", OperatingSystem: machineConfigDaemonOSRHCOS, NodeUpdaterClient: nil, kubeClient: nil, rootMount: "/", bootedOSImageURL: "test"}
	oldConfig := &mcfgv1.MachineConfig{Spec: mcfgv1.MachineConfigSpec{Config: ignv2_2types.Config{Ignition: ignv2_2types.Ignition{Version: "2.0.0"}}}}
	newConfig := &mcfgv1.MachineConfig{Spec: mcfgv1.MachineConfigSpec{Config: ignv2_2types.Config{Ignition: ignv2_2types.Ignition{Version: "2.2.0"}}}}
	isReconcilable := d.reconcilable(oldConfig, newConfig)
	checkIrreconcilableResults(t, "Ignition", isReconcilable)
	oldConfig.Spec.Config.Ignition.Version = "2.2.0"
	isReconcilable = d.reconcilable(oldConfig, newConfig)
	checkReconcilableResults(t, "Ignition", isReconcilable)
	oldConfig.Spec.Config.Networkd = ignv2_2types.Networkd{}
	newConfig.Spec.Config.Networkd = ignv2_2types.Networkd{Units: []ignv2_2types.Networkdunit{ignv2_2types.Networkdunit{Name: "test.network"}}}
	isReconcilable = d.reconcilable(oldConfig, newConfig)
	checkIrreconcilableResults(t, "Networkd", isReconcilable)
	oldConfig.Spec.Config.Networkd = newConfig.Spec.Config.Networkd
	isReconcilable = d.reconcilable(oldConfig, newConfig)
	checkReconcilableResults(t, "Networkd", isReconcilable)
	oldConfig.Spec.Config.Storage.Disks = []ignv2_2types.Disk{ignv2_2types.Disk{Device: "/one"}}
	isReconcilable = d.reconcilable(oldConfig, newConfig)
	checkIrreconcilableResults(t, "Disk", isReconcilable)
	newConfig.Spec.Config.Storage.Disks = oldConfig.Spec.Config.Storage.Disks
	isReconcilable = d.reconcilable(oldConfig, newConfig)
	checkReconcilableResults(t, "Disk", isReconcilable)
	oldFSPath := "/foo/bar"
	oldConfig.Spec.Config.Storage.Filesystems = []ignv2_2types.Filesystem{ignv2_2types.Filesystem{Name: "user", Path: &oldFSPath}}
	isReconcilable = d.reconcilable(oldConfig, newConfig)
	checkIrreconcilableResults(t, "Filesystem", isReconcilable)
	newConfig.Spec.Config.Storage.Filesystems = oldConfig.Spec.Config.Storage.Filesystems
	isReconcilable = d.reconcilable(oldConfig, newConfig)
	checkReconcilableResults(t, "Filesystem", isReconcilable)
	oldConfig.Spec.Config.Storage.Raid = []ignv2_2types.Raid{ignv2_2types.Raid{Name: "data", Level: "stripe"}}
	isReconcilable = d.reconcilable(oldConfig, newConfig)
	checkIrreconcilableResults(t, "Raid", isReconcilable)
	newConfig.Spec.Config.Storage.Raid = oldConfig.Spec.Config.Storage.Raid
	isReconcilable = d.reconcilable(oldConfig, newConfig)
	checkReconcilableResults(t, "Raid", isReconcilable)
	oldConfig = &mcfgv1.MachineConfig{}
	tempGroup := ignv2_2types.PasswdGroup{Name: "testGroup"}
	newMcfg := &mcfgv1.MachineConfig{Spec: mcfgv1.MachineConfigSpec{Config: ignv2_2types.Config{Passwd: ignv2_2types.Passwd{Groups: []ignv2_2types.PasswdGroup{tempGroup}}}}}
	isReconcilable = d.reconcilable(oldConfig, newMcfg)
	checkIrreconcilableResults(t, "PasswdGroups", isReconcilable)
}
func TestReconcilableSSH(t *testing.T) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	expectedError := fmt.Errorf("broken")
	testClient := RpmOstreeClientMock{GetBootedOSImageURLReturns: []GetBootedOSImageURLReturn{}, RunPivotReturns: []error{nil, expectedError}}
	d := Daemon{name: "nodeName", OperatingSystem: machineConfigDaemonOSRHCOS, NodeUpdaterClient: testClient, kubeClient: k8sfake.NewSimpleClientset(), rootMount: "/", bootedOSImageURL: "test"}
	oldMcfg := &mcfgv1.MachineConfig{Spec: mcfgv1.MachineConfigSpec{Config: ignv2_2types.Config{Ignition: ignv2_2types.Ignition{Version: "2.2.0"}}}}
	tempUser1 := ignv2_2types.PasswdUser{Name: "core", SSHAuthorizedKeys: []ignv2_2types.SSHAuthorizedKey{"5678", "abc"}}
	newMcfg := &mcfgv1.MachineConfig{Spec: mcfgv1.MachineConfigSpec{Config: ignv2_2types.Config{Ignition: ignv2_2types.Ignition{Version: "2.2.0"}, Passwd: ignv2_2types.Passwd{Users: []ignv2_2types.PasswdUser{tempUser1}}}}}
	errMsg := d.reconcilable(oldMcfg, newMcfg)
	checkReconcilableResults(t, "SSH", errMsg)
	tempUser2 := ignv2_2types.PasswdUser{Name: "core", SSHAuthorizedKeys: []ignv2_2types.SSHAuthorizedKey{"1234"}}
	oldMcfg.Spec.Config.Passwd.Users = append(oldMcfg.Spec.Config.Passwd.Users, tempUser2)
	tempUser3 := ignv2_2types.PasswdUser{Name: "another user", SSHAuthorizedKeys: []ignv2_2types.SSHAuthorizedKey{"5678"}}
	newMcfg.Spec.Config.Passwd.Users[0] = tempUser3
	errMsg = d.reconcilable(oldMcfg, newMcfg)
	checkIrreconcilableResults(t, "SSH", errMsg)
	tempUser4 := ignv2_2types.PasswdUser{Name: "core", SSHAuthorizedKeys: []ignv2_2types.SSHAuthorizedKey{"5678"}, HomeDir: "somedir"}
	newMcfg.Spec.Config.Passwd.Users[0] = tempUser4
	errMsg = d.reconcilable(oldMcfg, newMcfg)
	checkIrreconcilableResults(t, "SSH", errMsg)
	tempUser5 := ignv2_2types.PasswdUser{Name: "some user", SSHAuthorizedKeys: []ignv2_2types.SSHAuthorizedKey{"5678"}}
	newMcfg.Spec.Config.Passwd.Users = append(newMcfg.Spec.Config.Passwd.Users, tempUser5)
	errMsg = d.reconcilable(oldMcfg, newMcfg)
	checkIrreconcilableResults(t, "SSH", errMsg)
	tempUser6 := ignv2_2types.PasswdUser{Name: "core", SSHAuthorizedKeys: []ignv2_2types.SSHAuthorizedKey{}}
	newMcfg.Spec.Config.Passwd.Users[0] = tempUser6
	newMcfg.Spec.Config.Passwd.Users = newMcfg.Spec.Config.Passwd.Users[:len(newMcfg.Spec.Config.Passwd.Users)-1]
	errMsg = d.reconcilable(oldMcfg, newMcfg)
	checkIrreconcilableResults(t, "SSH", errMsg)
	newMcfg.Spec.Config.Passwd.Users = nil
	errMsg = d.reconcilable(oldMcfg, newMcfg)
	checkReconcilableResults(t, "SSH", errMsg)
}
func TestUpdateSSHKeys(t *testing.T) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	expectedError := fmt.Errorf("broken")
	testClient := RpmOstreeClientMock{GetBootedOSImageURLReturns: []GetBootedOSImageURLReturn{}, RunPivotReturns: []error{nil, expectedError}}
	d := Daemon{name: "nodeName", OperatingSystem: machineConfigDaemonOSRHCOS, NodeUpdaterClient: testClient, kubeClient: k8sfake.NewSimpleClientset(), rootMount: "/", bootedOSImageURL: "test"}
	tempUser := ignv2_2types.PasswdUser{Name: "core", SSHAuthorizedKeys: []ignv2_2types.SSHAuthorizedKey{"1234", "4567"}}
	newMcfg := &mcfgv1.MachineConfig{Spec: mcfgv1.MachineConfigSpec{Config: ignv2_2types.Config{Passwd: ignv2_2types.Passwd{Users: []ignv2_2types.PasswdUser{tempUser}}}}}
	d.atomicSSHKeysWriter = func(user ignv2_2types.PasswdUser, keys string) error {
		return nil
	}
	err := d.updateSSHKeys(newMcfg.Spec.Config.Passwd.Users)
	if err != nil {
		t.Errorf("Expected no error. Got %s.", err)
	}
	newMcfg2 := &mcfgv1.MachineConfig{}
	err = d.updateSSHKeys(newMcfg2.Spec.Config.Passwd.Users)
	if err != nil {
		t.Errorf("Expected no error. Got: %s", err)
	}
}
func TestInvalidIgnConfig(t *testing.T) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	expectedError := fmt.Errorf("broken")
	testClient := RpmOstreeClientMock{GetBootedOSImageURLReturns: []GetBootedOSImageURLReturn{}, RunPivotReturns: []error{nil, expectedError}}
	d := Daemon{name: "nodeName", OperatingSystem: machineConfigDaemonOSRHCOS, NodeUpdaterClient: testClient, kubeClient: k8sfake.NewSimpleClientset(), rootMount: "/", bootedOSImageURL: "test"}
	oldMcfg := &mcfgv1.MachineConfig{Spec: mcfgv1.MachineConfigSpec{Config: ignv2_2types.Config{Ignition: ignv2_2types.Ignition{Version: "2.2.0"}}}}
	tempFileContents := ignv2_2types.FileContents{Source: "data:,hello%20world%0A"}
	tempMode := 420
	newMcfg := &mcfgv1.MachineConfig{Spec: mcfgv1.MachineConfigSpec{Config: ignv2_2types.Config{Ignition: ignv2_2types.Ignition{Version: "2.2.0"}, Storage: ignv2_2types.Storage{Files: []ignv2_2types.File{{Node: ignv2_2types.Node{Path: "home/core/test", Filesystem: "root"}, FileEmbedded1: ignv2_2types.FileEmbedded1{Contents: tempFileContents, Mode: &tempMode}}}}}}}
	err := d.reconcilable(oldMcfg, newMcfg)
	assert.NotNil(t, err, "Expected error. Relative Paths should fail general ignition validation")
	newMcfg.Spec.Config.Storage.Files[0].Node.Path = "/home/core/test"
	err = d.reconcilable(oldMcfg, newMcfg)
	assert.Nil(t, err, "Expected no error. Absolute paths should not fail general ignition validation")
}
func checkReconcilableResults(t *testing.T, key string, reconcilableError error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if reconcilableError != nil {
		t.Errorf("%s values should be reconcilable. Received error: %v", key, reconcilableError)
	}
}
func checkIrreconcilableResults(t *testing.T, key string, reconcilableError error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	if reconcilableError == nil {
		t.Errorf("Different %s values should not be reconcilable.", key)
	}
}
func TestSkipReboot(t *testing.T) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := &Daemon{onceFrom: "test", skipReboot: true}
	require.Nil(t, d.reboot("", 0, nil))
	d = &Daemon{onceFrom: "", skipReboot: true}
	require.NotNil(t, d.reboot("", 0, exec.Command("true")))
}

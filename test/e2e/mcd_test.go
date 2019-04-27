package e2e_test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	mcv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"github.com/openshift/machine-config-operator/pkg/daemon/constants"
	"github.com/openshift/machine-config-operator/test/e2e/framework"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/util/wait"
)

func TestMCDToken(t *testing.T) {
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
	cs := framework.NewClientSet("")
	listOptions := metav1.ListOptions{LabelSelector: labels.SelectorFromSet(labels.Set{"k8s-app": "machine-config-daemon"}).String()}
	mcdList, err := cs.Pods("openshift-machine-config-operator").List(listOptions)
	if err != nil {
		t.Fatalf("%#v", err)
	}
	for _, pod := range mcdList.Items {
		res, err := cs.Pods(pod.Namespace).GetLogs(pod.Name, &v1.PodLogOptions{}).DoRaw()
		if err != nil {
			t.Errorf("%s", err)
		}
		for _, line := range strings.Split(string(res), "\n") {
			if strings.Contains(line, "Unable to rotate token") {
				t.Fatalf("found token rotation failure message: %s", line)
			}
		}
	}
}
func mcLabelForWorkers() map[string]string {
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
	mcLabels := make(map[string]string)
	mcLabels["machineconfiguration.openshift.io/role"] = "worker"
	return mcLabels
}
func createIgnFile(path, content, fs string, mode int) ignv2_2types.File {
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
	return ignv2_2types.File{FileEmbedded1: ignv2_2types.FileEmbedded1{Contents: ignv2_2types.FileContents{Source: content}, Mode: &mode}, Node: ignv2_2types.Node{Filesystem: fs, Path: path}}
}
func createMCToAddFile(name, filename, data, fs string) *mcv1.MachineConfig {
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
	mcName := fmt.Sprintf("%s-%s", name, uuid.NewUUID())
	mcadd := &mcv1.MachineConfig{}
	mcadd.ObjectMeta = metav1.ObjectMeta{Name: mcName, Labels: mcLabelForWorkers()}
	mcadd.Spec = mcv1.MachineConfigSpec{Config: ignv2_2types.Config{Ignition: ignv2_2types.Ignition{Version: "2.2.0"}, Storage: ignv2_2types.Storage{Files: []ignv2_2types.File{createIgnFile(filename, "data:,"+data, fs, 420)}}}}
	return mcadd
}
func TestMCDeployed(t *testing.T) {
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
	cs := framework.NewClientSet("")
	for i := 0; i < 10; i++ {
		mcadd := createMCToAddFile("add-a-file", fmt.Sprintf("/etc/mytestconf%d", i), "test", "root")
		_, err := cs.MachineConfigs().Create(mcadd)
		if err != nil {
			t.Errorf("failed to create machine config %v", err)
		}
		var newMCName string
		if err := wait.PollImmediate(2*time.Second, 5*time.Minute, func() (bool, error) {
			mcp, err := cs.MachineConfigPools().Get("worker", metav1.GetOptions{})
			if err != nil {
				return false, err
			}
			for _, mc := range mcp.Status.Configuration.Source {
				if mc.Name == mcadd.Name {
					newMCName = mcp.Status.Configuration.Name
					return true, nil
				}
			}
			return false, nil
		}); err != nil {
			t.Errorf("machine config hasn't been picked by the pool: %v", err)
		}
		visited := make(map[string]bool)
		if err := wait.Poll(2*time.Second, 10*time.Minute, func() (bool, error) {
			nodes, err := getNodesByRole(cs, "worker")
			if err != nil {
				return false, nil
			}
			for _, node := range nodes {
				if visited[node.Name] {
					continue
				}
				if node.Annotations[constants.CurrentMachineConfigAnnotationKey] == newMCName && node.Annotations[constants.MachineConfigDaemonStateAnnotationKey] == constants.MachineConfigDaemonStateDone {
					visited[node.Name] = true
					if len(visited) == len(nodes) {
						return true, nil
					}
					continue
				}
			}
			return false, nil
		}); err != nil {
			t.Errorf("machine config didn't result in file being on any worker: %v", err)
		}
	}
}
func TestUpdateSSH(t *testing.T) {
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
	cs := framework.NewClientSet("")
	mcName := fmt.Sprintf("sshkeys-worker-%s", uuid.NewUUID())
	mcadd := &mcv1.MachineConfig{}
	mcadd.ObjectMeta = metav1.ObjectMeta{Name: mcName, Labels: mcLabelForWorkers()}
	tempUser := ignv2_2types.PasswdUser{Name: "core", SSHAuthorizedKeys: []ignv2_2types.SSHAuthorizedKey{"1234_test", "abc_test"}}
	mcadd.Spec = mcv1.MachineConfigSpec{Config: ignv2_2types.Config{Ignition: ignv2_2types.Ignition{Version: "2.2.0"}, Passwd: ignv2_2types.Passwd{Users: []ignv2_2types.PasswdUser{tempUser}}}}
	_, err := cs.MachineConfigs().Create(mcadd)
	if err != nil {
		t.Errorf("failed to create machine config %v", err)
	}
	var newMCName string
	if err := wait.PollImmediate(2*time.Second, 5*time.Minute, func() (bool, error) {
		mcp, err := cs.MachineConfigPools().Get("worker", metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		for _, mc := range mcp.Status.Configuration.Source {
			if mc.Name == mcName {
				newMCName = mcp.Status.Configuration.Name
				return true, nil
			}
		}
		return false, nil
	}); err != nil {
		t.Errorf("machine config hasn't been picked by the pool: %v", err)
	}
	visited := make(map[string]bool)
	if err := wait.Poll(2*time.Second, 10*time.Minute, func() (bool, error) {
		nodes, err := getNodesByRole(cs, "worker")
		if err != nil {
			return false, err
		}
		for _, node := range nodes {
			if visited[node.Name] {
				continue
			}
			if node.Annotations[constants.CurrentMachineConfigAnnotationKey] == newMCName && node.Annotations[constants.MachineConfigDaemonStateAnnotationKey] == constants.MachineConfigDaemonStateDone {
				listOptions := metav1.ListOptions{FieldSelector: fields.SelectorFromSet(fields.Set{"spec.nodeName": node.Name}).String()}
				listOptions.LabelSelector = labels.SelectorFromSet(labels.Set{"k8s-app": "machine-config-daemon"}).String()
				mcdList, err := cs.Pods("openshift-machine-config-operator").List(listOptions)
				if err != nil {
					return false, nil
				}
				if len(mcdList.Items) != 1 {
					t.Logf("did not find any mcd pods")
					return false, nil
				}
				mcdName := mcdList.Items[0].ObjectMeta.Name
				found, err := exec.Command("oc", "rsh", "-n", "openshift-machine-config-operator", mcdName, "grep", "1234_test", "/rootfs/home/core/.ssh/authorized_keys").CombinedOutput()
				if err != nil {
					t.Logf("unable to read authorized_keys on daemon: %s got: %s got err: %v", mcdName, found, err)
					return false, nil
				}
				if !strings.Contains(string(found), "1234_test") {
					t.Logf("updated ssh keys not found in authorized_keys, got %s", found)
					return false, nil
				}
				visited[node.Name] = true
				if len(visited) == len(nodes) {
					return true, nil
				}
				continue
			}
		}
		return false, nil
	}); err != nil {
		t.Errorf("machine config didn't result in ssh keys being on any worker: %v", err)
	}
}
func getNodesByRole(cs *framework.ClientSet, role string) ([]v1.Node, error) {
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
	listOptions := metav1.ListOptions{LabelSelector: labels.SelectorFromSet(labels.Set{fmt.Sprintf("node-role.kubernetes.io/%s", role): ""}).String()}
	nodes, err := cs.Nodes().List(listOptions)
	if err != nil {
		return nil, err
	}
	return nodes.Items, nil
}
func TestPoolDegradedOnFailToRender(t *testing.T) {
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
	cs := framework.NewClientSet("")
	mcadd := createMCToAddFile("add-a-file", "/etc/mytestconfs", "test", "")
	mcadd.Spec.Config.Ignition.Version = ""
	_, err := cs.MachineConfigs().Create(mcadd)
	if err != nil {
		t.Errorf("failed to create machine config %v", err)
	}
	if err := wait.PollImmediate(2*time.Second, 5*time.Minute, func() (bool, error) {
		mcp, err := cs.MachineConfigPools().Get("worker", metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		if mcv1.IsMachineConfigPoolConditionTrue(mcp.Status.Conditions, mcv1.MachineConfigPoolDegraded) {
			return true, nil
		}
		return false, nil
	}); err != nil {
		t.Errorf("machine config pool never switched to Degraded on failure to render: %v", err)
	}
	if err := cs.MachineConfigs().Delete(mcadd.Name, &metav1.DeleteOptions{}); err != nil {
		t.Error(err)
	}
	if err := wait.PollImmediate(2*time.Second, 5*time.Minute, func() (bool, error) {
		mcp, err := cs.MachineConfigPools().Get("worker", metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		if mcv1.IsMachineConfigPoolConditionFalse(mcp.Status.Conditions, mcv1.MachineConfigPoolDegraded) {
			return true, nil
		}
		return false, nil
	}); err != nil {
		t.Errorf("machine config pool never switched back to Degraded=False: %v", err)
	}
}
func TestReconcileAfterBadMC(t *testing.T) {
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
	cs := framework.NewClientSet("")
	mcadd := createMCToAddFile("add-a-file", "/etc/mytestconfs", "test", "")
	mcp, err := cs.MachineConfigPools().Get("worker", metav1.GetOptions{})
	if err != nil {
		t.Error(err)
	}
	workerOldMc := mcp.Status.Configuration.Name
	_, err = cs.MachineConfigs().Create(mcadd)
	if err != nil {
		t.Errorf("failed to create machine config %v", err)
	}
	var newMCName string
	if err := wait.PollImmediate(2*time.Second, 5*time.Minute, func() (bool, error) {
		mcp, err := cs.MachineConfigPools().Get("worker", metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		for _, mc := range mcp.Status.Configuration.Source {
			if mc.Name == mcadd.Name {
				newMCName = mcp.Status.Configuration.Name
				return true, nil
			}
		}
		return false, nil
	}); err != nil {
		t.Errorf("machine config hasn't been picked by the pool: %v", err)
	}
	if err := wait.Poll(2*time.Second, 5*time.Minute, func() (bool, error) {
		nodes, err := getNodesByRole(cs, "worker")
		if err != nil {
			return false, err
		}
		for _, node := range nodes {
			if node.Annotations[constants.DesiredMachineConfigAnnotationKey] == newMCName && node.Annotations[constants.MachineConfigDaemonStateAnnotationKey] != constants.MachineConfigDaemonStateDone {
				if node.Annotations[constants.MachineConfigDaemonReasonAnnotationKey] != "" {
					return true, nil
				}
			}
		}
		return false, nil
	}); err != nil {
		t.Errorf("machine config hasn't been picked by any MCD: %v", err)
	}
	if err := wait.Poll(2*time.Second, 5*time.Minute, func() (bool, error) {
		mcp, err := cs.MachineConfigPools().Get("worker", metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		if mcp.Status.UnavailableMachineCount == 1 {
			return true, nil
		}
		return false, nil
	}); err != nil {
		t.Errorf("MCP isn't reporting unavailable with a bad MC: %v", err)
	}
	if err := cs.MachineConfigs().Delete(mcadd.Name, &metav1.DeleteOptions{}); err != nil {
		t.Error(err)
	}
	if err := wait.PollImmediate(2*time.Second, 5*time.Minute, func() (bool, error) {
		mcp, err := cs.MachineConfigPools().Get("worker", metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		if mcp.Status.Configuration.Name == workerOldMc {
			return true, nil
		}
		return false, nil
	}); err != nil {
		t.Errorf("old machine config hasn't been picked by the pool: %v", err)
	}
	visited := make(map[string]bool)
	if err := wait.Poll(2*time.Second, 10*time.Minute, func() (bool, error) {
		nodes, err := getNodesByRole(cs, "worker")
		if err != nil {
			return false, err
		}
		mcp, err = cs.MachineConfigPools().Get("worker", metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		for _, node := range nodes {
			if node.Annotations[constants.CurrentMachineConfigAnnotationKey] == workerOldMc && node.Annotations[constants.DesiredMachineConfigAnnotationKey] == workerOldMc && node.Annotations[constants.MachineConfigDaemonStateAnnotationKey] == constants.MachineConfigDaemonStateDone {
				visited[node.Name] = true
				if len(visited) == len(nodes) {
					if mcp.Status.UnavailableMachineCount == 0 && mcp.Status.ReadyMachineCount == int32(len(nodes)) && mcp.Status.UpdatedMachineCount == int32(len(nodes)) {
						return true, nil
					}
				}
				continue
			}
		}
		return false, nil
	}); err != nil {
		t.Errorf("machine config didn't roll back on any worker: %v", err)
	}
}

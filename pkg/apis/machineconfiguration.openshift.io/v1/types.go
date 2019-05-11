package v1

import (
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	kubeletconfigv1beta1 "k8s.io/kubelet/config/v1beta1"
)

type MCOConfig struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec				MCOConfigSpec	`json:"spec"`
}
type MCOConfigSpec struct{}
type MCOConfigList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`
	Items			[]MCOConfig	`json:"items"`
}
type ControllerConfig struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec				ControllerConfigSpec	`json:"spec"`
	Status				ControllerConfigStatus	`json:"status"`
}
type ControllerConfigSpec struct {
	ClusterDNSIP		string					`json:"clusterDNSIP"`
	CloudProviderConfig	string					`json:"cloudProviderConfig"`
	Platform			string					`json:"platform"`
	EtcdDiscoveryDomain	string					`json:"etcdDiscoveryDomain"`
	EtcdCAData			[]byte					`json:"etcdCAData"`
	EtcdMetricCAData	[]byte					`json:"etcdMetricCAData"`
	RootCAData			[]byte					`json:"rootCAData"`
	PullSecret			*corev1.ObjectReference	`json:"pullSecret,omitempty"`
	Images				map[string]string		`json:"images"`
	OSImageURL			string					`json:"osImageURL"`
}
type ControllerConfigStatus struct {
	ObservedGeneration	int64								`json:"observedGeneration,omitempty"`
	Conditions			[]ControllerConfigStatusCondition	`json:"conditions"`
}
type ControllerConfigStatusCondition struct {
	Type				ControllerConfigStatusConditionType	`json:"type"`
	Status				corev1.ConditionStatus				`json:"status"`
	LastTransitionTime	metav1.Time							`json:"lastTransitionTime"`
	Reason				string								`json:"reason,omitempty"`
	Message				string								`json:"message,omitempty"`
}
type ControllerConfigStatusConditionType string

const (
	TemplateContollerRunning	ControllerConfigStatusConditionType	= "TemplateContollerRunning"
	TemplateContollerCompleted	ControllerConfigStatusConditionType	= "TemplateContollerCompleted"
	TemplateContollerFailing	ControllerConfigStatusConditionType	= "TemplateContollerFailing"
)

type ControllerConfigList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`
	Items			[]ControllerConfig	`json:"items"`
}
type MachineConfig struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec				MachineConfigSpec	`json:"spec"`
}
type MachineConfigSpec struct {
	OSImageURL	string				`json:"osImageURL"`
	Config		ignv2_2types.Config	`json:"config"`
}
type MachineConfigList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`
	Items			[]MachineConfig	`json:"items"`
}
type MachineConfigPool struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec				MachineConfigPoolSpec	`json:"spec"`
	Status				MachineConfigPoolStatus	`json:"status"`
}
type MachineConfigPoolSpec struct {
	MachineConfigSelector	*metav1.LabelSelector	`json:"machineConfigSelector,omitempty"`
	NodeSelector			*metav1.LabelSelector	`json:"nodeSelector,omitempty"`
	Paused					bool					`json:"paused"`
	MaxUnavailable			*intstr.IntOrString		`json:"maxUnavailable"`
}
type MachineConfigPoolStatus struct {
	ObservedGeneration		int64									`json:"observedGeneration,omitempty"`
	Configuration			MachineConfigPoolStatusConfiguration	`json:"configuration"`
	MachineCount			int32									`json:"machineCount"`
	UpdatedMachineCount		int32									`json:"updatedMachineCount"`
	ReadyMachineCount		int32									`json:"readyMachineCount"`
	UnavailableMachineCount	int32									`json:"unavailableMachineCount"`
	Conditions				[]MachineConfigPoolCondition			`json:"conditions"`
}
type MachineConfigPoolStatusConfiguration struct {
	corev1.ObjectReference
	Source	[]corev1.ObjectReference	`json:"source,omitempty"`
}
type MachineConfigPoolCondition struct {
	Type				MachineConfigPoolConditionType	`json:"type"`
	Status				corev1.ConditionStatus			`json:"status"`
	LastTransitionTime	metav1.Time						`json:"lastTransitionTime"`
	Reason				string							`json:"reason"`
	Message				string							`json:"message"`
}
type MachineConfigPoolConditionType string

const (
	MachineConfigPoolUpdated	MachineConfigPoolConditionType	= "Updated"
	MachineConfigPoolUpdating	MachineConfigPoolConditionType	= "Updating"
)

type MachineConfigPoolList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`
	Items			[]MachineConfigPool	`json:"items"`
}
type KubeletConfig struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec				KubeletConfigSpec	`json:"spec,omitempty"`
	Status				KubeletConfigStatus	`json:"status,omitempty"`
}
type KubeletConfigSpec struct {
	MachineConfigPoolSelector	*metav1.LabelSelector						`json:"machineConfigPoolSelector,omitempty"`
	KubeletConfig				*kubeletconfigv1beta1.KubeletConfiguration	`json:"kubeletConfig,omitempty"`
}
type KubeletConfigStatus struct {
	ObservedGeneration	int64						`json:"observedGeneration,omitempty"`
	Conditions			[]KubeletConfigCondition	`json:"conditions"`
}
type KubeletConfigCondition struct {
	Type				KubeletConfigStatusConditionType	`json:"type"`
	Status				corev1.ConditionStatus				`json:"status"`
	LastTransitionTime	metav1.Time							`json:"lastTransitionTime"`
	Reason				string								`json:"reason,omitempty"`
	Message				string								`json:"message,omitempty"`
}
type KubeletConfigStatusConditionType string

const (
	KubeletConfigSuccess	KubeletConfigStatusConditionType	= "Success"
	KubeletConfigFailure	KubeletConfigStatusConditionType	= "Failure"
)

type KubeletConfigList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`
	Items			[]KubeletConfig	`json:"items"`
}
type ContainerRuntimeConfig struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec				ContainerRuntimeConfigSpec		`json:"spec,omitempty"`
	Status				ContainerRuntimeConfigStatus	`json:"status,omitempty"`
}
type ContainerRuntimeConfigSpec struct {
	MachineConfigPoolSelector	*metav1.LabelSelector			`json:"machineConfigPoolSelector,omitempty"`
	ContainerRuntimeConfig		*ContainerRuntimeConfiguration	`json:"containerRuntimeConfig,omitempty"`
}
type ContainerRuntimeConfiguration struct {
	PidsLimit	int64				`json:"pidsLimit,omitempty"`
	LogLevel	string				`json:"logLevel,omitempty"`
	LogSizeMax	resource.Quantity	`json:"logSizeMax,omitempty"`
	OverlaySize	resource.Quantity	`json:"overlaySize,omitempty"`
}
type ContainerRuntimeConfigStatus struct {
	ObservedGeneration	int64								`json:"observedGeneration,omitempty"`
	Conditions			[]ContainerRuntimeConfigCondition	`json:"conditions"`
}
type ContainerRuntimeConfigCondition struct {
	Type				ContainerRuntimeConfigStatusConditionType	`json:"type"`
	Status				corev1.ConditionStatus						`json:"status"`
	LastTransitionTime	metav1.Time									`json:"lastTransitionTime"`
	Reason				string										`json:"reason,omitempty"`
	Message				string										`json:"message,omitempty"`
}
type ContainerRuntimeConfigStatusConditionType string

const (
	ContainerRuntimeConfigSuccess	ContainerRuntimeConfigStatusConditionType	= "Success"
	ContainerRuntimeConfigFailure	ContainerRuntimeConfigStatusConditionType	= "Failure"
)

type ContainerRuntimeConfigList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata"`
	Items			[]ContainerRuntimeConfig	`json:"items"`
}

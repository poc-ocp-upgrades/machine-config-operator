package v1

import (
	"github.com/openshift/machine-config-operator/pkg/apis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	GroupName = apis.GroupName
)

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

func Resource(resource string) schema.GroupResource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder		runtime.SchemeBuilder
	localSchemeBuilder	= &SchemeBuilder
	AddToScheme		= localSchemeBuilder.AddToScheme
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(addKnownTypes)
}
func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(SchemeGroupVersion, &MCOConfig{}, &ContainerRuntimeConfig{}, &ContainerRuntimeConfigList{}, &ControllerConfig{}, &ControllerConfigList{}, &KubeletConfig{}, &KubeletConfigList{}, &MachineConfig{}, &MachineConfigList{}, &MachineConfigPool{}, &MachineConfigPoolList{})
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

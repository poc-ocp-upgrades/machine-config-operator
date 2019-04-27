package v1

import (
	internalinterfaces "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

type Interface interface {
	ContainerRuntimeConfigs() ContainerRuntimeConfigInformer
	ControllerConfigs() ControllerConfigInformer
	KubeletConfigs() KubeletConfigInformer
	MCOConfigs() MCOConfigInformer
	MachineConfigs() MachineConfigInformer
	MachineConfigPools() MachineConfigPoolInformer
}
type version struct {
	factory			internalinterfaces.SharedInformerFactory
	namespace		string
	tweakListOptions	internalinterfaces.TweakListOptionsFunc
}

func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
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
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}
func (v *version) ContainerRuntimeConfigs() ContainerRuntimeConfigInformer {
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
	return &containerRuntimeConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
func (v *version) ControllerConfigs() ControllerConfigInformer {
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
	return &controllerConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
func (v *version) KubeletConfigs() KubeletConfigInformer {
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
	return &kubeletConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
func (v *version) MCOConfigs() MCOConfigInformer {
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
	return &mCOConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
func (v *version) MachineConfigs() MachineConfigInformer {
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
	return &machineConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
func (v *version) MachineConfigPools() MachineConfigPoolInformer {
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
	return &machineConfigPoolInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

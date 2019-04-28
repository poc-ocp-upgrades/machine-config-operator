package v1

import (
	ignv2_2 "github.com/coreos/ignition/config/v2_2"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *MachineConfig) DeepCopyInto(out *MachineConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}
func (in *MachineConfig) DeepCopy() *MachineConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(MachineConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *MachineConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return in.DeepCopy()
}
func (in *MachineConfigSpec) DeepCopyInto(out *MachineConfigSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.Config = deepCopyIgnConfig(in.Config)
	return
}
func deepCopyIgnConfig(in ignv2_2types.Config) ignv2_2types.Config {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var out ignv2_2types.Config
	out.Ignition.Version = in.Ignition.Version
	return ignv2_2.Append(out, in)
}
func (in *MachineConfigSpec) DeepCopy() *MachineConfigSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(MachineConfigSpec)
	in.DeepCopyInto(out)
	return out
}

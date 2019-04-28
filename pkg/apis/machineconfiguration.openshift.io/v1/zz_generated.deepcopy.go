package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	v1beta1 "k8s.io/kubelet/config/v1beta1"
)

func (in *ContainerRuntimeConfig) DeepCopyInto(out *ContainerRuntimeConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *ContainerRuntimeConfig) DeepCopy() *ContainerRuntimeConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *ContainerRuntimeConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ContainerRuntimeConfigCondition) DeepCopyInto(out *ContainerRuntimeConfigCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}
func (in *ContainerRuntimeConfigCondition) DeepCopy() *ContainerRuntimeConfigCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfigCondition)
	in.DeepCopyInto(out)
	return out
}
func (in *ContainerRuntimeConfigList) DeepCopyInto(out *ContainerRuntimeConfigList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContainerRuntimeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ContainerRuntimeConfigList) DeepCopy() *ContainerRuntimeConfigList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfigList)
	in.DeepCopyInto(out)
	return out
}
func (in *ContainerRuntimeConfigList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ContainerRuntimeConfigSpec) DeepCopyInto(out *ContainerRuntimeConfigSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.MachineConfigPoolSelector != nil {
		in, out := &in.MachineConfigPoolSelector, &out.MachineConfigPoolSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerRuntimeConfig != nil {
		in, out := &in.ContainerRuntimeConfig, &out.ContainerRuntimeConfig
		*out = new(ContainerRuntimeConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *ContainerRuntimeConfigSpec) DeepCopy() *ContainerRuntimeConfigSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfigSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ContainerRuntimeConfigStatus) DeepCopyInto(out *ContainerRuntimeConfigStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ContainerRuntimeConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ContainerRuntimeConfigStatus) DeepCopy() *ContainerRuntimeConfigStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfigStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *ContainerRuntimeConfiguration) DeepCopyInto(out *ContainerRuntimeConfiguration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.LogSizeMax = in.LogSizeMax.DeepCopy()
	out.OverlaySize = in.OverlaySize.DeepCopy()
	return
}
func (in *ContainerRuntimeConfiguration) DeepCopy() *ContainerRuntimeConfiguration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfiguration)
	in.DeepCopyInto(out)
	return out
}
func (in *ControllerConfig) DeepCopyInto(out *ControllerConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *ControllerConfig) DeepCopy() *ControllerConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ControllerConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *ControllerConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ControllerConfigList) DeepCopyInto(out *ControllerConfigList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ControllerConfigList) DeepCopy() *ControllerConfigList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ControllerConfigList)
	in.DeepCopyInto(out)
	return out
}
func (in *ControllerConfigList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ControllerConfigSpec) DeepCopyInto(out *ControllerConfigSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.EtcdCAData != nil {
		in, out := &in.EtcdCAData, &out.EtcdCAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.EtcdMetricCAData != nil {
		in, out := &in.EtcdMetricCAData, &out.EtcdMetricCAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.RootCAData != nil {
		in, out := &in.RootCAData, &out.RootCAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}
func (in *ControllerConfigSpec) DeepCopy() *ControllerConfigSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ControllerConfigSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ControllerConfigStatus) DeepCopyInto(out *ControllerConfigStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ControllerConfigStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ControllerConfigStatus) DeepCopy() *ControllerConfigStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ControllerConfigStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *ControllerConfigStatusCondition) DeepCopyInto(out *ControllerConfigStatusCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}
func (in *ControllerConfigStatusCondition) DeepCopy() *ControllerConfigStatusCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ControllerConfigStatusCondition)
	in.DeepCopyInto(out)
	return out
}
func (in *KubeletConfig) DeepCopyInto(out *KubeletConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *KubeletConfig) DeepCopy() *KubeletConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(KubeletConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *KubeletConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *KubeletConfigCondition) DeepCopyInto(out *KubeletConfigCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}
func (in *KubeletConfigCondition) DeepCopy() *KubeletConfigCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(KubeletConfigCondition)
	in.DeepCopyInto(out)
	return out
}
func (in *KubeletConfigList) DeepCopyInto(out *KubeletConfigList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeletConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *KubeletConfigList) DeepCopy() *KubeletConfigList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(KubeletConfigList)
	in.DeepCopyInto(out)
	return out
}
func (in *KubeletConfigList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *KubeletConfigSpec) DeepCopyInto(out *KubeletConfigSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.MachineConfigPoolSelector != nil {
		in, out := &in.MachineConfigPoolSelector, &out.MachineConfigPoolSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeletConfig != nil {
		in, out := &in.KubeletConfig, &out.KubeletConfig
		*out = new(v1beta1.KubeletConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *KubeletConfigSpec) DeepCopy() *KubeletConfigSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(KubeletConfigSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *KubeletConfigStatus) DeepCopyInto(out *KubeletConfigStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]KubeletConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *KubeletConfigStatus) DeepCopy() *KubeletConfigStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(KubeletConfigStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *MCOConfig) DeepCopyInto(out *MCOConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}
func (in *MCOConfig) DeepCopy() *MCOConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(MCOConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *MCOConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *MCOConfigList) DeepCopyInto(out *MCOConfigList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MCOConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *MCOConfigList) DeepCopy() *MCOConfigList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(MCOConfigList)
	in.DeepCopyInto(out)
	return out
}
func (in *MCOConfigList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *MCOConfigSpec) DeepCopyInto(out *MCOConfigSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *MCOConfigSpec) DeepCopy() *MCOConfigSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(MCOConfigSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *MachineConfigList) DeepCopyInto(out *MachineConfigList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *MachineConfigList) DeepCopy() *MachineConfigList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(MachineConfigList)
	in.DeepCopyInto(out)
	return out
}
func (in *MachineConfigList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *MachineConfigPool) DeepCopyInto(out *MachineConfigPool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *MachineConfigPool) DeepCopy() *MachineConfigPool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(MachineConfigPool)
	in.DeepCopyInto(out)
	return out
}
func (in *MachineConfigPool) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *MachineConfigPoolCondition) DeepCopyInto(out *MachineConfigPoolCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}
func (in *MachineConfigPoolCondition) DeepCopy() *MachineConfigPoolCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolCondition)
	in.DeepCopyInto(out)
	return out
}
func (in *MachineConfigPoolList) DeepCopyInto(out *MachineConfigPoolList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineConfigPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *MachineConfigPoolList) DeepCopy() *MachineConfigPoolList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolList)
	in.DeepCopyInto(out)
	return out
}
func (in *MachineConfigPoolList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *MachineConfigPoolSpec) DeepCopyInto(out *MachineConfigPoolSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.MachineConfigSelector != nil {
		in, out := &in.MachineConfigSelector, &out.MachineConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}
func (in *MachineConfigPoolSpec) DeepCopy() *MachineConfigPoolSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *MachineConfigPoolStatus) DeepCopyInto(out *MachineConfigPoolStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MachineConfigPoolCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *MachineConfigPoolStatus) DeepCopy() *MachineConfigPoolStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *MachineConfigPoolStatusConfiguration) DeepCopyInto(out *MachineConfigPoolStatusConfiguration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.ObjectReference = in.ObjectReference
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *MachineConfigPoolStatusConfiguration) DeepCopy() *MachineConfigPoolStatusConfiguration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(MachineConfigPoolStatusConfiguration)
	in.DeepCopyInto(out)
	return out
}

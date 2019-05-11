package resourcemerge

import (
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"k8s.io/apimachinery/pkg/api/equality"
)

func EnsureMachineConfig(modified *bool, existing *mcfgv1.MachineConfig, required mcfgv1.MachineConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	EnsureObjectMeta(modified, &existing.ObjectMeta, required.ObjectMeta)
	ensureMachineConfigSpec(modified, &existing.Spec, required.Spec)
}
func EnsureControllerConfig(modified *bool, existing *mcfgv1.ControllerConfig, required mcfgv1.ControllerConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	EnsureObjectMeta(modified, &existing.ObjectMeta, required.ObjectMeta)
	ensureControllerConfigSpec(modified, &existing.Spec, required.Spec)
}
func EnsureMachineConfigPool(modified *bool, existing *mcfgv1.MachineConfigPool, required mcfgv1.MachineConfigPool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	EnsureObjectMeta(modified, &existing.ObjectMeta, required.ObjectMeta)
	if existing.Spec.MachineConfigSelector == nil {
		*modified = true
		existing.Spec.MachineConfigSelector = required.Spec.MachineConfigSelector
	}
	if !equality.Semantic.DeepEqual(existing.Spec.MachineConfigSelector, required.Spec.MachineConfigSelector) {
		*modified = true
		existing.Spec.MachineConfigSelector = required.Spec.MachineConfigSelector
	}
	if existing.Spec.NodeSelector == nil {
		*modified = true
		existing.Spec.NodeSelector = required.Spec.NodeSelector
	}
	if !equality.Semantic.DeepEqual(existing.Spec.NodeSelector, required.Spec.NodeSelector) {
		*modified = true
		existing.Spec.NodeSelector = required.Spec.NodeSelector
	}
}
func ensureMachineConfigSpec(modified *bool, existing *mcfgv1.MachineConfigSpec, required mcfgv1.MachineConfigSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setStringIfSet(modified, &existing.OSImageURL, required.OSImageURL)
	if !equality.Semantic.DeepEqual(existing.Config, required.Config) {
		*modified = true
		(*existing).Config = required.Config
	}
}
func ensureControllerConfigSpec(modified *bool, existing *mcfgv1.ControllerConfigSpec, required mcfgv1.ControllerConfigSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setStringIfSet(modified, &existing.ClusterDNSIP, required.ClusterDNSIP)
	setStringIfSet(modified, &existing.CloudProviderConfig, required.CloudProviderConfig)
	setStringIfSet(modified, &existing.Platform, required.Platform)
	setStringIfSet(modified, &existing.EtcdDiscoveryDomain, required.EtcdDiscoveryDomain)
	setStringIfSet(modified, &existing.OSImageURL, required.OSImageURL)
	setBytesIfSet(modified, &existing.EtcdCAData, required.EtcdCAData)
	setBytesIfSet(modified, &existing.EtcdMetricCAData, required.EtcdMetricCAData)
	setBytesIfSet(modified, &existing.RootCAData, required.RootCAData)
	if required.PullSecret != nil && !equality.Semantic.DeepEqual(existing.PullSecret, required.PullSecret) {
		existing.PullSecret = required.PullSecret
		*modified = true
	}
	mergeMap(modified, &existing.Images, required.Images)
}

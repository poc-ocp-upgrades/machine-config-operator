package fake

import (
	v1 "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/typed/machineconfiguration.openshift.io/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMachineconfigurationV1 struct{ *testing.Fake }

func (c *FakeMachineconfigurationV1) ContainerRuntimeConfigs() v1.ContainerRuntimeConfigInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeContainerRuntimeConfigs{c}
}
func (c *FakeMachineconfigurationV1) ControllerConfigs() v1.ControllerConfigInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeControllerConfigs{c}
}
func (c *FakeMachineconfigurationV1) KubeletConfigs() v1.KubeletConfigInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeKubeletConfigs{c}
}
func (c *FakeMachineconfigurationV1) MCOConfigs(namespace string) v1.MCOConfigInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeMCOConfigs{c, namespace}
}
func (c *FakeMachineconfigurationV1) MachineConfigs() v1.MachineConfigInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeMachineConfigs{c}
}
func (c *FakeMachineconfigurationV1) MachineConfigPools() v1.MachineConfigPoolInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeMachineConfigPools{c}
}
func (c *FakeMachineconfigurationV1) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ret *rest.RESTClient
	return ret
}

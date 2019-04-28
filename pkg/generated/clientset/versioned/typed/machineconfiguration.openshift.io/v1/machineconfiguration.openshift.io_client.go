package v1

import (
	v1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type MachineconfigurationV1Interface interface {
	RESTClient() rest.Interface
	ContainerRuntimeConfigsGetter
	ControllerConfigsGetter
	KubeletConfigsGetter
	MCOConfigsGetter
	MachineConfigsGetter
	MachineConfigPoolsGetter
}
type MachineconfigurationV1Client struct{ restClient rest.Interface }

func (c *MachineconfigurationV1Client) ContainerRuntimeConfigs() ContainerRuntimeConfigInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newContainerRuntimeConfigs(c)
}
func (c *MachineconfigurationV1Client) ControllerConfigs() ControllerConfigInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newControllerConfigs(c)
}
func (c *MachineconfigurationV1Client) KubeletConfigs() KubeletConfigInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newKubeletConfigs(c)
}
func (c *MachineconfigurationV1Client) MCOConfigs(namespace string) MCOConfigInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newMCOConfigs(c, namespace)
}
func (c *MachineconfigurationV1Client) MachineConfigs() MachineConfigInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newMachineConfigs(c)
}
func (c *MachineconfigurationV1Client) MachineConfigPools() MachineConfigPoolInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newMachineConfigPools(c)
}
func NewForConfig(c *rest.Config) (*MachineconfigurationV1Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MachineconfigurationV1Client{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *MachineconfigurationV1Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
func New(c rest.Interface) *MachineconfigurationV1Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &MachineconfigurationV1Client{c}
}
func setConfigDefaults(config *rest.Config) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	return nil
}
func (c *MachineconfigurationV1Client) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c == nil {
		return nil
	}
	return c.restClient
}

package v1

import (
	"time"
	v1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	scheme "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type ControllerConfigsGetter interface {
	ControllerConfigs() ControllerConfigInterface
}
type ControllerConfigInterface interface {
	Create(*v1.ControllerConfig) (*v1.ControllerConfig, error)
	Update(*v1.ControllerConfig) (*v1.ControllerConfig, error)
	UpdateStatus(*v1.ControllerConfig) (*v1.ControllerConfig, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ControllerConfig, error)
	List(opts metav1.ListOptions) (*v1.ControllerConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ControllerConfig, err error)
	ControllerConfigExpansion
}
type controllerConfigs struct{ client rest.Interface }

func newControllerConfigs(c *MachineconfigurationV1Client) *controllerConfigs {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &controllerConfigs{client: c.RESTClient()}
}
func (c *controllerConfigs) Get(name string, options metav1.GetOptions) (result *v1.ControllerConfig, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.ControllerConfig{}
	err = c.client.Get().Resource("controllerconfigs").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *controllerConfigs) List(opts metav1.ListOptions) (result *v1.ControllerConfigList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ControllerConfigList{}
	err = c.client.Get().Resource("controllerconfigs").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *controllerConfigs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("controllerconfigs").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *controllerConfigs) Create(controllerConfig *v1.ControllerConfig) (result *v1.ControllerConfig, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.ControllerConfig{}
	err = c.client.Post().Resource("controllerconfigs").Body(controllerConfig).Do().Into(result)
	return
}
func (c *controllerConfigs) Update(controllerConfig *v1.ControllerConfig) (result *v1.ControllerConfig, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.ControllerConfig{}
	err = c.client.Put().Resource("controllerconfigs").Name(controllerConfig.Name).Body(controllerConfig).Do().Into(result)
	return
}
func (c *controllerConfigs) UpdateStatus(controllerConfig *v1.ControllerConfig) (result *v1.ControllerConfig, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.ControllerConfig{}
	err = c.client.Put().Resource("controllerconfigs").Name(controllerConfig.Name).SubResource("status").Body(controllerConfig).Do().Into(result)
	return
}
func (c *controllerConfigs) Delete(name string, options *metav1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("controllerconfigs").Name(name).Body(options).Do().Error()
}
func (c *controllerConfigs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("controllerconfigs").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *controllerConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ControllerConfig, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.ControllerConfig{}
	err = c.client.Patch(pt).Resource("controllerconfigs").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}

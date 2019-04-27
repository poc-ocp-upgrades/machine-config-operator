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

type MCOConfigsGetter interface {
	MCOConfigs(namespace string) MCOConfigInterface
}
type MCOConfigInterface interface {
	Create(*v1.MCOConfig) (*v1.MCOConfig, error)
	Update(*v1.MCOConfig) (*v1.MCOConfig, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.MCOConfig, error)
	List(opts metav1.ListOptions) (*v1.MCOConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MCOConfig, err error)
	MCOConfigExpansion
}
type mCOConfigs struct {
	client	rest.Interface
	ns	string
}

func newMCOConfigs(c *MachineconfigurationV1Client, namespace string) *mCOConfigs {
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
	return &mCOConfigs{client: c.RESTClient(), ns: namespace}
}
func (c *mCOConfigs) Get(name string, options metav1.GetOptions) (result *v1.MCOConfig, err error) {
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
	result = &v1.MCOConfig{}
	err = c.client.Get().Namespace(c.ns).Resource("mcoconfigs").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *mCOConfigs) List(opts metav1.ListOptions) (result *v1.MCOConfigList, err error) {
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
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MCOConfigList{}
	err = c.client.Get().Namespace(c.ns).Resource("mcoconfigs").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *mCOConfigs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
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
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("mcoconfigs").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *mCOConfigs) Create(mCOConfig *v1.MCOConfig) (result *v1.MCOConfig, err error) {
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
	result = &v1.MCOConfig{}
	err = c.client.Post().Namespace(c.ns).Resource("mcoconfigs").Body(mCOConfig).Do().Into(result)
	return
}
func (c *mCOConfigs) Update(mCOConfig *v1.MCOConfig) (result *v1.MCOConfig, err error) {
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
	result = &v1.MCOConfig{}
	err = c.client.Put().Namespace(c.ns).Resource("mcoconfigs").Name(mCOConfig.Name).Body(mCOConfig).Do().Into(result)
	return
}
func (c *mCOConfigs) Delete(name string, options *metav1.DeleteOptions) error {
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
	return c.client.Delete().Namespace(c.ns).Resource("mcoconfigs").Name(name).Body(options).Do().Error()
}
func (c *mCOConfigs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
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
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Namespace(c.ns).Resource("mcoconfigs").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *mCOConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MCOConfig, err error) {
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
	result = &v1.MCOConfig{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("mcoconfigs").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}

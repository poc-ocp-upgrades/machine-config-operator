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

type KubeletConfigsGetter interface{ KubeletConfigs() KubeletConfigInterface }
type KubeletConfigInterface interface {
	Create(*v1.KubeletConfig) (*v1.KubeletConfig, error)
	Update(*v1.KubeletConfig) (*v1.KubeletConfig, error)
	UpdateStatus(*v1.KubeletConfig) (*v1.KubeletConfig, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.KubeletConfig, error)
	List(opts metav1.ListOptions) (*v1.KubeletConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KubeletConfig, err error)
	KubeletConfigExpansion
}
type kubeletConfigs struct{ client rest.Interface }

func newKubeletConfigs(c *MachineconfigurationV1Client) *kubeletConfigs {
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
	return &kubeletConfigs{client: c.RESTClient()}
}
func (c *kubeletConfigs) Get(name string, options metav1.GetOptions) (result *v1.KubeletConfig, err error) {
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
	result = &v1.KubeletConfig{}
	err = c.client.Get().Resource("kubeletconfigs").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *kubeletConfigs) List(opts metav1.ListOptions) (result *v1.KubeletConfigList, err error) {
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
	result = &v1.KubeletConfigList{}
	err = c.client.Get().Resource("kubeletconfigs").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *kubeletConfigs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
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
	return c.client.Get().Resource("kubeletconfigs").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *kubeletConfigs) Create(kubeletConfig *v1.KubeletConfig) (result *v1.KubeletConfig, err error) {
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
	result = &v1.KubeletConfig{}
	err = c.client.Post().Resource("kubeletconfigs").Body(kubeletConfig).Do().Into(result)
	return
}
func (c *kubeletConfigs) Update(kubeletConfig *v1.KubeletConfig) (result *v1.KubeletConfig, err error) {
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
	result = &v1.KubeletConfig{}
	err = c.client.Put().Resource("kubeletconfigs").Name(kubeletConfig.Name).Body(kubeletConfig).Do().Into(result)
	return
}
func (c *kubeletConfigs) UpdateStatus(kubeletConfig *v1.KubeletConfig) (result *v1.KubeletConfig, err error) {
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
	result = &v1.KubeletConfig{}
	err = c.client.Put().Resource("kubeletconfigs").Name(kubeletConfig.Name).SubResource("status").Body(kubeletConfig).Do().Into(result)
	return
}
func (c *kubeletConfigs) Delete(name string, options *metav1.DeleteOptions) error {
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
	return c.client.Delete().Resource("kubeletconfigs").Name(name).Body(options).Do().Error()
}
func (c *kubeletConfigs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
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
	return c.client.Delete().Resource("kubeletconfigs").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *kubeletConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KubeletConfig, err error) {
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
	result = &v1.KubeletConfig{}
	err = c.client.Patch(pt).Resource("kubeletconfigs").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}

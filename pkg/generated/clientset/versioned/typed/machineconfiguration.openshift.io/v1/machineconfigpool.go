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

type MachineConfigPoolsGetter interface {
	MachineConfigPools() MachineConfigPoolInterface
}
type MachineConfigPoolInterface interface {
	Create(*v1.MachineConfigPool) (*v1.MachineConfigPool, error)
	Update(*v1.MachineConfigPool) (*v1.MachineConfigPool, error)
	UpdateStatus(*v1.MachineConfigPool) (*v1.MachineConfigPool, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.MachineConfigPool, error)
	List(opts metav1.ListOptions) (*v1.MachineConfigPoolList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MachineConfigPool, err error)
	MachineConfigPoolExpansion
}
type machineConfigPools struct{ client rest.Interface }

func newMachineConfigPools(c *MachineconfigurationV1Client) *machineConfigPools {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &machineConfigPools{client: c.RESTClient()}
}
func (c *machineConfigPools) Get(name string, options metav1.GetOptions) (result *v1.MachineConfigPool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.MachineConfigPool{}
	err = c.client.Get().Resource("machineconfigpools").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *machineConfigPools) List(opts metav1.ListOptions) (result *v1.MachineConfigPoolList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MachineConfigPoolList{}
	err = c.client.Get().Resource("machineconfigpools").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *machineConfigPools) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("machineconfigpools").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *machineConfigPools) Create(machineConfigPool *v1.MachineConfigPool) (result *v1.MachineConfigPool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.MachineConfigPool{}
	err = c.client.Post().Resource("machineconfigpools").Body(machineConfigPool).Do().Into(result)
	return
}
func (c *machineConfigPools) Update(machineConfigPool *v1.MachineConfigPool) (result *v1.MachineConfigPool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.MachineConfigPool{}
	err = c.client.Put().Resource("machineconfigpools").Name(machineConfigPool.Name).Body(machineConfigPool).Do().Into(result)
	return
}
func (c *machineConfigPools) UpdateStatus(machineConfigPool *v1.MachineConfigPool) (result *v1.MachineConfigPool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.MachineConfigPool{}
	err = c.client.Put().Resource("machineconfigpools").Name(machineConfigPool.Name).SubResource("status").Body(machineConfigPool).Do().Into(result)
	return
}
func (c *machineConfigPools) Delete(name string, options *metav1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("machineconfigpools").Name(name).Body(options).Do().Error()
}
func (c *machineConfigPools) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("machineconfigpools").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *machineConfigPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MachineConfigPool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1.MachineConfigPool{}
	err = c.client.Patch(pt).Resource("machineconfigpools").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}

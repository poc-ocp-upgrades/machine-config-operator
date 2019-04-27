package fake

import (
	machineconfigurationopenshiftiov1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeContainerRuntimeConfigs struct{ Fake *FakeMachineconfigurationV1 }

var containerruntimeconfigsResource = schema.GroupVersionResource{Group: "machineconfiguration.openshift.io", Version: "v1", Resource: "containerruntimeconfigs"}
var containerruntimeconfigsKind = schema.GroupVersionKind{Group: "machineconfiguration.openshift.io", Version: "v1", Kind: "ContainerRuntimeConfig"}

func (c *FakeContainerRuntimeConfigs) Get(name string, options v1.GetOptions) (result *machineconfigurationopenshiftiov1.ContainerRuntimeConfig, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(containerruntimeconfigsResource, name), &machineconfigurationopenshiftiov1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.ContainerRuntimeConfig), err
}
func (c *FakeContainerRuntimeConfigs) List(opts v1.ListOptions) (result *machineconfigurationopenshiftiov1.ContainerRuntimeConfigList, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootListAction(containerruntimeconfigsResource, containerruntimeconfigsKind, opts), &machineconfigurationopenshiftiov1.ContainerRuntimeConfigList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &machineconfigurationopenshiftiov1.ContainerRuntimeConfigList{ListMeta: obj.(*machineconfigurationopenshiftiov1.ContainerRuntimeConfigList).ListMeta}
	for _, item := range obj.(*machineconfigurationopenshiftiov1.ContainerRuntimeConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeContainerRuntimeConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
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
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(containerruntimeconfigsResource, opts))
}
func (c *FakeContainerRuntimeConfigs) Create(containerRuntimeConfig *machineconfigurationopenshiftiov1.ContainerRuntimeConfig) (result *machineconfigurationopenshiftiov1.ContainerRuntimeConfig, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(containerruntimeconfigsResource, containerRuntimeConfig), &machineconfigurationopenshiftiov1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.ContainerRuntimeConfig), err
}
func (c *FakeContainerRuntimeConfigs) Update(containerRuntimeConfig *machineconfigurationopenshiftiov1.ContainerRuntimeConfig) (result *machineconfigurationopenshiftiov1.ContainerRuntimeConfig, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(containerruntimeconfigsResource, containerRuntimeConfig), &machineconfigurationopenshiftiov1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.ContainerRuntimeConfig), err
}
func (c *FakeContainerRuntimeConfigs) UpdateStatus(containerRuntimeConfig *machineconfigurationopenshiftiov1.ContainerRuntimeConfig) (*machineconfigurationopenshiftiov1.ContainerRuntimeConfig, error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(containerruntimeconfigsResource, "status", containerRuntimeConfig), &machineconfigurationopenshiftiov1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.ContainerRuntimeConfig), err
}
func (c *FakeContainerRuntimeConfigs) Delete(name string, options *v1.DeleteOptions) error {
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
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(containerruntimeconfigsResource, name), &machineconfigurationopenshiftiov1.ContainerRuntimeConfig{})
	return err
}
func (c *FakeContainerRuntimeConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
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
	action := testing.NewRootDeleteCollectionAction(containerruntimeconfigsResource, listOptions)
	_, err := c.Fake.Invokes(action, &machineconfigurationopenshiftiov1.ContainerRuntimeConfigList{})
	return err
}
func (c *FakeContainerRuntimeConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *machineconfigurationopenshiftiov1.ContainerRuntimeConfig, err error) {
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
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(containerruntimeconfigsResource, name, pt, data, subresources...), &machineconfigurationopenshiftiov1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.ContainerRuntimeConfig), err
}

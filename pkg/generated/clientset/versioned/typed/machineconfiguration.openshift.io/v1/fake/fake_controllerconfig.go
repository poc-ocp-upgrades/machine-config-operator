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

type FakeControllerConfigs struct{ Fake *FakeMachineconfigurationV1 }

var controllerconfigsResource = schema.GroupVersionResource{Group: "machineconfiguration.openshift.io", Version: "v1", Resource: "controllerconfigs"}
var controllerconfigsKind = schema.GroupVersionKind{Group: "machineconfiguration.openshift.io", Version: "v1", Kind: "ControllerConfig"}

func (c *FakeControllerConfigs) Get(name string, options v1.GetOptions) (result *machineconfigurationopenshiftiov1.ControllerConfig, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(controllerconfigsResource, name), &machineconfigurationopenshiftiov1.ControllerConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.ControllerConfig), err
}
func (c *FakeControllerConfigs) List(opts v1.ListOptions) (result *machineconfigurationopenshiftiov1.ControllerConfigList, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootListAction(controllerconfigsResource, controllerconfigsKind, opts), &machineconfigurationopenshiftiov1.ControllerConfigList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &machineconfigurationopenshiftiov1.ControllerConfigList{ListMeta: obj.(*machineconfigurationopenshiftiov1.ControllerConfigList).ListMeta}
	for _, item := range obj.(*machineconfigurationopenshiftiov1.ControllerConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeControllerConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(controllerconfigsResource, opts))
}
func (c *FakeControllerConfigs) Create(controllerConfig *machineconfigurationopenshiftiov1.ControllerConfig) (result *machineconfigurationopenshiftiov1.ControllerConfig, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(controllerconfigsResource, controllerConfig), &machineconfigurationopenshiftiov1.ControllerConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.ControllerConfig), err
}
func (c *FakeControllerConfigs) Update(controllerConfig *machineconfigurationopenshiftiov1.ControllerConfig) (result *machineconfigurationopenshiftiov1.ControllerConfig, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(controllerconfigsResource, controllerConfig), &machineconfigurationopenshiftiov1.ControllerConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.ControllerConfig), err
}
func (c *FakeControllerConfigs) UpdateStatus(controllerConfig *machineconfigurationopenshiftiov1.ControllerConfig) (*machineconfigurationopenshiftiov1.ControllerConfig, error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(controllerconfigsResource, "status", controllerConfig), &machineconfigurationopenshiftiov1.ControllerConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.ControllerConfig), err
}
func (c *FakeControllerConfigs) Delete(name string, options *v1.DeleteOptions) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(controllerconfigsResource, name), &machineconfigurationopenshiftiov1.ControllerConfig{})
	return err
}
func (c *FakeControllerConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewRootDeleteCollectionAction(controllerconfigsResource, listOptions)
	_, err := c.Fake.Invokes(action, &machineconfigurationopenshiftiov1.ControllerConfigList{})
	return err
}
func (c *FakeControllerConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *machineconfigurationopenshiftiov1.ControllerConfig, err error) {
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(controllerconfigsResource, name, pt, data, subresources...), &machineconfigurationopenshiftiov1.ControllerConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.ControllerConfig), err
}

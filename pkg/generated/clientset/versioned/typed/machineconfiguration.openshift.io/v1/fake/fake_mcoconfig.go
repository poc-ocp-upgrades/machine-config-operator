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

type FakeMCOConfigs struct {
	Fake	*FakeMachineconfigurationV1
	ns	string
}

var mcoconfigsResource = schema.GroupVersionResource{Group: "machineconfiguration.openshift.io", Version: "v1", Resource: "mcoconfigs"}
var mcoconfigsKind = schema.GroupVersionKind{Group: "machineconfiguration.openshift.io", Version: "v1", Kind: "MCOConfig"}

func (c *FakeMCOConfigs) Get(name string, options v1.GetOptions) (result *machineconfigurationopenshiftiov1.MCOConfig, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(mcoconfigsResource, c.ns, name), &machineconfigurationopenshiftiov1.MCOConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MCOConfig), err
}
func (c *FakeMCOConfigs) List(opts v1.ListOptions) (result *machineconfigurationopenshiftiov1.MCOConfigList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(mcoconfigsResource, mcoconfigsKind, c.ns, opts), &machineconfigurationopenshiftiov1.MCOConfigList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &machineconfigurationopenshiftiov1.MCOConfigList{ListMeta: obj.(*machineconfigurationopenshiftiov1.MCOConfigList).ListMeta}
	for _, item := range obj.(*machineconfigurationopenshiftiov1.MCOConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeMCOConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(mcoconfigsResource, c.ns, opts))
}
func (c *FakeMCOConfigs) Create(mCOConfig *machineconfigurationopenshiftiov1.MCOConfig) (result *machineconfigurationopenshiftiov1.MCOConfig, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(mcoconfigsResource, c.ns, mCOConfig), &machineconfigurationopenshiftiov1.MCOConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MCOConfig), err
}
func (c *FakeMCOConfigs) Update(mCOConfig *machineconfigurationopenshiftiov1.MCOConfig) (result *machineconfigurationopenshiftiov1.MCOConfig, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(mcoconfigsResource, c.ns, mCOConfig), &machineconfigurationopenshiftiov1.MCOConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MCOConfig), err
}
func (c *FakeMCOConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(mcoconfigsResource, c.ns, name), &machineconfigurationopenshiftiov1.MCOConfig{})
	return err
}
func (c *FakeMCOConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(mcoconfigsResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &machineconfigurationopenshiftiov1.MCOConfigList{})
	return err
}
func (c *FakeMCOConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *machineconfigurationopenshiftiov1.MCOConfig, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(mcoconfigsResource, c.ns, name, pt, data, subresources...), &machineconfigurationopenshiftiov1.MCOConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MCOConfig), err
}

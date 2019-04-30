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

type FakeMachineConfigPools struct{ Fake *FakeMachineconfigurationV1 }

var machineconfigpoolsResource = schema.GroupVersionResource{Group: "machineconfiguration.openshift.io", Version: "v1", Resource: "machineconfigpools"}
var machineconfigpoolsKind = schema.GroupVersionKind{Group: "machineconfiguration.openshift.io", Version: "v1", Kind: "MachineConfigPool"}

func (c *FakeMachineConfigPools) Get(name string, options v1.GetOptions) (result *machineconfigurationopenshiftiov1.MachineConfigPool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(machineconfigpoolsResource, name), &machineconfigurationopenshiftiov1.MachineConfigPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MachineConfigPool), err
}
func (c *FakeMachineConfigPools) List(opts v1.ListOptions) (result *machineconfigurationopenshiftiov1.MachineConfigPoolList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootListAction(machineconfigpoolsResource, machineconfigpoolsKind, opts), &machineconfigurationopenshiftiov1.MachineConfigPoolList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &machineconfigurationopenshiftiov1.MachineConfigPoolList{ListMeta: obj.(*machineconfigurationopenshiftiov1.MachineConfigPoolList).ListMeta}
	for _, item := range obj.(*machineconfigurationopenshiftiov1.MachineConfigPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeMachineConfigPools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(machineconfigpoolsResource, opts))
}
func (c *FakeMachineConfigPools) Create(machineConfigPool *machineconfigurationopenshiftiov1.MachineConfigPool) (result *machineconfigurationopenshiftiov1.MachineConfigPool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(machineconfigpoolsResource, machineConfigPool), &machineconfigurationopenshiftiov1.MachineConfigPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MachineConfigPool), err
}
func (c *FakeMachineConfigPools) Update(machineConfigPool *machineconfigurationopenshiftiov1.MachineConfigPool) (result *machineconfigurationopenshiftiov1.MachineConfigPool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(machineconfigpoolsResource, machineConfigPool), &machineconfigurationopenshiftiov1.MachineConfigPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MachineConfigPool), err
}
func (c *FakeMachineConfigPools) UpdateStatus(machineConfigPool *machineconfigurationopenshiftiov1.MachineConfigPool) (*machineconfigurationopenshiftiov1.MachineConfigPool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(machineconfigpoolsResource, "status", machineConfigPool), &machineconfigurationopenshiftiov1.MachineConfigPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MachineConfigPool), err
}
func (c *FakeMachineConfigPools) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(machineconfigpoolsResource, name), &machineconfigurationopenshiftiov1.MachineConfigPool{})
	return err
}
func (c *FakeMachineConfigPools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewRootDeleteCollectionAction(machineconfigpoolsResource, listOptions)
	_, err := c.Fake.Invokes(action, &machineconfigurationopenshiftiov1.MachineConfigPoolList{})
	return err
}
func (c *FakeMachineConfigPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *machineconfigurationopenshiftiov1.MachineConfigPool, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(machineconfigpoolsResource, name, pt, data, subresources...), &machineconfigurationopenshiftiov1.MachineConfigPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.MachineConfigPool), err
}

package v1

import (
	v1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type MachineConfigPoolLister interface {
	List(selector labels.Selector) (ret []*v1.MachineConfigPool, err error)
	Get(name string) (*v1.MachineConfigPool, error)
	MachineConfigPoolListerExpansion
}
type machineConfigPoolLister struct{ indexer cache.Indexer }

func NewMachineConfigPoolLister(indexer cache.Indexer) MachineConfigPoolLister {
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
	return &machineConfigPoolLister{indexer: indexer}
}
func (s *machineConfigPoolLister) List(selector labels.Selector) (ret []*v1.MachineConfigPool, err error) {
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
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MachineConfigPool))
	})
	return ret, err
}
func (s *machineConfigPoolLister) Get(name string) (*v1.MachineConfigPool, error) {
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
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("machineconfigpool"), name)
	}
	return obj.(*v1.MachineConfigPool), nil
}

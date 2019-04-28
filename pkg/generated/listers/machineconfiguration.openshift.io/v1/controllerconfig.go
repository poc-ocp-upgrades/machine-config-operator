package v1

import (
	v1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ControllerConfigLister interface {
	List(selector labels.Selector) (ret []*v1.ControllerConfig, err error)
	Get(name string) (*v1.ControllerConfig, error)
	ControllerConfigListerExpansion
}
type controllerConfigLister struct{ indexer cache.Indexer }

func NewControllerConfigLister(indexer cache.Indexer) ControllerConfigLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &controllerConfigLister{indexer: indexer}
}
func (s *controllerConfigLister) List(selector labels.Selector) (ret []*v1.ControllerConfig, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ControllerConfig))
	})
	return ret, err
}
func (s *controllerConfigLister) Get(name string) (*v1.ControllerConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("controllerconfig"), name)
	}
	return obj.(*v1.ControllerConfig), nil
}

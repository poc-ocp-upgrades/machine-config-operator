package v1

import (
	v1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type MCOConfigLister interface {
	List(selector labels.Selector) (ret []*v1.MCOConfig, err error)
	MCOConfigs(namespace string) MCOConfigNamespaceLister
	MCOConfigListerExpansion
}
type mCOConfigLister struct{ indexer cache.Indexer }

func NewMCOConfigLister(indexer cache.Indexer) MCOConfigLister {
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
	return &mCOConfigLister{indexer: indexer}
}
func (s *mCOConfigLister) List(selector labels.Selector) (ret []*v1.MCOConfig, err error) {
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
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MCOConfig))
	})
	return ret, err
}
func (s *mCOConfigLister) MCOConfigs(namespace string) MCOConfigNamespaceLister {
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
	return mCOConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type MCOConfigNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1.MCOConfig, err error)
	Get(name string) (*v1.MCOConfig, error)
	MCOConfigNamespaceListerExpansion
}
type mCOConfigNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s mCOConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.MCOConfig, err error) {
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
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MCOConfig))
	})
	return ret, err
}
func (s mCOConfigNamespaceLister) Get(name string) (*v1.MCOConfig, error) {
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
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("mcoconfig"), name)
	}
	return obj.(*v1.MCOConfig), nil
}

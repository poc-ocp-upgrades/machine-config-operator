package externalversions

import (
	"fmt"
	v1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}
type genericInformer struct {
	informer	cache.SharedIndexInformer
	resource	schema.GroupResource
}

func (f *genericInformer) Informer() cache.SharedIndexInformer {
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
	return f.informer
}
func (f *genericInformer) Lister() cache.GenericLister {
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
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
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
	switch resource {
	case v1.SchemeGroupVersion.WithResource("containerruntimeconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machineconfiguration().V1().ContainerRuntimeConfigs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("controllerconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machineconfiguration().V1().ControllerConfigs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("kubeletconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machineconfiguration().V1().KubeletConfigs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("mcoconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machineconfiguration().V1().MCOConfigs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("machineconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machineconfiguration().V1().MachineConfigs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("machineconfigpools"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Machineconfiguration().V1().MachineConfigPools().Informer()}, nil
	}
	return nil, fmt.Errorf("no informer found for %v", resource)
}

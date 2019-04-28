package v1

import (
	time "time"
	machineconfigurationopenshiftiov1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	versioned "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/machine-config-operator/pkg/generated/listers/machineconfiguration.openshift.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

type KubeletConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.KubeletConfigLister
}
type kubeletConfigInformer struct {
	factory			internalinterfaces.SharedInformerFactory
	tweakListOptions	internalinterfaces.TweakListOptionsFunc
}

func NewKubeletConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredKubeletConfigInformer(client, resyncPeriod, indexers, nil)
}
func NewFilteredKubeletConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.NewSharedIndexInformer(&cache.ListWatch{ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.MachineconfigurationV1().KubeletConfigs().List(options)
	}, WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.MachineconfigurationV1().KubeletConfigs().Watch(options)
	}}, &machineconfigurationopenshiftiov1.KubeletConfig{}, resyncPeriod, indexers)
}
func (f *kubeletConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredKubeletConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}
func (f *kubeletConfigInformer) Informer() cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.factory.InformerFor(&machineconfigurationopenshiftiov1.KubeletConfig{}, f.defaultInformer)
}
func (f *kubeletConfigInformer) Lister() v1.KubeletConfigLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return v1.NewKubeletConfigLister(f.Informer().GetIndexer())
}

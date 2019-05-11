package common

import (
	"math/rand"
	"time"
	configinformers "github.com/openshift/client-go/config/informers/externalversions"
	"github.com/openshift/machine-config-operator/internal/clients"
	mcfginformers "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions"
	apiextinformers "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
)

const (
	minResyncPeriod = 20 * time.Minute
)

func resyncPeriod() func() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func() time.Duration {
		factor := rand.Float64() + 1
		return time.Duration(float64(minResyncPeriod.Nanoseconds()) * factor)
	}
}

type ControllerContext struct {
	ClientBuilder									*clients.Builder
	NamespacedInformerFactory						mcfginformers.SharedInformerFactory
	InformerFactory									mcfginformers.SharedInformerFactory
	KubeInformerFactory								informers.SharedInformerFactory
	KubeNamespacedInformerFactory					informers.SharedInformerFactory
	OpenShiftConfigKubeNamespacedInformerFactory	informers.SharedInformerFactory
	APIExtInformerFactory							apiextinformers.SharedInformerFactory
	ConfigInformerFactory							configinformers.SharedInformerFactory
	AvailableResources								map[schema.GroupVersionResource]bool
	Stop											<-chan struct{}
	InformersStarted								chan struct{}
	ResyncPeriod									func() time.Duration
}

func CreateControllerContext(cb *clients.Builder, stop <-chan struct{}, targetNamespace string) *ControllerContext {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client := cb.MachineConfigClientOrDie("machine-config-shared-informer")
	kubeClient := cb.KubeClientOrDie("kube-shared-informer")
	apiExtClient := cb.APIExtClientOrDie("apiext-shared-informer")
	configClient := cb.ConfigClientOrDie("config-shared-informer")
	sharedInformers := mcfginformers.NewSharedInformerFactory(client, resyncPeriod()())
	sharedNamespacedInformers := mcfginformers.NewFilteredSharedInformerFactory(client, resyncPeriod()(), targetNamespace, nil)
	kubeSharedInformer := informers.NewSharedInformerFactory(kubeClient, resyncPeriod()())
	kubeNamespacedSharedInformer := informers.NewFilteredSharedInformerFactory(kubeClient, resyncPeriod()(), targetNamespace, nil)
	openShiftConfigKubeNamespacedSharedInformer := informers.NewFilteredSharedInformerFactory(kubeClient, resyncPeriod()(), "openshift-config", nil)
	apiExtSharedInformer := apiextinformers.NewSharedInformerFactory(apiExtClient, resyncPeriod()())
	configSharedInformer := configinformers.NewSharedInformerFactory(configClient, resyncPeriod()())
	return &ControllerContext{ClientBuilder: cb, NamespacedInformerFactory: sharedNamespacedInformers, InformerFactory: sharedInformers, KubeInformerFactory: kubeSharedInformer, KubeNamespacedInformerFactory: kubeNamespacedSharedInformer, OpenShiftConfigKubeNamespacedInformerFactory: openShiftConfigKubeNamespacedSharedInformer, APIExtInformerFactory: apiExtSharedInformer, ConfigInformerFactory: configSharedInformer, Stop: stop, InformersStarted: make(chan struct{}), ResyncPeriod: resyncPeriod()}
}

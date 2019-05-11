package operator

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"github.com/golang/glog"
	configclientset "github.com/openshift/client-go/config/clientset/versioned"
	v1 "k8s.io/api/core/v1"
	apiextclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apiextinformersv1beta1 "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions/apiextensions/v1beta1"
	apiextlistersv1beta1 "k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	appsinformersv1 "k8s.io/client-go/informers/apps/v1"
	coreinformersv1 "k8s.io/client-go/informers/core/v1"
	rbacinformersv1 "k8s.io/client-go/informers/rbac/v1"
	"k8s.io/client-go/kubernetes"
	coreclientsetv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	appslisterv1 "k8s.io/client-go/listers/apps/v1"
	corelisterv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	configv1 "github.com/openshift/api/config/v1"
	configinformersv1 "github.com/openshift/client-go/config/informers/externalversions/config/v1"
	configlistersv1 "github.com/openshift/client-go/config/listers/config/v1"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	templatectrl "github.com/openshift/machine-config-operator/pkg/controller/template"
	mcfgclientset "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned"
	"github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/scheme"
	mcfginformersv1 "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/machineconfiguration.openshift.io/v1"
	mcfglistersv1 "github.com/openshift/machine-config-operator/pkg/generated/listers/machineconfiguration.openshift.io/v1"
	"github.com/openshift/machine-config-operator/pkg/version"
)

const (
	maxRetries				= 15
	osImageConfigMapName	= "machine-config-osimageurl"
)

type Operator struct {
	namespace, name			string
	inClusterBringup		bool
	imagesFile				string
	vStore					*versionStore
	client					mcfgclientset.Interface
	kubeClient				kubernetes.Interface
	apiExtClient			apiextclientset.Interface
	configClient			configclientset.Interface
	eventRecorder			record.EventRecorder
	syncHandler				func(ic string) error
	crdLister				apiextlistersv1beta1.CustomResourceDefinitionLister
	mcpLister				mcfglistersv1.MachineConfigPoolLister
	ccLister				mcfglistersv1.ControllerConfigLister
	mcLister				mcfglistersv1.MachineConfigLister
	deployLister			appslisterv1.DeploymentLister
	daemonsetLister			appslisterv1.DaemonSetLister
	infraLister				configlistersv1.InfrastructureLister
	networkLister			configlistersv1.NetworkLister
	mcoCmLister				corelisterv1.ConfigMapLister
	clusterCmLister			corelisterv1.ConfigMapLister
	crdListerSynced			cache.InformerSynced
	deployListerSynced		cache.InformerSynced
	daemonsetListerSynced	cache.InformerSynced
	infraListerSynced		cache.InformerSynced
	networkListerSynced		cache.InformerSynced
	mcpListerSynced			cache.InformerSynced
	ccListerSynced			cache.InformerSynced
	mcListerSynced			cache.InformerSynced
	mcoCmListerSynced		cache.InformerSynced
	clusterCmListerSynced	cache.InformerSynced
	queue					workqueue.RateLimitingInterface
	stopCh					<-chan struct{}
}

func New(namespace, name string, imagesFile string, mcpInformer mcfginformersv1.MachineConfigPoolInformer, ccInformer mcfginformersv1.ControllerConfigInformer, mcInformer mcfginformersv1.MachineConfigInformer, controllerConfigInformer mcfginformersv1.ControllerConfigInformer, serviceAccountInfomer coreinformersv1.ServiceAccountInformer, crdInformer apiextinformersv1beta1.CustomResourceDefinitionInformer, deployInformer appsinformersv1.DeploymentInformer, daemonsetInformer appsinformersv1.DaemonSetInformer, clusterRoleInformer rbacinformersv1.ClusterRoleInformer, clusterRoleBindingInformer rbacinformersv1.ClusterRoleBindingInformer, mcoCmInformer coreinformersv1.ConfigMapInformer, clusterCmInfomer coreinformersv1.ConfigMapInformer, infraInformer configinformersv1.InfrastructureInformer, networkInformer configinformersv1.NetworkInformer, client mcfgclientset.Interface, kubeClient kubernetes.Interface, apiExtClient apiextclientset.Interface, configClient configclientset.Interface) *Operator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&coreclientsetv1.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	optr := &Operator{namespace: namespace, name: name, imagesFile: imagesFile, vStore: newVersionStore(), client: client, kubeClient: kubeClient, apiExtClient: apiExtClient, configClient: configClient, eventRecorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "machineconfigoperator"}), queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "machineconfigoperator")}
	for _, i := range []cache.SharedIndexInformer{controllerConfigInformer.Informer(), serviceAccountInfomer.Informer(), crdInformer.Informer(), deployInformer.Informer(), daemonsetInformer.Informer(), clusterRoleInformer.Informer(), clusterRoleBindingInformer.Informer(), mcoCmInformer.Informer(), infraInformer.Informer(), networkInformer.Informer()} {
		i.AddEventHandler(optr.eventHandler())
	}
	optr.syncHandler = optr.sync
	optr.clusterCmLister = clusterCmInfomer.Lister()
	optr.clusterCmListerSynced = clusterCmInfomer.Informer().HasSynced
	optr.mcoCmLister = mcoCmInformer.Lister()
	optr.mcoCmListerSynced = mcoCmInformer.Informer().HasSynced
	optr.crdLister = crdInformer.Lister()
	optr.crdListerSynced = crdInformer.Informer().HasSynced
	optr.mcpLister = mcpInformer.Lister()
	optr.mcpListerSynced = mcpInformer.Informer().HasSynced
	optr.ccLister = ccInformer.Lister()
	optr.ccListerSynced = ccInformer.Informer().HasSynced
	optr.mcLister = mcInformer.Lister()
	optr.mcListerSynced = mcInformer.Informer().HasSynced
	optr.deployLister = deployInformer.Lister()
	optr.deployListerSynced = deployInformer.Informer().HasSynced
	optr.daemonsetLister = daemonsetInformer.Lister()
	optr.daemonsetListerSynced = daemonsetInformer.Informer().HasSynced
	optr.infraLister = infraInformer.Lister()
	optr.infraListerSynced = infraInformer.Informer().HasSynced
	optr.networkLister = networkInformer.Lister()
	optr.networkListerSynced = networkInformer.Informer().HasSynced
	optr.vStore.Set("operator", os.Getenv("RELEASE_VERSION"))
	return optr
}
func (optr *Operator) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	defer optr.queue.ShutDown()
	glog.Info("Starting MachineConfigOperator")
	defer glog.Info("Shutting down MachineConfigOperator")
	apiClient := optr.apiExtClient.ApiextensionsV1beta1()
	_, err := apiClient.CustomResourceDefinitions().Get("machineconfigpools.machineconfiguration.openshift.io", metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			glog.Infof("Couldn't find machineconfigpool CRD, in cluster bringup mode")
			optr.inClusterBringup = true
		} else {
			glog.Errorf("While checking for cluster bringup: %v", err)
		}
	}
	if !cache.WaitForCacheSync(stopCh, optr.crdListerSynced, optr.deployListerSynced, optr.daemonsetListerSynced, optr.infraListerSynced, optr.mcoCmListerSynced, optr.clusterCmListerSynced, optr.networkListerSynced) {
		glog.Error("failed to sync caches")
		return
	}
	if !optr.inClusterBringup {
		if !cache.WaitForCacheSync(stopCh, optr.mcpListerSynced, optr.ccListerSynced, optr.mcListerSynced) {
			glog.Error("failed to sync caches")
			return
		}
	}
	optr.stopCh = stopCh
	for i := 0; i < workers; i++ {
		go wait.Until(optr.worker, time.Second, stopCh)
	}
	<-stopCh
}
func (optr *Operator) eventHandler() cache.ResourceEventHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	workQueueKey := fmt.Sprintf("%s/%s", optr.namespace, optr.name)
	return cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		optr.queue.Add(workQueueKey)
	}, UpdateFunc: func(old, new interface{}) {
		optr.queue.Add(workQueueKey)
	}, DeleteFunc: func(obj interface{}) {
		optr.queue.Add(workQueueKey)
	}}
}
func (optr *Operator) worker() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for optr.processNextWorkItem() {
	}
}
func (optr *Operator) processNextWorkItem() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, quit := optr.queue.Get()
	if quit {
		return false
	}
	defer optr.queue.Done(key)
	err := optr.syncHandler(key.(string))
	optr.handleErr(err, key)
	return true
}
func (optr *Operator) handleErr(err error, key interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		optr.queue.Forget(key)
		return
	}
	if optr.queue.NumRequeues(key) < maxRetries {
		glog.V(2).Infof("Error syncing operator %v: %v", key, err)
		optr.queue.AddRateLimited(key)
		return
	}
	utilruntime.HandleError(err)
	glog.V(2).Infof("Dropping operator %q out of the queue: %v", key, err)
	optr.queue.Forget(key)
	optr.queue.AddAfter(key, 1*time.Minute)
}
func (optr *Operator) sync(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	startTime := time.Now()
	glog.V(4).Infof("Started syncing operator %q (%v)", key, startTime)
	defer func() {
		glog.V(4).Infof("Finished syncing operator %q (%v)", key, time.Since(startTime))
	}()
	if err := optr.syncCustomResourceDefinitions(); err != nil {
		return err
	}
	if optr.inClusterBringup {
		if !cache.WaitForCacheSync(optr.stopCh, optr.mcpListerSynced, optr.mcListerSynced, optr.ccListerSynced) {
			return errors.New("failed to sync caches for informers")
		}
	}
	namespace, _, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	imgsRaw, err := ioutil.ReadFile(optr.imagesFile)
	if err != nil {
		return err
	}
	imgs := Images{}
	if err := json.Unmarshal(imgsRaw, &imgs); err != nil {
		return err
	}
	etcdCA, err := optr.getCAsFromConfigMap("openshift-config", "etcd-serving-ca", "ca-bundle.crt")
	if err != nil {
		return err
	}
	etcdMetricCA, err := optr.getCAsFromConfigMap("openshift-config", "etcd-metric-serving-ca", "ca-bundle.crt")
	if err != nil {
		return err
	}
	rootCA, err := optr.getCAsFromConfigMap("kube-system", "root-ca", "ca.crt")
	if err != nil {
		return err
	}
	kubeAPIServerServingCABytes, err := optr.getCAsFromConfigMap("openshift-config", "initial-kube-apiserver-server-ca", "ca-bundle.crt")
	if err != nil {
		return err
	}
	bundle := make([]byte, 0)
	bundle = append(bundle, rootCA...)
	bundle = append(bundle, kubeAPIServerServingCABytes...)
	osimageurl, err := optr.getOsImageURL(namespace)
	if err != nil {
		return err
	}
	imgs.MachineOSContent = osimageurl
	infra, network, err := optr.getGlobalConfig()
	if err != nil {
		return err
	}
	spec, err := createDiscoveredControllerConfigSpec(infra, network)
	if err != nil {
		return err
	}
	spec.EtcdCAData = etcdCA
	spec.EtcdMetricCAData = etcdMetricCA
	spec.RootCAData = bundle
	spec.PullSecret = &v1.ObjectReference{Namespace: "openshift-config", Name: "pull-secret"}
	spec.OSImageURL = imgs.MachineOSContent
	spec.Images = map[string]string{templatectrl.EtcdImageKey: imgs.Etcd, templatectrl.SetupEtcdEnvKey: imgs.SetupEtcdEnv, templatectrl.InfraImageKey: imgs.InfraImage, templatectrl.KubeClientAgentImageKey: imgs.KubeClientAgent}
	rc := getRenderConfig(namespace, string(kubeAPIServerServingCABytes), spec, imgs, infra.Status.APIServerURL)
	var syncFuncs = []syncFunc{{"pools", optr.syncMachineConfigPools}, {"mcc", optr.syncMachineConfigController}, {"mcs", optr.syncMachineConfigServer}, {"mcd", optr.syncMachineConfigDaemon}, {"required-pools", optr.syncRequiredMachineConfigPools}}
	return optr.syncAll(rc, syncFuncs)
}
func (optr *Operator) getOsImageURL(namespace string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cm, err := optr.mcoCmLister.ConfigMaps(namespace).Get(osImageConfigMapName)
	if err != nil {
		return "", err
	}
	return cm.Data["osImageURL"], nil
}
func (optr *Operator) getCAsFromConfigMap(namespace, name, key string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cm, err := optr.clusterCmLister.ConfigMaps(namespace).Get(name)
	if err != nil {
		return nil, err
	}
	if bd, bdok := cm.BinaryData[key]; bdok {
		return bd, nil
	} else if d, dok := cm.Data[key]; dok {
		raw, err := base64.StdEncoding.DecodeString(d)
		if err != nil {
			return []byte(d), nil
		}
		return raw, nil
	} else {
		return nil, fmt.Errorf("%s not found in %s/%s", key, namespace, name)
	}
}
func (optr *Operator) getGlobalConfig() (*configv1.Infrastructure, *configv1.Network, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	infra, err := optr.infraLister.Get("cluster")
	if err != nil {
		return nil, nil, err
	}
	network, err := optr.networkLister.Get("cluster")
	if err != nil {
		return nil, nil, err
	}
	return infra, network, nil
}
func getRenderConfig(tnamespace, kubeAPIServerServingCA string, ccSpec *mcfgv1.ControllerConfigSpec, imgs Images, apiServerURL string) renderConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return renderConfig{TargetNamespace: tnamespace, Version: version.Raw, ControllerConfig: *ccSpec, Images: imgs, APIServerURL: apiServerURL, KubeAPIServerServingCA: kubeAPIServerServingCA}
}

package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"time"
	"github.com/golang/glog"
	"github.com/openshift/machine-config-operator/lib/resourceapply"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"github.com/openshift/machine-config-operator/pkg/controller/common"
	mcfgclientset "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned"
	"github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/scheme"
	mcfginformersv1 "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/machineconfiguration.openshift.io/v1"
	mcfglistersv1 "github.com/openshift/machine-config-operator/pkg/generated/listers/machineconfiguration.openshift.io/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformersv1 "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	corev1clientset "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
)

const (
	maxRetries = 15
)

var controllerKind = mcfgv1.SchemeGroupVersion.WithKind("ControllerConfig")

type Controller struct {
	templatesDir		string
	client			mcfgclientset.Interface
	kubeClient		clientset.Interface
	eventRecorder		record.EventRecorder
	syncHandler		func(ccKey string) error
	enqueueControllerConfig	func(*mcfgv1.ControllerConfig)
	ccLister		mcfglistersv1.ControllerConfigLister
	mcLister		mcfglistersv1.MachineConfigLister
	ccListerSynced		cache.InformerSynced
	mcListerSynced		cache.InformerSynced
	secretsInformerSynced	cache.InformerSynced
	queue			workqueue.RateLimitingInterface
}

func New(templatesDir string, ccInformer mcfginformersv1.ControllerConfigInformer, mcInformer mcfginformersv1.MachineConfigInformer, secretsInformer coreinformersv1.SecretInformer, kubeClient clientset.Interface, mcfgClient mcfgclientset.Interface) *Controller {
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
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&corev1clientset.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	ctrl := &Controller{templatesDir: templatesDir, client: mcfgClient, kubeClient: kubeClient, eventRecorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "machineconfigcontroller-templatecontroller"}), queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "machineconfigcontroller-templatecontroller")}
	ccInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: ctrl.addControllerConfig, UpdateFunc: ctrl.updateControllerConfig, DeleteFunc: ctrl.deleteControllerConfig})
	mcInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: ctrl.addMachineConfig, UpdateFunc: ctrl.updateMachineConfig, DeleteFunc: ctrl.deleteMachineConfig})
	secretsInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: ctrl.addSecret, UpdateFunc: ctrl.updateSecret, DeleteFunc: ctrl.deleteSecret})
	ctrl.syncHandler = ctrl.syncControllerConfig
	ctrl.enqueueControllerConfig = ctrl.enqueue
	ctrl.ccLister = ccInformer.Lister()
	ctrl.mcLister = mcInformer.Lister()
	ctrl.ccListerSynced = ccInformer.Informer().HasSynced
	ctrl.mcListerSynced = mcInformer.Informer().HasSynced
	ctrl.secretsInformerSynced = secretsInformer.Informer().HasSynced
	return ctrl
}
func (ctrl *Controller) filterSecret(secret *v1.Secret) {
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
	if secret.Name == "pull-secret" {
		cfg, err := ctrl.ccLister.Get(common.ControllerConfigName)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("couldn't get ControllerConfig on secret callback %#v", err))
			return
		}
		glog.V(4).Infof("Re-syncing ControllerConfig %s due to secret change", cfg.Name)
		ctrl.enqueueControllerConfig(cfg)
	}
}
func (ctrl *Controller) addSecret(obj interface{}) {
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
	secret := obj.(*v1.Secret)
	if secret.DeletionTimestamp != nil {
		ctrl.deleteSecret(secret)
		return
	}
	glog.V(4).Infof("Add Secret %v", secret)
	ctrl.filterSecret(secret)
}
func (ctrl *Controller) updateSecret(old, new interface{}) {
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
	secret := new.(*v1.Secret)
	glog.V(4).Infof("Update Secret %v", secret)
	ctrl.filterSecret(secret)
}
func (ctrl *Controller) deleteSecret(obj interface{}) {
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
	secret, ok := obj.(*v1.Secret)
	glog.V(4).Infof("Delete Secret %v", secret)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
			return
		}
		secret, ok = tombstone.Obj.(*v1.Secret)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a Secret %#v", obj))
			return
		}
	}
	if secret.Name == "pull-secret" {
		cfg, err := ctrl.ccLister.Get(common.ControllerConfigName)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("Couldn't get ControllerConfig on secret callback %#v", err))
			return
		}
		glog.V(4).Infof("Re-syncing ControllerConfig %s due to secret deletion", cfg.Name)
	}
}
func (ctrl *Controller) Run(workers int, stopCh <-chan struct{}) {
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
	defer utilruntime.HandleCrash()
	defer ctrl.queue.ShutDown()
	if !cache.WaitForCacheSync(stopCh, ctrl.ccListerSynced, ctrl.mcListerSynced, ctrl.secretsInformerSynced) {
		return
	}
	glog.Info("Starting MachineConfigController-TemplateController")
	defer glog.Info("Shutting down MachineConfigController-TemplateController")
	for i := 0; i < workers; i++ {
		go wait.Until(ctrl.worker, time.Second, stopCh)
	}
	<-stopCh
}
func (ctrl *Controller) addControllerConfig(obj interface{}) {
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
	cfg := obj.(*mcfgv1.ControllerConfig)
	glog.V(4).Infof("Adding ControllerConfig %s", cfg.Name)
	ctrl.enqueueControllerConfig(cfg)
}
func (ctrl *Controller) updateControllerConfig(old, cur interface{}) {
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
	oldCfg := old.(*mcfgv1.ControllerConfig)
	curCfg := cur.(*mcfgv1.ControllerConfig)
	glog.V(4).Infof("Updating ControllerConfig %s", oldCfg.Name)
	ctrl.enqueueControllerConfig(curCfg)
}
func (ctrl *Controller) deleteControllerConfig(obj interface{}) {
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
	cfg, ok := obj.(*mcfgv1.ControllerConfig)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
			return
		}
		cfg, ok = tombstone.Obj.(*mcfgv1.ControllerConfig)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a ControllerConfig %#v", obj))
			return
		}
	}
	glog.V(4).Infof("Deleting ControllerConfig %s", cfg.Name)
}
func (ctrl *Controller) addMachineConfig(obj interface{}) {
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
	mc := obj.(*mcfgv1.MachineConfig)
	if mc.DeletionTimestamp != nil {
		ctrl.deleteMachineConfig(mc)
		return
	}
	if controllerRef := metav1.GetControllerOf(mc); controllerRef != nil {
		cfg := ctrl.resolveControllerRef(controllerRef)
		if cfg == nil {
			return
		}
		glog.V(4).Infof("MachineConfig %s added", mc.Name)
		ctrl.enqueueControllerConfig(cfg)
		return
	}
}
func (ctrl *Controller) updateMachineConfig(old, cur interface{}) {
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
	curMC := cur.(*mcfgv1.MachineConfig)
	if controllerRef := metav1.GetControllerOf(curMC); controllerRef != nil {
		cfg := ctrl.resolveControllerRef(controllerRef)
		if cfg == nil {
			return
		}
		glog.V(4).Infof("MachineConfig %s updated", curMC.Name)
		ctrl.enqueueControllerConfig(cfg)
		return
	}
}
func (ctrl *Controller) deleteMachineConfig(obj interface{}) {
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
	mc, ok := obj.(*mcfgv1.MachineConfig)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
			return
		}
		mc, ok = tombstone.Obj.(*mcfgv1.MachineConfig)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a MachineConfig %#v", obj))
			return
		}
	}
	controllerRef := metav1.GetControllerOf(mc)
	if controllerRef == nil {
		return
	}
	cfg := ctrl.resolveControllerRef(controllerRef)
	if cfg == nil {
		return
	}
	glog.V(4).Infof("MachineConfig %s deleted.", mc.Name)
	ctrl.enqueueControllerConfig(cfg)
}
func (ctrl *Controller) resolveControllerRef(controllerRef *metav1.OwnerReference) *mcfgv1.ControllerConfig {
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
	if controllerRef.Kind != controllerKind.Kind {
		return nil
	}
	cfg, err := ctrl.ccLister.Get(controllerRef.Name)
	if err != nil {
		return nil
	}
	if cfg.UID != controllerRef.UID {
		return nil
	}
	return cfg
}
func (ctrl *Controller) enqueue(config *mcfgv1.ControllerConfig) {
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
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(config)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", config, err))
		return
	}
	ctrl.queue.Add(key)
}
func (ctrl *Controller) enqueueRateLimited(controllerconfig *mcfgv1.ControllerConfig) {
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
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(controllerconfig)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", controllerconfig, err))
		return
	}
	ctrl.queue.AddRateLimited(key)
}
func (ctrl *Controller) enqueueAfter(controllerconfig *mcfgv1.ControllerConfig, after time.Duration) {
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
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(controllerconfig)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", controllerconfig, err))
		return
	}
	ctrl.queue.AddAfter(key, after)
}
func (ctrl *Controller) worker() {
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
	for ctrl.processNextWorkItem() {
	}
}
func (ctrl *Controller) processNextWorkItem() bool {
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
	key, quit := ctrl.queue.Get()
	if quit {
		return false
	}
	defer ctrl.queue.Done(key)
	err := ctrl.syncHandler(key.(string))
	ctrl.handleErr(err, key)
	return true
}
func (ctrl *Controller) handleErr(err error, key interface{}) {
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
	if err == nil {
		ctrl.queue.Forget(key)
		return
	}
	if ctrl.queue.NumRequeues(key) < maxRetries {
		glog.V(2).Infof("Error syncing controllerconfig %v: %v", key, err)
		ctrl.queue.AddRateLimited(key)
		return
	}
	utilruntime.HandleError(err)
	glog.V(2).Infof("Dropping controllerconfig %q out of the queue: %v", key, err)
	ctrl.queue.Forget(key)
	ctrl.queue.AddAfter(key, 1*time.Minute)
}
func (ctrl *Controller) syncControllerConfig(key string) error {
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
	startTime := time.Now()
	glog.V(4).Infof("Started syncing controllerconfig %q (%v)", key, startTime)
	defer func() {
		glog.V(4).Infof("Finished syncing controllerconfig %q (%v)", key, time.Since(startTime))
	}()
	_, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	controllerconfig, err := ctrl.ccLister.Get(name)
	if errors.IsNotFound(err) {
		glog.V(2).Infof("ControllerConfig %v has been deleted", key)
		return nil
	}
	if err != nil {
		return err
	}
	cfg := controllerconfig.DeepCopy()
	if cfg.GetGeneration() != cfg.Status.ObservedGeneration {
		if err := ctrl.syncRunningStatus(cfg); err != nil {
			return err
		}
	}
	var pullSecretRaw []byte
	if cfg.Spec.PullSecret != nil {
		secret, err := ctrl.kubeClient.CoreV1().Secrets(cfg.Spec.PullSecret.Namespace).Get(cfg.Spec.PullSecret.Name, metav1.GetOptions{})
		if err != nil {
			return ctrl.syncFailingStatus(cfg, err)
		}
		if secret.Type != corev1.SecretTypeDockerConfigJson {
			return ctrl.syncFailingStatus(cfg, fmt.Errorf("expected secret type %s found %s", corev1.SecretTypeDockerConfigJson, secret.Type))
		}
		pullSecretRaw = secret.Data[corev1.DockerConfigJsonKey]
	}
	mcs, err := getMachineConfigsForControllerConfig(ctrl.templatesDir, cfg, pullSecretRaw)
	if err != nil {
		return ctrl.syncFailingStatus(cfg, err)
	}
	for _, mc := range mcs {
		_, updated, err := resourceapply.ApplyMachineConfig(ctrl.client.MachineconfigurationV1(), mc)
		if err != nil {
			return ctrl.syncFailingStatus(cfg, err)
		}
		if updated {
			glog.V(4).Infof("Machineconfig %s was updated", mc.Name)
		}
	}
	return ctrl.syncCompletedStatus(cfg)
}
func getMachineConfigsForControllerConfig(templatesDir string, config *mcfgv1.ControllerConfig, pullSecretRaw []byte) ([]*mcfgv1.MachineConfig, error) {
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
	buf := &bytes.Buffer{}
	if err := json.Compact(buf, pullSecretRaw); err != nil {
		return nil, fmt.Errorf("couldn't compact pullsecret %q: %v", string(pullSecretRaw), err)
	}
	rc := &RenderConfig{ControllerConfigSpec: &config.Spec, PullSecret: string(buf.Bytes())}
	mcs, err := generateTemplateMachineConfigs(rc, templatesDir)
	if err != nil {
		return nil, err
	}
	for _, mc := range mcs {
		oref := metav1.NewControllerRef(config, controllerKind)
		mc.SetOwnerReferences([]metav1.OwnerReference{*oref})
	}
	sort.Slice(mcs, func(i, j int) bool {
		return mcs[i].Name < mcs[j].Name
	})
	return mcs, nil
}
func RunBootstrap(templatesDir string, config *mcfgv1.ControllerConfig, pullSecretRaw []byte) ([]*mcfgv1.MachineConfig, error) {
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
	return getMachineConfigsForControllerConfig(templatesDir, config, pullSecretRaw)
}

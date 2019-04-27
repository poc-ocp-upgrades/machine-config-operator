package render

import (
	"fmt"
	"reflect"
	"time"
	"github.com/coreos/ignition/config/validate"
	"github.com/golang/glog"
	"github.com/openshift/machine-config-operator/lib/resourceapply"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"github.com/openshift/machine-config-operator/pkg/controller/common"
	mcfgclientset "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned"
	"github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/scheme"
	mcfginformersv1 "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions/machineconfiguration.openshift.io/v1"
	mcfglistersv1 "github.com/openshift/machine-config-operator/pkg/generated/listers/machineconfiguration.openshift.io/v1"
	"github.com/openshift/machine-config-operator/pkg/version"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
)

const (
	maxRetries	= 15
	renderDelay	= 5 * time.Second
)

var (
	controllerKind		= mcfgv1.SchemeGroupVersion.WithKind("MachineConfigPool")
	machineconfigKind	= mcfgv1.SchemeGroupVersion.WithKind("MachineConfig")
)

type Controller struct {
	client				mcfgclientset.Interface
	eventRecorder			record.EventRecorder
	syncHandler			func(mcp string) error
	enqueueMachineConfigPool	func(*mcfgv1.MachineConfigPool)
	mcpLister			mcfglistersv1.MachineConfigPoolLister
	mcLister			mcfglistersv1.MachineConfigLister
	mcpListerSynced			cache.InformerSynced
	mcListerSynced			cache.InformerSynced
	ccLister			mcfglistersv1.ControllerConfigLister
	ccListerSynced			cache.InformerSynced
	queue				workqueue.RateLimitingInterface
}

func New(mcpInformer mcfginformersv1.MachineConfigPoolInformer, mcInformer mcfginformersv1.MachineConfigInformer, ccInformer mcfginformersv1.ControllerConfigInformer, kubeClient clientset.Interface, mcfgClient mcfgclientset.Interface) *Controller {
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
	eventBroadcaster.StartRecordingToSink(&corev1.EventSinkImpl{Interface: kubeClient.CoreV1().Events("")})
	ctrl := &Controller{client: mcfgClient, eventRecorder: eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "machineconfigcontroller-rendercontroller"}), queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "machineconfigcontroller-rendercontroller")}
	mcpInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: ctrl.addMachineConfigPool, UpdateFunc: ctrl.updateMachineConfigPool, DeleteFunc: ctrl.deleteMachineConfigPool})
	mcInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: ctrl.addMachineConfig, UpdateFunc: ctrl.updateMachineConfig, DeleteFunc: ctrl.deleteMachineConfig})
	ctrl.syncHandler = ctrl.syncMachineConfigPool
	ctrl.enqueueMachineConfigPool = ctrl.enqueueDefault
	ctrl.mcpLister = mcpInformer.Lister()
	ctrl.mcLister = mcInformer.Lister()
	ctrl.mcpListerSynced = mcpInformer.Informer().HasSynced
	ctrl.mcListerSynced = mcInformer.Informer().HasSynced
	ctrl.ccLister = ccInformer.Lister()
	ctrl.ccListerSynced = ccInformer.Informer().HasSynced
	return ctrl
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
	defer utilruntime.HandleCrash()
	defer ctrl.queue.ShutDown()
	if !cache.WaitForCacheSync(stopCh, ctrl.mcpListerSynced, ctrl.mcListerSynced, ctrl.ccListerSynced) {
		return
	}
	glog.Info("Starting MachineConfigController-RenderController")
	defer glog.Info("Shutting down MachineConfigController-RenderController")
	for i := 0; i < workers; i++ {
		go wait.Until(ctrl.worker, time.Second, stopCh)
	}
	<-stopCh
}
func (ctrl *Controller) addMachineConfigPool(obj interface{}) {
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
	pool := obj.(*mcfgv1.MachineConfigPool)
	glog.V(4).Infof("Adding MachineConfigPool %s", pool.Name)
	ctrl.enqueueMachineConfigPool(pool)
}
func (ctrl *Controller) updateMachineConfigPool(old, cur interface{}) {
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
	oldPool := old.(*mcfgv1.MachineConfigPool)
	curPool := cur.(*mcfgv1.MachineConfigPool)
	glog.V(4).Infof("Updating MachineConfigPool %s", oldPool.Name)
	ctrl.enqueueMachineConfigPool(curPool)
}
func (ctrl *Controller) deleteMachineConfigPool(obj interface{}) {
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
	pool, ok := obj.(*mcfgv1.MachineConfigPool)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
			return
		}
		pool, ok = tombstone.Obj.(*mcfgv1.MachineConfigPool)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a MachineConfigPool %#v", obj))
			return
		}
	}
	glog.V(4).Infof("Deleting MachineConfigPool %s", pool.Name)
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
	mc := obj.(*mcfgv1.MachineConfig)
	if mc.DeletionTimestamp != nil {
		ctrl.deleteMachineConfig(mc)
		return
	}
	controllerRef := metav1.GetControllerOf(mc)
	if controllerRef != nil {
		if pool := ctrl.resolveControllerRef(controllerRef); pool != nil {
			glog.V(4).Infof("MachineConfig %s added", mc.Name)
			ctrl.enqueueMachineConfigPool(pool)
			return
		}
	}
	pools, err := ctrl.getPoolsForMachineConfig(mc)
	if err != nil {
		glog.Errorf("error finding pools for machineconfig: %v", err)
		return
	}
	glog.V(4).Infof("MachineConfig %s added", mc.Name)
	for _, p := range pools {
		ctrl.enqueueMachineConfigPool(p)
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
	oldMC := old.(*mcfgv1.MachineConfig)
	curMC := cur.(*mcfgv1.MachineConfig)
	curControllerRef := metav1.GetControllerOf(curMC)
	oldControllerRef := metav1.GetControllerOf(oldMC)
	controllerRefChanged := !reflect.DeepEqual(curControllerRef, oldControllerRef)
	if controllerRefChanged && oldControllerRef != nil {
		glog.Errorf("machineconfig has changed controller, not allowed.")
		return
	}
	if curControllerRef != nil {
		if pool := ctrl.resolveControllerRef(curControllerRef); pool != nil {
			glog.V(4).Infof("MachineConfig %s updated", curMC.Name)
			ctrl.enqueueMachineConfigPool(pool)
			return
		}
	}
	pools, err := ctrl.getPoolsForMachineConfig(curMC)
	if err != nil {
		glog.Errorf("error finding pools for machineconfig: %v", err)
		return
	}
	glog.V(4).Infof("MachineConfig %s updated", curMC.Name)
	for _, p := range pools {
		ctrl.enqueueMachineConfigPool(p)
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
	if controllerRef != nil {
		if pool := ctrl.resolveControllerRef(controllerRef); pool != nil {
			glog.V(4).Infof("MachineConfig %s deleted", mc.Name)
			ctrl.enqueueMachineConfigPool(pool)
			return
		}
	}
	pools, err := ctrl.getPoolsForMachineConfig(mc)
	if err != nil {
		glog.Errorf("error finding pools for machineconfig: %v", err)
		return
	}
	glog.V(4).Infof("MachineConfig %s deleted", mc.Name)
	for _, p := range pools {
		ctrl.enqueueMachineConfigPool(p)
	}
}
func (ctrl *Controller) resolveControllerRef(controllerRef *metav1.OwnerReference) *mcfgv1.MachineConfigPool {
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
	pool, err := ctrl.mcpLister.Get(controllerRef.Name)
	if err != nil {
		return nil
	}
	if pool.UID != controllerRef.UID {
		return nil
	}
	return pool
}
func (ctrl *Controller) getPoolsForMachineConfig(config *mcfgv1.MachineConfig) ([]*mcfgv1.MachineConfigPool, error) {
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
	if len(config.Labels) == 0 {
		return nil, fmt.Errorf("no MachineConfigPool found for MachineConfig %v because it has no labels", config.Name)
	}
	pList, err := ctrl.mcpLister.List(labels.Everything())
	if err != nil {
		return nil, err
	}
	var pools []*mcfgv1.MachineConfigPool
	for _, p := range pList {
		selector, err := metav1.LabelSelectorAsSelector(p.Spec.MachineConfigSelector)
		if err != nil {
			return nil, fmt.Errorf("invalid label selector: %v", err)
		}
		if selector.Empty() || !selector.Matches(labels.Set(config.Labels)) {
			continue
		}
		pools = append(pools, p)
	}
	if len(pools) == 0 {
		return nil, fmt.Errorf("could not find any MachineConfigPool set for MachineConfig %s with labels: %v", config.Name, config.Labels)
	}
	return pools, nil
}
func (ctrl *Controller) enqueue(pool *mcfgv1.MachineConfigPool) {
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
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(pool)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", pool, err))
		return
	}
	ctrl.queue.Add(key)
}
func (ctrl *Controller) enqueueRateLimited(pool *mcfgv1.MachineConfigPool) {
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
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(pool)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", pool, err))
		return
	}
	ctrl.queue.AddRateLimited(key)
}
func (ctrl *Controller) enqueueAfter(pool *mcfgv1.MachineConfigPool, after time.Duration) {
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
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(pool)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", pool, err))
		return
	}
	ctrl.queue.AddAfter(key, after)
}
func (ctrl *Controller) enqueueDefault(pool *mcfgv1.MachineConfigPool) {
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
	ctrl.enqueueAfter(pool, renderDelay)
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
	if err == nil {
		ctrl.queue.Forget(key)
		return
	}
	if ctrl.queue.NumRequeues(key) < maxRetries {
		glog.V(2).Infof("Error syncing machineconfigpool %v: %v", key, err)
		ctrl.queue.AddRateLimited(key)
		return
	}
	utilruntime.HandleError(err)
	glog.V(2).Infof("Dropping machineconfigpool %q out of the queue: %v", key, err)
	ctrl.queue.Forget(key)
	ctrl.queue.AddAfter(key, 1*time.Minute)
}
func (ctrl *Controller) syncMachineConfigPool(key string) error {
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
	glog.V(4).Infof("Started syncing machineconfigpool %q (%v)", key, startTime)
	defer func() {
		glog.V(4).Infof("Finished syncing machineconfigpool %q (%v)", key, time.Since(startTime))
	}()
	_, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	machineconfigpool, err := ctrl.mcpLister.Get(name)
	if errors.IsNotFound(err) {
		glog.V(2).Infof("MachineConfigPool %v has been deleted", key)
		return nil
	}
	if err != nil {
		return err
	}
	pool := machineconfigpool.DeepCopy()
	everything := metav1.LabelSelector{}
	if reflect.DeepEqual(pool.Spec.MachineConfigSelector, &everything) {
		ctrl.eventRecorder.Eventf(pool, v1.EventTypeWarning, "SelectingAll", "This machineconfigpool is selecting all machineconfigs. A non-empty selector is require.")
		return nil
	}
	selector, err := metav1.LabelSelectorAsSelector(pool.Spec.MachineConfigSelector)
	if err != nil {
		return err
	}
	if err := mcfgv1.IsControllerConfigCompleted(common.ControllerConfigName, ctrl.ccLister.Get); err != nil {
		return err
	}
	mcs, err := ctrl.mcLister.List(selector)
	if err != nil {
		return err
	}
	if len(mcs) == 0 {
		return ctrl.syncFailingStatus(pool, fmt.Errorf("no MachineConfigs found matching selector %v", selector))
	}
	if err := ctrl.syncGeneratedMachineConfig(pool, mcs); err != nil {
		return ctrl.syncFailingStatus(pool, err)
	}
	return ctrl.syncAvailableStatus(pool)
}
func (ctrl *Controller) syncAvailableStatus(pool *mcfgv1.MachineConfigPool) error {
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
	if mcfgv1.IsMachineConfigPoolConditionFalse(pool.Status.Conditions, mcfgv1.MachineConfigPoolRenderDegraded) {
		return nil
	}
	sdegraded := mcfgv1.NewMachineConfigPoolCondition(mcfgv1.MachineConfigPoolRenderDegraded, v1.ConditionFalse, "", "")
	mcfgv1.SetMachineConfigPoolCondition(&pool.Status, *sdegraded)
	if _, err := ctrl.client.MachineconfigurationV1().MachineConfigPools().UpdateStatus(pool); err != nil {
		return err
	}
	return nil
}
func (ctrl *Controller) syncFailingStatus(pool *mcfgv1.MachineConfigPool, err error) error {
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
	sdegraded := mcfgv1.NewMachineConfigPoolCondition(mcfgv1.MachineConfigPoolRenderDegraded, v1.ConditionTrue, fmt.Sprintf("Failed to render configuration for pool %s: %v", pool.Name, err), "")
	mcfgv1.SetMachineConfigPoolCondition(&pool.Status, *sdegraded)
	if _, updateErr := ctrl.client.MachineconfigurationV1().MachineConfigPools().UpdateStatus(pool); updateErr != nil {
		glog.Errorf("Error updating MachineConfigPool %s: %v", pool.Name, updateErr)
	}
	return err
}
func (ctrl *Controller) garbageCollectRenderedConfigs(pool *mcfgv1.MachineConfigPool) error {
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
	return nil
}
func (ctrl *Controller) syncGeneratedMachineConfig(pool *mcfgv1.MachineConfigPool, configs []*mcfgv1.MachineConfig) error {
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
	if len(configs) == 0 {
		return nil
	}
	cc, err := ctrl.ccLister.Get(common.ControllerConfigName)
	if err != nil {
		return err
	}
	generated, err := generateRenderedMachineConfig(pool, configs, cc)
	if err != nil {
		return err
	}
	source := []v1.ObjectReference{}
	for _, cfg := range configs {
		source = append(source, v1.ObjectReference{Kind: machineconfigKind.Kind, Name: cfg.GetName(), APIVersion: machineconfigKind.GroupVersion().String()})
	}
	_, err = ctrl.mcLister.Get(generated.Name)
	if apierrors.IsNotFound(err) {
		_, err = ctrl.client.MachineconfigurationV1().MachineConfigs().Create(generated)
		glog.V(2).Infof("Generated machineconfig %s from %d configs: %s", generated.Name, len(source), source)
	}
	if err != nil {
		return err
	}
	if pool.Status.Configuration.Name == generated.Name {
		_, _, err = resourceapply.ApplyMachineConfig(ctrl.client.MachineconfigurationV1(), generated)
		return err
	}
	pool.Status.Configuration.Name = generated.Name
	pool.Status.Configuration.Source = source
	_, err = ctrl.client.MachineconfigurationV1().MachineConfigPools().UpdateStatus(pool)
	if err != nil {
		return err
	}
	if err := ctrl.garbageCollectRenderedConfigs(pool); err != nil {
		return err
	}
	return nil
}
func generateRenderedMachineConfig(pool *mcfgv1.MachineConfigPool, configs []*mcfgv1.MachineConfig, cconfig *mcfgv1.ControllerConfig) (*mcfgv1.MachineConfig, error) {
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
	for _, config := range configs {
		rpt := validate.ValidateWithoutSource(reflect.ValueOf(config.Spec.Config.Ignition))
		if rpt.IsFatal() {
			return nil, fmt.Errorf("machine config: %v contains invalid ignition config: %v", config.ObjectMeta.Name, rpt)
		}
	}
	merged := mcfgv1.MergeMachineConfigs(configs, cconfig.Spec.OSImageURL)
	hashedName, err := getMachineConfigHashedName(pool, merged)
	if err != nil {
		return nil, err
	}
	oref := metav1.NewControllerRef(pool, controllerKind)
	merged.SetName(hashedName)
	merged.SetOwnerReferences([]metav1.OwnerReference{*oref})
	if merged.Annotations == nil {
		merged.Annotations = map[string]string{}
	}
	merged.Annotations[common.GeneratedByControllerVersionAnnotationKey] = version.Version.String()
	return merged, nil
}
func RunBootstrap(pools []*mcfgv1.MachineConfigPool, configs []*mcfgv1.MachineConfig, cconfig *mcfgv1.ControllerConfig) ([]*mcfgv1.MachineConfigPool, []*mcfgv1.MachineConfig, error) {
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
	var (
		opools		[]*mcfgv1.MachineConfigPool
		oconfigs	[]*mcfgv1.MachineConfig
	)
	for _, pool := range pools {
		pcs, err := getMachineConfigsForPool(pool, configs)
		if err != nil {
			return nil, nil, err
		}
		generated, err := generateRenderedMachineConfig(pool, pcs, cconfig)
		if err != nil {
			return nil, nil, err
		}
		pool.Status.Configuration.Name = generated.Name
		opools = append(opools, pool)
		oconfigs = append(oconfigs, generated)
	}
	return opools, oconfigs, nil
}
func getMachineConfigsForPool(pool *mcfgv1.MachineConfigPool, configs []*mcfgv1.MachineConfig) ([]*mcfgv1.MachineConfig, error) {
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
	selector, err := metav1.LabelSelectorAsSelector(pool.Spec.MachineConfigSelector)
	if err != nil {
		return nil, fmt.Errorf("invalid label selector: %v", err)
	}
	var out []*mcfgv1.MachineConfig
	if selector.Empty() {
		return out, nil
	}
	for idx, config := range configs {
		if selector.Matches(labels.Set(config.Labels)) {
			out = append(out, configs[idx])
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("couldn't find any MachineConfigs for pool: %v", pool.Name)
	}
	return out, nil
}

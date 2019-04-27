package render

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	ctrlcommon "github.com/openshift/machine-config-operator/pkg/controller/common"
	"github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/fake"
	informers "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/diff"
	utilrand "k8s.io/apimachinery/pkg/util/rand"
	k8sfake "k8s.io/client-go/kubernetes/fake"
	core "k8s.io/client-go/testing"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
)

var (
	alwaysReady	= func() bool {
		return true
	}
	noResyncPeriodFunc	= func() time.Duration {
		return 0
	}
)

type fixture struct {
	t		*testing.T
	client		*fake.Clientset
	mcpLister	[]*mcfgv1.MachineConfigPool
	mcLister	[]*mcfgv1.MachineConfig
	ccLister	[]*mcfgv1.ControllerConfig
	actions		[]core.Action
	objects		[]runtime.Object
}

func newFixture(t *testing.T) *fixture {
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
	f := &fixture{}
	f.t = t
	f.objects = []runtime.Object{}
	return f
}
func newMachineConfigPool(name string, selector *metav1.LabelSelector, currentMachineConfig string) *mcfgv1.MachineConfigPool {
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
	mcp := &mcfgv1.MachineConfigPool{TypeMeta: metav1.TypeMeta{APIVersion: mcfgv1.SchemeGroupVersion.String()}, ObjectMeta: metav1.ObjectMeta{Name: name, UID: types.UID(utilrand.String(5))}, Spec: mcfgv1.MachineConfigPoolSpec{MachineConfigSelector: selector}, Status: mcfgv1.MachineConfigPoolStatus{Configuration: mcfgv1.MachineConfigPoolStatusConfiguration{ObjectReference: corev1.ObjectReference{Name: currentMachineConfig}}}}
	sdegraded := mcfgv1.NewMachineConfigPoolCondition(mcfgv1.MachineConfigPoolRenderDegraded, corev1.ConditionFalse, "", "")
	mcfgv1.SetMachineConfigPoolCondition(&mcp.Status, *sdegraded)
	return mcp
}
func newMachineConfig(name string, labels map[string]string, osurl string, files []ignv2_2types.File) *mcfgv1.MachineConfig {
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
	if labels == nil {
		labels = map[string]string{}
	}
	ignCfg := ctrlcommon.NewIgnConfig()
	ignCfg.Storage.Files = files
	return &mcfgv1.MachineConfig{TypeMeta: metav1.TypeMeta{APIVersion: mcfgv1.SchemeGroupVersion.String()}, ObjectMeta: metav1.ObjectMeta{Name: name, Labels: labels, UID: types.UID(utilrand.String(5))}, Spec: mcfgv1.MachineConfigSpec{OSImageURL: osurl, Config: ignCfg}}
}
func (f *fixture) newController() *Controller {
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
	f.client = fake.NewSimpleClientset(f.objects...)
	i := informers.NewSharedInformerFactory(f.client, noResyncPeriodFunc())
	c := New(i.Machineconfiguration().V1().MachineConfigPools(), i.Machineconfiguration().V1().MachineConfigs(), i.Machineconfiguration().V1().ControllerConfigs(), k8sfake.NewSimpleClientset(), f.client)
	c.mcpListerSynced = alwaysReady
	c.mcListerSynced = alwaysReady
	c.ccListerSynced = alwaysReady
	c.eventRecorder = &record.FakeRecorder{}
	stopCh := make(chan struct{})
	defer close(stopCh)
	i.Start(stopCh)
	i.WaitForCacheSync(stopCh)
	for _, c := range f.ccLister {
		i.Machineconfiguration().V1().ControllerConfigs().Informer().GetIndexer().Add(c)
	}
	for _, c := range f.mcpLister {
		i.Machineconfiguration().V1().MachineConfigPools().Informer().GetIndexer().Add(c)
	}
	for _, m := range f.mcLister {
		i.Machineconfiguration().V1().MachineConfigs().Informer().GetIndexer().Add(m)
	}
	for _, m := range f.ccLister {
		i.Machineconfiguration().V1().ControllerConfigs().Informer().GetIndexer().Add(m)
	}
	return c
}
func (f *fixture) run(mcpname string) {
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
	f.runController(mcpname, false)
}
func (f *fixture) runExpectError(mcpname string) {
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
	f.runController(mcpname, true)
}
func (f *fixture) runController(mcpname string, expectError bool) {
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
	c := f.newController()
	err := c.syncHandler(mcpname)
	if !expectError && err != nil {
		f.t.Errorf("error syncing machineconfigpool: %v", err)
	} else if expectError && err == nil {
		f.t.Error("expected error syncing machineconfigpool, got nil")
	}
	actions := filterInformerActions(f.client.Actions())
	for i, action := range actions {
		if len(f.actions) < i+1 {
			f.t.Errorf("%d unexpected actions: %+v", len(actions)-len(f.actions), actions[i:])
			break
		}
		expectedAction := f.actions[i]
		checkAction(expectedAction, action, f.t)
	}
	if len(f.actions) > len(actions) {
		f.t.Errorf("%d additional expected actions:%+v", len(f.actions)-len(actions), f.actions[len(actions):])
	}
}
func checkAction(expected, actual core.Action, t *testing.T) {
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
	if !(expected.Matches(actual.GetVerb(), actual.GetResource().Resource) && actual.GetSubresource() == expected.GetSubresource()) {
		t.Errorf("Expected\n\t%#v\ngot\n\t%#v", expected, actual)
		return
	}
	if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
		t.Errorf("Action has wrong type. Expected: %t. Got: %t", expected, actual)
		return
	}
	switch a := actual.(type) {
	case core.CreateAction:
		e, _ := expected.(core.CreateAction)
		expObject := e.GetObject()
		object := a.GetObject()
		if !equality.Semantic.DeepEqual(expObject, object) {
			t.Errorf("Action %s %s has wrong object\nDiff:\n %s", a.GetVerb(), a.GetResource().Resource, diff.ObjectGoPrintDiff(expObject, object))
		}
	case core.UpdateAction:
		e, _ := expected.(core.UpdateAction)
		expObject := e.GetObject()
		object := a.GetObject()
		if !equality.Semantic.DeepEqual(expObject, object) {
			t.Errorf("Action %s %s has wrong object\nDiff:\n %s", a.GetVerb(), a.GetResource().Resource, diff.ObjectGoPrintDiff(expObject, object))
		}
	case core.PatchAction:
		e, _ := expected.(core.PatchAction)
		expPatch := e.GetPatch()
		patch := a.GetPatch()
		if !equality.Semantic.DeepEqual(expPatch, expPatch) {
			t.Errorf("Action %s %s has wrong patch\nDiff:\n %s", a.GetVerb(), a.GetResource().Resource, diff.ObjectGoPrintDiff(expPatch, patch))
		}
	}
}
func filterInformerActions(actions []core.Action) []core.Action {
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
	ret := []core.Action{}
	for _, action := range actions {
		if len(action.GetNamespace()) == 0 && (action.Matches("list", "machineconfigpools") || action.Matches("watch", "machineconfigpools") || action.Matches("list", "controllerconfigs") || action.Matches("watch", "controllerconfigs") || action.Matches("list", "machineconfigs") || action.Matches("watch", "machineconfigs")) {
			continue
		}
		ret = append(ret, action)
	}
	return ret
}
func (f *fixture) expectGetMachineConfigAction(config *mcfgv1.MachineConfig) {
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
	f.actions = append(f.actions, core.NewRootGetAction(schema.GroupVersionResource{Resource: "machineconfigs"}, config.Name))
}
func (f *fixture) expectCreateMachineConfigAction(config *mcfgv1.MachineConfig) {
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
	f.actions = append(f.actions, core.NewRootCreateAction(schema.GroupVersionResource{Resource: "machineconfigs"}, config))
}
func (f *fixture) expectPatchMachineConfigAction(config *mcfgv1.MachineConfig, patch []byte) {
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
	f.actions = append(f.actions, core.NewRootPatchAction(schema.GroupVersionResource{Resource: "machineconfigs"}, config.Name, types.MergePatchType, patch))
}
func (f *fixture) expectUpdateMachineConfigAction(config *mcfgv1.MachineConfig) {
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
	f.actions = append(f.actions, core.NewRootUpdateAction(schema.GroupVersionResource{Resource: "machineconfigs"}, config))
}
func (f *fixture) expectUpdateMachineConfigPoolStatus(pool *mcfgv1.MachineConfigPool) {
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
	f.actions = append(f.actions, core.NewRootUpdateSubresourceAction(schema.GroupVersionResource{Resource: "machineconfigpools"}, "status", pool))
}
func newControllerConfig(name string) *mcfgv1.ControllerConfig {
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
	return &mcfgv1.ControllerConfig{TypeMeta: metav1.TypeMeta{APIVersion: mcfgv1.SchemeGroupVersion.String()}, ObjectMeta: metav1.ObjectMeta{Name: name, UID: types.UID(utilrand.String(5))}, Spec: mcfgv1.ControllerConfigSpec{EtcdDiscoveryDomain: fmt.Sprintf("%s.tt.testing", name), OSImageURL: "dummy"}, Status: mcfgv1.ControllerConfigStatus{Conditions: []mcfgv1.ControllerConfigStatusCondition{{Type: mcfgv1.TemplateContollerCompleted, Status: corev1.ConditionTrue}}}}
}
func TestCreatesGeneratedMachineConfig(t *testing.T) {
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
	f := newFixture(t)
	mcp := newMachineConfigPool("test-cluster-master", metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "master"), "")
	files := []ignv2_2types.File{{Node: ignv2_2types.Node{Path: "/dummy/0"}}, {Node: ignv2_2types.Node{Path: "/dummy/1"}}}
	mcs := []*mcfgv1.MachineConfig{newMachineConfig("00-test-cluster-master", map[string]string{"node-role": "master"}, "dummy://", []ignv2_2types.File{files[0]}), newMachineConfig("05-extra-master", map[string]string{"node-role": "master"}, "dummy://1", []ignv2_2types.File{files[1]})}
	cc := newControllerConfig(ctrlcommon.ControllerConfigName)
	f.ccLister = append(f.ccLister, cc)
	f.mcpLister = append(f.mcpLister, mcp)
	f.objects = append(f.objects, mcp)
	f.mcLister = append(f.mcLister, mcs...)
	for idx := range mcs {
		f.objects = append(f.objects, mcs[idx])
	}
}
func TestIgnValidationGenerateRenderedMachineConfig(t *testing.T) {
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
	mcp := newMachineConfigPool("test-cluster-master", metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "master"), "")
	files := []ignv2_2types.File{{Node: ignv2_2types.Node{Path: "/dummy/0"}}, {Node: ignv2_2types.Node{Path: "/dummy/1"}}}
	mcs := []*mcfgv1.MachineConfig{newMachineConfig("00-test-cluster-master", map[string]string{"node-role": "master"}, "dummy://", []ignv2_2types.File{files[0]}), newMachineConfig("05-extra-master", map[string]string{"node-role": "master"}, "dummy://1", []ignv2_2types.File{files[1]})}
	cc := newControllerConfig(ctrlcommon.ControllerConfigName)
	_, err := generateRenderedMachineConfig(mcp, mcs, cc)
	if err != nil {
		t.Fatalf("expected no error. Got: %v", err)
	}
	mcs[1].Spec.Config.Ignition.Version = ""
	_, err = generateRenderedMachineConfig(mcp, mcs, cc)
	if err == nil {
		t.Fatalf("expected error. mcs contains a machine config with invalid ignconfig version")
	}
}
func TestUpdatesGeneratedMachineConfig(t *testing.T) {
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
	f := newFixture(t)
	mcp := newMachineConfigPool("test-cluster-master", metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "master"), "")
	files := []ignv2_2types.File{{Node: ignv2_2types.Node{Path: "/dummy/0"}}, {Node: ignv2_2types.Node{Path: "/dummy/1"}}}
	mcs := []*mcfgv1.MachineConfig{newMachineConfig("00-test-cluster-master", map[string]string{"node-role": "master"}, "dummy://", []ignv2_2types.File{files[0]}), newMachineConfig("05-extra-master", map[string]string{"node-role": "master"}, "dummy://1", []ignv2_2types.File{files[1]})}
	cc := newControllerConfig(ctrlcommon.ControllerConfigName)
	gmc, err := generateRenderedMachineConfig(mcp, mcs, cc)
	if err != nil {
		t.Fatal(err)
	}
	gmc.Spec.OSImageURL = "why-did-you-change-it"
	mcp.Status.Configuration.Name = gmc.Name
	f.ccLister = append(f.ccLister, cc)
	f.mcpLister = append(f.mcpLister, mcp)
	f.objects = append(f.objects, mcp)
	f.mcLister = append(f.mcLister, mcs...)
	for idx := range mcs {
		f.objects = append(f.objects, mcs[idx])
	}
	f.mcLister = append(f.mcLister, gmc)
	f.objects = append(f.objects, gmc)
	expmc, err := generateRenderedMachineConfig(mcp, mcs, cc)
	if err != nil {
		t.Fatal(err)
	}
	f.expectGetMachineConfigAction(expmc)
	f.expectUpdateMachineConfigAction(expmc)
	f.run(getKey(mcp, t))
}
func TestGenerateMachineConfigNoOverrideOSImageURL(t *testing.T) {
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
	mcp := newMachineConfigPool("test-cluster-master", metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "master"), "")
	mcs := []*mcfgv1.MachineConfig{newMachineConfig("00-test-cluster-master", map[string]string{"node-role": "master"}, "dummy-test-1", []ignv2_2types.File{}), newMachineConfig("00-test-cluster-master-0", map[string]string{"node-role": "master"}, "dummy-change", []ignv2_2types.File{})}
	cc := newControllerConfig(ctrlcommon.ControllerConfigName)
	gmc, err := generateRenderedMachineConfig(mcp, mcs, cc)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "dummy", gmc.Spec.OSImageURL)
}
func TestDoNothing(t *testing.T) {
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
	f := newFixture(t)
	mcp := newMachineConfigPool("test-cluster-master", metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "master"), "")
	files := []ignv2_2types.File{{Node: ignv2_2types.Node{Path: "/dummy/0"}}, {Node: ignv2_2types.Node{Path: "/dummy/1"}}}
	mcs := []*mcfgv1.MachineConfig{newMachineConfig("00-test-cluster-master", map[string]string{"node-role": "master"}, "dummy://", []ignv2_2types.File{files[0]}), newMachineConfig("05-extra-master", map[string]string{"node-role": "master"}, "dummy://1", []ignv2_2types.File{files[1]})}
	cc := newControllerConfig(ctrlcommon.ControllerConfigName)
	gmc, err := generateRenderedMachineConfig(mcp, mcs, cc)
	if err != nil {
		t.Fatal(err)
	}
	mcp.Status.Configuration.Name = gmc.Name
	f.ccLister = append(f.ccLister, cc)
	f.mcpLister = append(f.mcpLister, mcp)
	f.objects = append(f.objects, mcp)
	f.mcLister = append(f.mcLister, mcs...)
	for idx := range mcs {
		f.objects = append(f.objects, mcs[idx])
	}
	f.mcLister = append(f.mcLister, gmc)
	f.objects = append(f.objects, gmc)
	f.expectGetMachineConfigAction(gmc)
	f.run(getKey(mcp, t))
}
func TestGetMachineConfigsForPool(t *testing.T) {
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
	masterPool := newMachineConfigPool("test-cluster-master", metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "master"), "")
	files := []ignv2_2types.File{{Node: ignv2_2types.Node{Path: "/dummy/0"}}, {Node: ignv2_2types.Node{Path: "/dummy/1"}}, {Node: ignv2_2types.Node{Path: "/dummy/2"}}}
	mcs := []*mcfgv1.MachineConfig{newMachineConfig("00-test-cluster-master", map[string]string{"node-role": "master"}, "dummy://", []ignv2_2types.File{files[0]}), newMachineConfig("05-extra-master", map[string]string{"node-role": "master"}, "dummy://1", []ignv2_2types.File{files[1]}), newMachineConfig("00-test-cluster-worker", map[string]string{"node-role": "worker"}, "dummy://2", []ignv2_2types.File{files[2]})}
	masterConfigs, err := getMachineConfigsForPool(masterPool, mcs)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if len(masterConfigs) != 2 {
		t.Fatalf("expected to select 2 configs for pool master got: %v", len(masterConfigs))
	}
	workerPool := newMachineConfigPool("test-cluster-worker", metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "worker"), "")
	_, err = getMachineConfigsForPool(workerPool, mcs[:2])
	if err == nil {
		t.Fatalf("expected error, no worker configs found")
	}
}
func getKey(config *mcfgv1.MachineConfigPool, t *testing.T) string {
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
		t.Errorf("Unexpected error getting key for config %v: %v", config.Name, err)
		return ""
	}
	return key
}
func TestMachineConfigsNoBailWithoutPool(t *testing.T) {
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
	f := newFixture(t)
	mc := newMachineConfig("00-test-cluster-worker", map[string]string{"node-role": "worker"}, "dummy://2", []ignv2_2types.File{})
	oref := metav1.NewControllerRef(newControllerConfig("test"), mcfgv1.SchemeGroupVersion.WithKind("ControllerConfig"))
	mc.SetOwnerReferences([]metav1.OwnerReference{*oref})
	mcp := newMachineConfigPool("test-cluster-master", metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "worker"), "")
	f.mcpLister = append(f.mcpLister, mcp)
	c := f.newController()
	queue := []*mcfgv1.MachineConfigPool{}
	c.enqueueMachineConfigPool = func(mcp *mcfgv1.MachineConfigPool) {
		queue = append(queue, mcp)
	}
	c.addMachineConfig(mc)
	c.updateMachineConfig(mc, mc)
	c.deleteMachineConfig(mc)
	require.Len(t, queue, 3)
}

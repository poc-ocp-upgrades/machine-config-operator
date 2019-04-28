package containerruntimeconfig

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"k8s.io/apimachinery/pkg/api/resource"
	"github.com/golang/glog"
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
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	apicfgv1 "github.com/openshift/api/config/v1"
	fakeconfigv1client "github.com/openshift/client-go/config/clientset/versioned/fake"
	configv1informer "github.com/openshift/client-go/config/informers/externalversions"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"github.com/openshift/machine-config-operator/pkg/controller/common"
	"github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/fake"
	informers "github.com/openshift/machine-config-operator/pkg/generated/informers/externalversions"
)

var (
	alwaysReady	= func() bool {
		return true
	}
	noResyncPeriodFunc	= func() time.Duration {
		return 0
	}
)

const (
	templateDir = "../../../templates"
)

type fixture struct {
	t		*testing.T
	client		*fake.Clientset
	imgClient	*fakeconfigv1client.Clientset
	ccLister	[]*mcfgv1.ControllerConfig
	mcpLister	[]*mcfgv1.MachineConfigPool
	mccrLister	[]*mcfgv1.ContainerRuntimeConfig
	imgLister	[]*apicfgv1.Image
	cvLister	[]*apicfgv1.ClusterVersion
	actions		[]core.Action
	objects		[]runtime.Object
	imgObjects	[]runtime.Object
}

func newFixture(t *testing.T) *fixture {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f := &fixture{}
	f.t = t
	f.objects = []runtime.Object{}
	return f
}
func (f *fixture) validateActions() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	actions := filterInformerActions(f.client.Actions())
	for i, action := range actions {
		glog.Infof("Action: %v", action)
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
func newMachineConfig(name string, labels map[string]string, osurl string, files []ignv2_2types.File) *mcfgv1.MachineConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if labels == nil {
		labels = map[string]string{}
	}
	return &mcfgv1.MachineConfig{TypeMeta: metav1.TypeMeta{APIVersion: mcfgv1.SchemeGroupVersion.String()}, ObjectMeta: metav1.ObjectMeta{Name: name, Labels: labels, UID: types.UID(utilrand.String(5))}, Spec: mcfgv1.MachineConfigSpec{OSImageURL: osurl, Config: ignv2_2types.Config{Storage: ignv2_2types.Storage{Files: files}}}}
}
func newControllerConfig(name, platform string) *mcfgv1.ControllerConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := &mcfgv1.ControllerConfig{TypeMeta: metav1.TypeMeta{APIVersion: mcfgv1.SchemeGroupVersion.String()}, ObjectMeta: metav1.ObjectMeta{Name: name, UID: types.UID(utilrand.String(5))}, Spec: mcfgv1.ControllerConfigSpec{EtcdDiscoveryDomain: fmt.Sprintf("%s.tt.testing", name), Platform: platform}}
	return cc
}
func newMachineConfigPool(name string, labels map[string]string, selector *metav1.LabelSelector, currentMachineConfig string) *mcfgv1.MachineConfigPool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &mcfgv1.MachineConfigPool{TypeMeta: metav1.TypeMeta{APIVersion: mcfgv1.SchemeGroupVersion.String()}, ObjectMeta: metav1.ObjectMeta{Name: name, Labels: labels, UID: types.UID(utilrand.String(5))}, Spec: mcfgv1.MachineConfigPoolSpec{MachineConfigSelector: selector}, Status: mcfgv1.MachineConfigPoolStatus{Configuration: mcfgv1.MachineConfigPoolStatusConfiguration{ObjectReference: corev1.ObjectReference{Name: currentMachineConfig}}}}
}
func newContainerRuntimeConfig(name string, ctrconf *mcfgv1.ContainerRuntimeConfiguration, selector *metav1.LabelSelector) *mcfgv1.ContainerRuntimeConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &mcfgv1.ContainerRuntimeConfig{TypeMeta: metav1.TypeMeta{APIVersion: mcfgv1.SchemeGroupVersion.String()}, ObjectMeta: metav1.ObjectMeta{Name: name, UID: types.UID(utilrand.String(5)), Generation: 1}, Spec: mcfgv1.ContainerRuntimeConfigSpec{ContainerRuntimeConfig: ctrconf, MachineConfigPoolSelector: selector}, Status: mcfgv1.ContainerRuntimeConfigStatus{}}
}
func newImageConfig(name string, regconf *apicfgv1.RegistrySources) *apicfgv1.Image {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &apicfgv1.Image{TypeMeta: metav1.TypeMeta{APIVersion: apicfgv1.SchemeGroupVersion.String()}, ObjectMeta: metav1.ObjectMeta{Name: name, UID: types.UID(utilrand.String(5)), Generation: 1}, Spec: apicfgv1.ImageSpec{RegistrySources: *regconf}}
}
func newClusterVersionConfig(name, desiredImage string) *apicfgv1.ClusterVersion {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &apicfgv1.ClusterVersion{TypeMeta: metav1.TypeMeta{APIVersion: apicfgv1.SchemeGroupVersion.String()}, ObjectMeta: metav1.ObjectMeta{Name: name, UID: types.UID(utilrand.String(5)), Generation: 1}, Status: apicfgv1.ClusterVersionStatus{Desired: apicfgv1.Update{Image: desiredImage}}}
}
func (f *fixture) newController() *Controller {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.client = fake.NewSimpleClientset(f.objects...)
	f.imgClient = fakeconfigv1client.NewSimpleClientset(f.imgObjects...)
	i := informers.NewSharedInformerFactory(f.client, noResyncPeriodFunc())
	ci := configv1informer.NewSharedInformerFactory(f.imgClient, noResyncPeriodFunc())
	c := New(templateDir, i.Machineconfiguration().V1().MachineConfigPools(), i.Machineconfiguration().V1().ControllerConfigs(), i.Machineconfiguration().V1().ContainerRuntimeConfigs(), ci.Config().V1().Images(), ci.Config().V1().ClusterVersions(), k8sfake.NewSimpleClientset(), f.client, f.imgClient)
	c.patchContainerRuntimeConfigsFunc = func(name string, patch []byte) error {
		f.client.Invokes(core.NewRootPatchAction(schema.GroupVersionResource{Version: "v1", Group: "machineconfiguration.openshift.io", Resource: "containerruntimeconfigs"}, name, types.MergePatchType, patch), nil)
		return nil
	}
	c.mcpListerSynced = alwaysReady
	c.mccrListerSynced = alwaysReady
	c.ccListerSynced = alwaysReady
	c.imgListerSynced = alwaysReady
	c.clusterVersionListerSynced = alwaysReady
	c.eventRecorder = &record.FakeRecorder{}
	stopCh := make(chan struct{})
	defer close(stopCh)
	i.Start(stopCh)
	i.WaitForCacheSync(stopCh)
	ci.Start(stopCh)
	i.WaitForCacheSync(stopCh)
	for _, c := range f.ccLister {
		i.Machineconfiguration().V1().ControllerConfigs().Informer().GetIndexer().Add(c)
	}
	for _, c := range f.mcpLister {
		i.Machineconfiguration().V1().MachineConfigPools().Informer().GetIndexer().Add(c)
	}
	for _, c := range f.mccrLister {
		i.Machineconfiguration().V1().ContainerRuntimeConfigs().Informer().GetIndexer().Add(c)
	}
	for _, c := range f.imgLister {
		ci.Config().V1().Images().Informer().GetIndexer().Add(c)
	}
	for _, c := range f.cvLister {
		ci.Config().V1().ClusterVersions().Informer().GetIndexer().Add(c)
	}
	return c
}
func (f *fixture) run(mcpname string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.runController(mcpname, false)
}
func (f *fixture) runExpectError(mcpname string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.runController(mcpname, true)
}
func (f *fixture) runController(mcpname string, expectError bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := f.newController()
	err := c.syncImgHandler(mcpname)
	if !expectError && err != nil {
		f.t.Errorf("error syncing image config: %v", err)
	} else if expectError && err == nil {
		f.t.Error("expected error syncing image config, got nil")
	}
	err = c.syncHandler(mcpname)
	if !expectError && err != nil {
		f.t.Errorf("error syncing containerruntimeconfigs: %v", err)
	} else if expectError && err == nil {
		f.t.Error("expected error syncing containerruntimeconfigs, got nil")
	}
	f.validateActions()
}
func filterInformerActions(actions []core.Action) []core.Action {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := []core.Action{}
	for _, action := range actions {
		if len(action.GetNamespace()) == 0 && (action.Matches("list", "machineconfigpools") || action.Matches("watch", "machineconfigpools") || action.Matches("list", "controllerconfigs") || action.Matches("watch", "controllerconfigs") || action.Matches("list", "containerruntimeconfigs") || action.Matches("watch", "containerruntimeconfigs") || action.Matches("list", "machineconfigs") || action.Matches("watch", "machineconfigs")) {
			continue
		}
		ret = append(ret, action)
	}
	return ret
}
func checkAction(expected, actual core.Action, t *testing.T) {
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
		if a.GetVerb() != e.GetVerb() || a.GetResource().Resource != e.GetResource().Resource {
			t.Errorf("Action %s:%s has wrong Resource %s:%s", a.GetVerb(), e.GetVerb(), a.GetResource().Resource, e.GetResource().Resource)
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
func (f *fixture) expectGetContainerRuntimeConfigAction(config *mcfgv1.ContainerRuntimeConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.actions = append(f.actions, core.NewRootGetAction(schema.GroupVersionResource{Resource: "containerruntimeconfigs"}, config.Name))
}
func (f *fixture) expectGetMachineConfigAction(config *mcfgv1.MachineConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.actions = append(f.actions, core.NewRootGetAction(schema.GroupVersionResource{Resource: "machineconfigs"}, config.Name))
}
func (f *fixture) expectCreateMachineConfigAction(config *mcfgv1.MachineConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.actions = append(f.actions, core.NewRootCreateAction(schema.GroupVersionResource{Resource: "machineconfigs"}, config))
}
func (f *fixture) expectUpdateMachineConfigAction(config *mcfgv1.MachineConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.actions = append(f.actions, core.NewRootUpdateAction(schema.GroupVersionResource{Resource: "machineconfigs"}, config))
}
func (f *fixture) expectPatchContainerRuntimeConfig(config *mcfgv1.ContainerRuntimeConfig, patch []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.actions = append(f.actions, core.NewRootPatchAction(schema.GroupVersionResource{Version: "v1", Group: "machineconfiguration.openshift.io", Resource: "containerruntimeconfigs"}, config.Name, types.MergePatchType, patch))
}
func (f *fixture) expectUpdateContainerRuntimeConfig(config *mcfgv1.ContainerRuntimeConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.actions = append(f.actions, core.NewRootUpdateSubresourceAction(schema.GroupVersionResource{Version: "v1", Group: "machineconfiguration.openshift.io", Resource: "containerruntimeconfigs"}, "status", config))
}

var ctrcfgPatchBytes = []uint8{0x7b, 0x22, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x5b, 0x22, 0x39, 0x39, 0x2d, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x73, 0x78, 0x32, 0x76, 0x72, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x5d, 0x7d, 0x7d}

func TestContainerRuntimeConfigCreate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, platform := range []string{"aws", "none", "unrecognized"} {
		t.Run(platform, func(t *testing.T) {
			f := newFixture(t)
			cc := newControllerConfig(common.ControllerConfigName, platform)
			mcp := newMachineConfigPool("master", map[string]string{"custom-crio": "my-config"}, metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "master"), "v0")
			mcp2 := newMachineConfigPool("worker", map[string]string{"custom-crio": "storage-config"}, metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "worker"), "v0")
			ctrcfg1 := newContainerRuntimeConfig("set-log-level", &mcfgv1.ContainerRuntimeConfiguration{LogLevel: "debug", LogSizeMax: resource.MustParse("9k"), OverlaySize: resource.MustParse("3G")}, metav1.AddLabelToSelector(&metav1.LabelSelector{}, "custom-crio", "my-config"))
			mcs1 := newMachineConfig(getManagedKeyCtrCfg(mcp, ctrcfg1), map[string]string{"node-role": "master"}, "dummy://", []ignv2_2types.File{{}})
			f.ccLister = append(f.ccLister, cc)
			f.mcpLister = append(f.mcpLister, mcp)
			f.mcpLister = append(f.mcpLister, mcp2)
			f.mccrLister = append(f.mccrLister, ctrcfg1)
			f.objects = append(f.objects, ctrcfg1)
			f.expectGetMachineConfigAction(mcs1)
			f.expectUpdateContainerRuntimeConfig(ctrcfg1)
			f.expectUpdateContainerRuntimeConfig(ctrcfg1)
			f.expectCreateMachineConfigAction(mcs1)
			f.expectPatchContainerRuntimeConfig(ctrcfg1, ctrcfgPatchBytes)
			f.expectUpdateContainerRuntimeConfig(ctrcfg1)
			f.run(getKey(ctrcfg1, t))
		})
	}
}
func TestContainerRuntimeConfigUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, platform := range []string{"aws", "none", "unrecognized"} {
		t.Run(platform, func(t *testing.T) {
			f := newFixture(t)
			cc := newControllerConfig(common.ControllerConfigName, platform)
			mcp := newMachineConfigPool("master", map[string]string{"custom-crio": "my-config"}, metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "master"), "v0")
			mcp2 := newMachineConfigPool("worker", map[string]string{"custom-crio": "storage-config"}, metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "worker"), "v0")
			ctrcfg1 := newContainerRuntimeConfig("set-log-level", &mcfgv1.ContainerRuntimeConfiguration{LogLevel: "debug", LogSizeMax: resource.MustParse("9k"), OverlaySize: resource.MustParse("3G")}, metav1.AddLabelToSelector(&metav1.LabelSelector{}, "custom-crio", "my-config"))
			mcs := newMachineConfig(getManagedKeyCtrCfg(mcp, ctrcfg1), map[string]string{"node-role": "master"}, "dummy://", []ignv2_2types.File{{}})
			f.ccLister = append(f.ccLister, cc)
			f.mcpLister = append(f.mcpLister, mcp)
			f.mcpLister = append(f.mcpLister, mcp2)
			f.mccrLister = append(f.mccrLister, ctrcfg1)
			f.objects = append(f.objects, ctrcfg1)
			f.expectGetMachineConfigAction(mcs)
			f.expectUpdateContainerRuntimeConfig(ctrcfg1)
			f.expectUpdateContainerRuntimeConfig(ctrcfg1)
			f.expectCreateMachineConfigAction(mcs)
			f.expectPatchContainerRuntimeConfig(ctrcfg1, ctrcfgPatchBytes)
			f.expectUpdateContainerRuntimeConfig(ctrcfg1)
			c := f.newController()
			stopCh := make(chan struct{})
			err := c.syncHandler(getKey(ctrcfg1, t))
			if err != nil {
				t.Errorf("syncHandler returned %v", err)
			}
			f.validateActions()
			close(stopCh)
			f = newFixture(t)
			ctrcfgUpdate := ctrcfg1.DeepCopy()
			ctrcfgUpdate.Spec.ContainerRuntimeConfig.LogLevel = "warn"
			f.ccLister = append(f.ccLister, cc)
			f.mcpLister = append(f.mcpLister, mcp)
			f.mcpLister = append(f.mcpLister, mcp2)
			f.mccrLister = append(f.mccrLister, ctrcfg1)
			f.objects = append(f.objects, mcs, ctrcfgUpdate)
			c = f.newController()
			stopCh = make(chan struct{})
			glog.Info("Applying update")
			err = c.syncHandler(getKey(ctrcfgUpdate, t))
			if err != nil {
				t.Errorf("syncHandler returned: %v", err)
			}
			f.expectGetMachineConfigAction(mcs)
			f.expectUpdateContainerRuntimeConfig(ctrcfgUpdate)
			f.expectUpdateContainerRuntimeConfig(ctrcfgUpdate)
			f.expectUpdateMachineConfigAction(mcs)
			f.expectPatchContainerRuntimeConfig(ctrcfgUpdate, ctrcfgPatchBytes)
			f.expectUpdateContainerRuntimeConfig(ctrcfgUpdate)
			f.validateActions()
			close(stopCh)
		})
	}
}
func TestImageConfigCreate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, platform := range []string{"aws", "none", "unrecognized"} {
		t.Run(platform, func(t *testing.T) {
			f := newFixture(t)
			cc := newControllerConfig(common.ControllerConfigName, platform)
			mcp := newMachineConfigPool("master", map[string]string{"custom-crio": "my-config"}, metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "master"), "v0")
			mcp2 := newMachineConfigPool("worker", map[string]string{"custom-crio": "storage-config"}, metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "worker"), "v0")
			imgcfg1 := newImageConfig("cluster", &apicfgv1.RegistrySources{InsecureRegistries: []string{"blah.io"}})
			cvcfg1 := newClusterVersionConfig("version", "test.io/myuser/myimage:test")
			mcs1 := newMachineConfig(getManagedKeyReg(mcp, imgcfg1), map[string]string{"node-role": "master"}, "dummy://", []ignv2_2types.File{{}})
			mcs2 := newMachineConfig(getManagedKeyReg(mcp2, imgcfg1), map[string]string{"node-role": "worker"}, "dummy://", []ignv2_2types.File{{}})
			f.ccLister = append(f.ccLister, cc)
			f.mcpLister = append(f.mcpLister, mcp)
			f.mcpLister = append(f.mcpLister, mcp2)
			f.imgLister = append(f.imgLister, imgcfg1)
			f.cvLister = append(f.cvLister, cvcfg1)
			f.imgObjects = append(f.imgObjects, imgcfg1)
			f.expectGetMachineConfigAction(mcs1)
			f.expectCreateMachineConfigAction(mcs1)
			f.expectGetMachineConfigAction(mcs2)
			f.expectCreateMachineConfigAction(mcs2)
			f.run("cluster")
		})
	}
}
func TestImageConfigUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, platform := range []string{"aws", "none", "unrecognized"} {
		t.Run(platform, func(t *testing.T) {
			f := newFixture(t)
			cc := newControllerConfig(common.ControllerConfigName, platform)
			mcp := newMachineConfigPool("master", map[string]string{"custom-crio": "my-config"}, metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "master"), "v0")
			mcp2 := newMachineConfigPool("worker", map[string]string{"custom-crio": "storage-config"}, metav1.AddLabelToSelector(&metav1.LabelSelector{}, "node-role", "worker"), "v0")
			imgcfg1 := newImageConfig("cluster", &apicfgv1.RegistrySources{InsecureRegistries: []string{"blah.io"}})
			cvcfg1 := newClusterVersionConfig("version", "test.io/myuser/myimage:test")
			mcs1 := newMachineConfig(getManagedKeyReg(mcp, imgcfg1), map[string]string{"node-role": "master"}, "dummy://", []ignv2_2types.File{{}})
			mcs2 := newMachineConfig(getManagedKeyReg(mcp2, imgcfg1), map[string]string{"node-role": "worker"}, "dummy://", []ignv2_2types.File{{}})
			f.ccLister = append(f.ccLister, cc)
			f.mcpLister = append(f.mcpLister, mcp)
			f.mcpLister = append(f.mcpLister, mcp2)
			f.imgLister = append(f.imgLister, imgcfg1)
			f.cvLister = append(f.cvLister, cvcfg1)
			f.imgObjects = append(f.imgObjects, imgcfg1)
			f.expectGetMachineConfigAction(mcs1)
			f.expectCreateMachineConfigAction(mcs1)
			f.expectGetMachineConfigAction(mcs2)
			f.expectCreateMachineConfigAction(mcs2)
			c := f.newController()
			stopCh := make(chan struct{})
			err := c.syncImgHandler("cluster")
			if err != nil {
				t.Errorf("syncImgHandler returned %v", err)
			}
			f.validateActions()
			close(stopCh)
			f = newFixture(t)
			imgcfgUpdate := imgcfg1.DeepCopy()
			imgcfgUpdate.Spec.RegistrySources.InsecureRegistries = []string{"test.io"}
			f.ccLister = append(f.ccLister, cc)
			f.mcpLister = append(f.mcpLister, mcp)
			f.mcpLister = append(f.mcpLister, mcp2)
			f.imgLister = append(f.imgLister, imgcfg1)
			f.cvLister = append(f.cvLister, cvcfg1)
			f.imgObjects = append(f.imgObjects, imgcfg1)
			f.objects = append(f.objects, mcs1, mcs2)
			c = f.newController()
			stopCh = make(chan struct{})
			glog.Info("Applying update")
			err = c.syncImgHandler("")
			if err != nil {
				t.Errorf("syncImgHandler returned: %v", err)
			}
			f.expectGetMachineConfigAction(mcs1)
			f.expectUpdateMachineConfigAction(mcs1)
			f.expectGetMachineConfigAction(mcs2)
			f.expectUpdateMachineConfigAction(mcs2)
			f.validateActions()
			close(stopCh)
		})
	}
}
func TestRegistriesValidation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	failureTests := []struct {
		name	string
		config	*apicfgv1.RegistrySources
	}{{name: "adding registry used by payload to blocked registries", config: &apicfgv1.RegistrySources{BlockedRegistries: []string{"blah.io", "docker.io"}, InsecureRegistries: []string{"test.io"}}}}
	successTests := []struct {
		name	string
		config	*apicfgv1.RegistrySources
	}{{name: "adding registry used by payload to blocked registries", config: &apicfgv1.RegistrySources{BlockedRegistries: []string{"docker.io"}, InsecureRegistries: []string{"blah.io"}}}}
	for _, test := range failureTests {
		imgcfg := newImageConfig(test.name, test.config)
		cvcfg := newClusterVersionConfig("version", "blah.io/myuser/myimage:test")
		insecure, blocked, err := getValidRegistries(&cvcfg.Status, &imgcfg.Spec)
		if err == nil {
			t.Errorf("%s: failed", test.name)
		}
		for _, reg := range blocked {
			if reg == "blah.io" {
				t.Errorf("%s: failed to filter out registry being used by payload", test.name)
			}
		}
		if len(insecure) == 0 {
			t.Errorf("%s: failed to copy over the insecure registries from the image spec", test.name)
		}
	}
	for _, test := range successTests {
		imgcfg := newImageConfig(test.name, test.config)
		cvcfg := newClusterVersionConfig("version", "blah.io/myuser/myimage:test")
		insecure, blocked, err := getValidRegistries(&cvcfg.Status, &imgcfg.Spec)
		if err != nil {
			t.Errorf("%s: failed", test.name)
		}
		for _, reg := range blocked {
			if reg == "blah.io" {
				t.Errorf("%s: failed to filter out registry being used by payload", test.name)
			}
		}
		if len(insecure) == 0 {
			t.Errorf("%s: failed to copy over the insecure registries from the image spec", test.name)
		}
	}
}
func TestContainerRuntimeConfigOptions(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	failureTests := []struct {
		name	string
		config	*mcfgv1.ContainerRuntimeConfiguration
	}{{name: "invalid value of pids limit", config: &mcfgv1.ContainerRuntimeConfiguration{PidsLimit: 10}}, {name: "inalid value of max log size", config: &mcfgv1.ContainerRuntimeConfiguration{LogSizeMax: resource.MustParse("3k")}}, {name: "inalid value of log level", config: &mcfgv1.ContainerRuntimeConfiguration{LogLevel: "invalid"}}}
	successTests := []struct {
		name	string
		config	*mcfgv1.ContainerRuntimeConfiguration
	}{{name: "valid pids limit", config: &mcfgv1.ContainerRuntimeConfiguration{PidsLimit: 2048}}, {name: "valid max log size", config: &mcfgv1.ContainerRuntimeConfiguration{LogSizeMax: resource.MustParse("10k")}}, {name: "valid log level", config: &mcfgv1.ContainerRuntimeConfiguration{LogLevel: "debug"}}}
	for _, test := range failureTests {
		ctrcfg := newContainerRuntimeConfig(test.name, test.config, metav1.AddLabelToSelector(&metav1.LabelSelector{}, "", ""))
		err := validateUserContainerRuntimeConfig(ctrcfg)
		if err == nil {
			t.Errorf("%s: failed", test.name)
		}
	}
	for _, test := range successTests {
		ctrcfg := newContainerRuntimeConfig(test.name, test.config, metav1.AddLabelToSelector(&metav1.LabelSelector{}, "", ""))
		err := validateUserContainerRuntimeConfig(ctrcfg)
		if err != nil {
			t.Errorf("%s: failed with %v. should have succeeded", test.name, err)
		}
	}
}
func getKey(config *mcfgv1.ContainerRuntimeConfig, t *testing.T) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(config)
	if err != nil {
		t.Errorf("Unexpected error getting key for config %v: %v", config.Name, err)
		return ""
	}
	return key
}

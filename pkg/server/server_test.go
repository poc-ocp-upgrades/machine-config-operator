package server

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"reflect"
	"testing"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	yaml "github.com/ghodss/yaml"
	"github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	daemonconsts "github.com/openshift/machine-config-operator/pkg/daemon/constants"
	"github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/fake"
)

const (
	testPool	= "test-pool"
	testConfig	= "test-config"
	testDir		= "./testdata"
)

var (
	testKubeConfig = fmt.Sprintf("%s/kubeconfig", testDir)
)

func TestStringDecode(t *testing.T) {
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
	inp := "data:,Hello%2C%20world!"
	exp := "Hello, world!"
	dec, err := getDecodedContent(inp)
	if err != nil {
		t.Errorf("expected error to be nil, received: %v", err)
	}
	if exp != dec {
		t.Errorf("string decode failed. exp: %s, got: %s", exp, dec)
	}
}
func TestStringEncode(t *testing.T) {
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
	inp := "Hello, world!"
	exp := "data:,Hello%2C%20world!"
	enc := getEncodedContent(inp)
	if exp != enc {
		t.Errorf("string encode failed. exp: %s, got: %s", exp, enc)
	}
}
func TestBootstrapServer(t *testing.T) {
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
	mp, err := getTestMachineConfigPool()
	if err != nil {
		t.Fatal(err)
	}
	mcPath := filepath.Join(testDir, "machine-configs", testConfig+".yaml")
	mcData, err := ioutil.ReadFile(mcPath)
	if err != nil {
		t.Fatalf("unexpected error while reading machine-config: %s, err: %v", mcPath, err)
	}
	mc := new(v1.MachineConfig)
	err = yaml.Unmarshal([]byte(mcData), mc)
	if err != nil {
		t.Fatalf("unexpected error while unmarshaling machine-config: %s, err: %v", mcPath, err)
	}
	kc, _, err := getKubeConfigContent(t)
	if err != nil {
		t.Fatal(err)
	}
	appendFileToIgnition(&mc.Spec.Config, defaultMachineKubeConfPath, string(kc))
	anno, err := getNodeAnnotation(mp.Status.Configuration.Name)
	if err != nil {
		t.Fatalf("unexpected error while creating annotations err: %v", err)
	}
	appendFileToIgnition(&mc.Spec.Config, daemonconsts.InitialNodeAnnotationsFilePath, anno)
	bs := &bootstrapServer{serverBaseDir: testDir, kubeconfigFunc: func() ([]byte, []byte, error) {
		return getKubeConfigContent(t)
	}}
	if err != nil {
		t.Fatal(err)
	}
	res, err := bs.GetConfig(poolRequest{machineConfigPool: testPool})
	if err != nil {
		t.Fatalf("expected err to be nil, received: %v", err)
	}
	validateIgnitionFiles(t, res.Storage.Files, mc.Spec.Config.Storage.Files)
	validateIgnitionSystemd(t, res.Systemd.Units, mc.Spec.Config.Systemd.Units)
}
func TestClusterServer(t *testing.T) {
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
	mp, err := getTestMachineConfigPool()
	if err != nil {
		t.Fatal(err)
	}
	mcPath := filepath.Join(testDir, "machine-configs", testConfig+".yaml")
	mcData, err := ioutil.ReadFile(mcPath)
	if err != nil {
		t.Fatalf("unexpected error while reading machine-config: %s, err: %v", mcPath, err)
	}
	origMC := new(v1.MachineConfig)
	err = yaml.Unmarshal([]byte(mcData), origMC)
	if err != nil {
		t.Fatalf("unexpected error while unmarshaling machine-config: %s, err: %v", mcPath, err)
	}
	cs := fake.NewSimpleClientset()
	_, err = cs.MachineconfigurationV1().MachineConfigPools().Create(mp)
	if err != nil {
		t.Logf("err: %v", err)
	}
	_, err = cs.MachineconfigurationV1().MachineConfigs().Create(origMC)
	if err != nil {
		t.Logf("err: %v", err)
	}
	csc := &clusterServer{machineClient: cs.MachineconfigurationV1(), kubeconfigFunc: func() ([]byte, []byte, error) {
		return getKubeConfigContent(t)
	}}
	mc := new(v1.MachineConfig)
	err = yaml.Unmarshal([]byte(mcData), mc)
	if err != nil {
		t.Fatalf("unexpected error while unmarshaling machine-config: %s, err: %v", mcPath, err)
	}
	kc, _, err := getKubeConfigContent(t)
	if err != nil {
		t.Fatal(err)
	}
	appendFileToIgnition(&mc.Spec.Config, defaultMachineKubeConfPath, string(kc))
	anno, err := getNodeAnnotation(mp.Status.Configuration.Name)
	if err != nil {
		t.Fatalf("unexpected error while creating annotations err: %v", err)
	}
	appendFileToIgnition(&mc.Spec.Config, daemonconsts.InitialNodeAnnotationsFilePath, anno)
	res, err := csc.GetConfig(poolRequest{machineConfigPool: testPool})
	if err != nil {
		t.Fatalf("expected err to be nil, received: %v", err)
	}
	validateIgnitionFiles(t, res.Storage.Files, mc.Spec.Config.Storage.Files)
	validateIgnitionSystemd(t, res.Systemd.Units, mc.Spec.Config.Systemd.Units)
}
func getKubeConfigContent(t *testing.T) ([]byte, []byte, error) {
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
	return []byte("dummy-kubeconfig"), []byte("dummy-root-ca"), nil
}
func validateIgnitionFiles(t *testing.T, exp, got []ignv2_2types.File) {
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
	expMap := createFileMap(exp)
	gotMap := createFileMap(got)
	for k, v := range expMap {
		f, ok := gotMap[k]
		if !ok {
			t.Errorf("could not find file: %s", k)
		}
		if !reflect.DeepEqual(v, f) {
			t.Errorf("file validation failed for: %s, exp: %v, got: %v", k, v, f)
		}
	}
}
func validateIgnitionSystemd(t *testing.T, exp, got []ignv2_2types.Unit) {
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
	expMap := createUnitMap(exp)
	gotMap := createUnitMap(got)
	for k, v := range expMap {
		f, ok := gotMap[k]
		if !ok {
			t.Errorf("could not find file: %s", k)
		}
		if !reflect.DeepEqual(v, f) {
			t.Errorf("file validation failed for: %s, exp: %v, got: %v", k, v, f)
		}
	}
}
func createUnitMap(units []ignv2_2types.Unit) map[string]ignv2_2types.Unit {
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
	m := make(map[string]ignv2_2types.Unit)
	for i := range units {
		m[units[i].Name] = units[i]
	}
	return m
}
func createFileMap(files []ignv2_2types.File) map[string]ignv2_2types.File {
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
	m := make(map[string]ignv2_2types.File)
	for i := range files {
		file := path.Join(files[i].Filesystem, files[i].Path)
		m[file] = files[i]
	}
	return m
}
func getTestMachineConfigPool() (*v1.MachineConfigPool, error) {
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
	mpPath := path.Join(testDir, "machine-pools", testPool+".yaml")
	mpData, err := ioutil.ReadFile(mpPath)
	if err != nil {
		return nil, fmt.Errorf("unexpected error while reading machine-pool: %s, err: %v", mpPath, err)
	}
	mp := new(v1.MachineConfigPool)
	err = yaml.Unmarshal(mpData, mp)
	if err != nil {
		return nil, fmt.Errorf("unexpected error while unmarshaling machine-pool: %s, err: %v", mpPath, err)
	}
	return mp, nil
}

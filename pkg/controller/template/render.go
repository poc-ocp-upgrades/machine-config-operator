package template

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"github.com/Masterminds/sprig"
	ctconfig "github.com/coreos/container-linux-config-transpiler/config"
	cttypes "github.com/coreos/container-linux-config-transpiler/config/types"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	"github.com/ghodss/yaml"
	"github.com/golang/glog"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	"github.com/openshift/machine-config-operator/pkg/controller/common"
	"github.com/openshift/machine-config-operator/pkg/version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RenderConfig struct {
	*mcfgv1.ControllerConfigSpec
	PullSecret	string
}

const (
	filesDir		= "files"
	unitsDir		= "units"
	platformAWS		= "aws"
	platformAzure		= "azure"
	platformOpenstack	= "openstack"
	platformLibvirt		= "libvirt"
	platformNone		= "none"
	platformVSphere		= "vsphere"
	platformBase		= "_base"
)

func generateTemplateMachineConfigs(config *RenderConfig, templateDir string) ([]*mcfgv1.MachineConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	infos, err := ioutil.ReadDir(templateDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir %q: %v", templateDir, err)
	}
	cfgs := []*mcfgv1.MachineConfig{}
	for _, info := range infos {
		if !info.IsDir() {
			glog.Infof("ignoring non-directory path %q", info.Name())
			continue
		}
		role := info.Name()
		path := filepath.Join(templateDir, role)
		roleConfigs, err := GenerateMachineConfigsForRole(config, role, path)
		if err != nil {
			return nil, fmt.Errorf("failed to create MachineConfig for role %s: %v", role, err)
		}
		cfgs = append(cfgs, roleConfigs...)
	}
	for _, cfg := range cfgs {
		if cfg.Annotations == nil {
			cfg.Annotations = map[string]string{}
		}
		cfg.Annotations[common.GeneratedByControllerVersionAnnotationKey] = version.Version.String()
	}
	return cfgs, nil
}
func GenerateMachineConfigsForRole(config *RenderConfig, role string, path string) ([]*mcfgv1.MachineConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir %q: %v", path, err)
	}
	cfgs := []*mcfgv1.MachineConfig{}
	for _, info := range infos {
		if !info.IsDir() {
			glog.Infof("ignoring non-directory path %q", info.Name())
			continue
		}
		name := info.Name()
		namePath := filepath.Join(path, name)
		nameConfig, err := generateMachineConfigForName(config, role, name, namePath)
		if err != nil {
			return nil, err
		}
		cfgs = append(cfgs, nameConfig)
	}
	return cfgs, nil
}
func platformFromControllerConfigSpec(ic *mcfgv1.ControllerConfigSpec) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch ic.Platform {
	case "":
		return "", fmt.Errorf("cannot generateMachineConfigs with an empty platform field")
	case platformBase:
		return "", fmt.Errorf("platform _base unsupported")
	case platformAWS, platformAzure, platformOpenstack, platformLibvirt, platformNone:
		return ic.Platform, nil
	default:
		glog.Warningf("Warning: the controller config referenced an unsupported platform: %s", ic.Platform)
		return platformNone, nil
	}
}
func filterTemplates(toFilter map[string]string, path string, config *RenderConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Size() == 0 {
			delete(toFilter, info.Name())
			return nil
		}
		filedata, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %q: %v", path, err)
		}
		renderedData, err := renderTemplate(*config, path, filedata)
		if err != nil {
			return err
		}
		toFilter[info.Name()] = string(renderedData)
		return nil
	}
	return filepath.Walk(path, walkFn)
}
func generateMachineConfigForName(config *RenderConfig, role, name, path string) (*mcfgv1.MachineConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	platform, err := platformFromControllerConfigSpec(config.ControllerConfigSpec)
	if err != nil {
		return nil, err
	}
	platformDirs := []string{}
	for _, dir := range []string{platformBase, platform} {
		platformPath := filepath.Join(path, dir)
		exists, err := existsDir(platformPath)
		if err != nil {
			return nil, err
		}
		if !exists {
			glog.Errorf("could not find expected template directory %s", platformPath)
			return nil, fmt.Errorf("platform %s unsupported", config.Platform)
		}
		platformDirs = append(platformDirs, platformPath)
	}
	files := map[string]string{}
	units := map[string]string{}
	for _, platformDir := range platformDirs {
		p := filepath.Join(platformDir, filesDir)
		exists, err := existsDir(p)
		if err != nil {
			return nil, err
		}
		if exists {
			if err := filterTemplates(files, p, config); err != nil {
				return nil, err
			}
		}
		p = filepath.Join(platformDir, unitsDir)
		exists, err = existsDir(p)
		if err != nil {
			return nil, err
		}
		if exists {
			if err := filterTemplates(units, p, config); err != nil {
				return nil, err
			}
		}
	}
	keySortVals := func(m map[string]string) []string {
		ks := []string{}
		for k := range m {
			ks = append(ks, k)
		}
		sort.Strings(ks)
		vs := []string{}
		for _, k := range ks {
			vs = append(vs, m[k])
		}
		return vs
	}
	ignCfg, err := transpileToIgn(keySortVals(files), keySortVals(units))
	if err != nil {
		return nil, fmt.Errorf("error transpiling ct config to Ignition config: %v", err)
	}
	mcfg := MachineConfigFromIgnConfig(role, name, ignCfg)
	mcfg.Spec.OSImageURL = config.OSImageURL
	return mcfg, nil
}

const (
	machineConfigRoleLabelKey = "machineconfiguration.openshift.io/role"
)

func MachineConfigFromIgnConfig(role string, name string, ignCfg *ignv2_2types.Config) *mcfgv1.MachineConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	labels := map[string]string{machineConfigRoleLabelKey: role}
	return &mcfgv1.MachineConfig{ObjectMeta: metav1.ObjectMeta{Labels: labels, Name: name}, Spec: mcfgv1.MachineConfigSpec{OSImageURL: "", Config: *ignCfg}}
}
func transpileToIgn(files, units []string) (*ignv2_2types.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ctCfg cttypes.Config
	for _, d := range files {
		f := new(cttypes.File)
		if err := yaml.Unmarshal([]byte(d), f); err != nil {
			return nil, fmt.Errorf("failed to unmarshal file into struct: %v", err)
		}
		ctCfg.Storage.Files = append(ctCfg.Storage.Files, *f)
	}
	for _, d := range units {
		u := new(cttypes.SystemdUnit)
		if err := yaml.Unmarshal([]byte(d), u); err != nil {
			return nil, fmt.Errorf("failed to unmarshal systemd unit into struct: %v", err)
		}
		ctCfg.Systemd.Units = append(ctCfg.Systemd.Units, *u)
	}
	ignCfg, rep := ctconfig.Convert(ctCfg, "", nil)
	if rep.IsFatal() {
		return nil, fmt.Errorf("failed to convert config to Ignition config %s", rep)
	}
	return &ignCfg, nil
}
func renderTemplate(config RenderConfig, path string, b []byte) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	funcs := sprig.TxtFuncMap()
	funcs["skip"] = skipMissing
	funcs["etcdServerCertDNSNames"] = etcdServerCertDNSNames
	funcs["etcdPeerCertDNSNames"] = etcdPeerCertDNSNames
	funcs["cloudProvider"] = cloudProvider
	funcs["cloudConfigFlag"] = cloudConfigFlag
	tmpl, err := template.New(path).Funcs(funcs).Parse(string(b))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template %s: %v", path, err)
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, config); err != nil {
		return nil, fmt.Errorf("failed to execute template: %v", err)
	}
	return buf.Bytes(), nil
}

var skipKeyValidate = regexp.MustCompile(`^[_a-z]\w*$`)

func skipMissing(key string) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !skipKeyValidate.Match([]byte(key)) {
		return nil, fmt.Errorf("invalid key for skipKey")
	}
	return fmt.Sprintf("{{.%s}}", key), nil
}
func etcdServerCertDNSNames(cfg RenderConfig) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var dnsNames = []string{"localhost", "etcd.kube-system.svc", "etcd.kube-system.svc.cluster.local", "etcd.openshift-etcd.svc", "etcd.openshift-etcd.svc.cluster.local", "${ETCD_DNS_NAME}"}
	return strings.Join(dnsNames, ","), nil
}
func etcdPeerCertDNSNames(cfg RenderConfig) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if cfg.EtcdDiscoveryDomain == "" {
		return nil, fmt.Errorf("invalid configuration")
	}
	var dnsNames = []string{"${ETCD_DNS_NAME}", cfg.EtcdDiscoveryDomain}
	return strings.Join(dnsNames, ","), nil
}
func cloudProvider(cfg RenderConfig) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch cfg.Platform {
	case platformAWS:
		return platformAWS, nil
	case platformAzure:
		return platformAzure, nil
	case platformOpenstack:
		return platformOpenstack, nil
	case platformVSphere:
		return platformVSphere, nil
	}
	return "", nil
}
func cloudConfigFlag(cfg RenderConfig) interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(cfg.CloudProviderConfig) == 0 {
		return ""
	}
	flag := "--cloud-config=/etc/kubernetes/cloud.conf"
	switch cfg.Platform {
	case platformAzure, platformOpenstack:
		return flag
	default:
		return ""
	}
}
func existsDir(path string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("failed to open dir %q: %v", path, err)
	}
	if !info.IsDir() {
		return false, fmt.Errorf("expected template directory %q is not a directory", path)
	}
	return true, nil
}

package containerruntimeconfig

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"github.com/BurntSushi/toml"
	"github.com/containers/image/docker/reference"
	"github.com/containers/image/pkg/sysregistriesv2"
	storageconfig "github.com/containers/storage/pkg/config"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	crioconfig "github.com/kubernetes-sigs/cri-o/pkg/config"
	apicfgv1 "github.com/openshift/api/config/v1"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	ctrlcommon "github.com/openshift/machine-config-operator/pkg/controller/common"
	"github.com/vincent-petithory/dataurl"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	minLogSize		= 8192
	minPidsLimit		= 20
	crioConfigPath		= "/etc/crio/crio.conf"
	storageConfigPath	= "/etc/containers/storage.conf"
	registriesConfigPath	= "/etc/containers/registries.conf"
)

var errParsingReference error = errors.New("error parsing reference of desired image from cluster version config")

type tomlConfigStorage struct {
	Storage struct {
		Driver		string					`toml:"driver"`
		RunRoot		string					`toml:"runroot"`
		GraphRoot	string					`toml:"graphroot"`
		Options		struct{ storageconfig.OptionsConfig }	`toml:"options"`
	} `toml:"storage"`
}
type tomlConfigCRIO struct {
	Crio struct {
		crioconfig.RootConfig
		API	struct{ crioconfig.APIConfig }		`toml:"api"`
		Runtime	struct{ crioconfig.RuntimeConfig }	`toml:"runtime"`
		Image	struct{ crioconfig.ImageConfig }	`toml:"image"`
		Network	struct{ crioconfig.NetworkConfig }	`toml:"network"`
	} `toml:"crio"`
}
type tomlConfigRegistries struct {
	Registries			[]sysregistriesv2.Registry	`toml:"registry"`
	sysregistriesv2.V1TOMLConfig	`toml:"registries"`
}
type updateConfig func(data []byte, internal *mcfgv1.ContainerRuntimeConfiguration) ([]byte, error)

func createNewCtrRuntimeConfigIgnition(storageTOMLConfig, crioTOMLConfig []byte) ignv2_2types.Config {
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
	tempIgnConfig := ctrlcommon.NewIgnConfig()
	mode := 0644
	if storageTOMLConfig != nil {
		storagedu := dataurl.New(storageTOMLConfig, "text/plain")
		storagedu.Encoding = dataurl.EncodingASCII
		storageTempFile := ignv2_2types.File{Node: ignv2_2types.Node{Filesystem: "root", Path: storageConfigPath}, FileEmbedded1: ignv2_2types.FileEmbedded1{Mode: &mode, Contents: ignv2_2types.FileContents{Source: storagedu.String()}}}
		tempIgnConfig.Storage.Files = append(tempIgnConfig.Storage.Files, storageTempFile)
	}
	if crioTOMLConfig != nil {
		criodu := dataurl.New(crioTOMLConfig, "text/plain")
		criodu.Encoding = dataurl.EncodingASCII
		crioTempFile := ignv2_2types.File{Node: ignv2_2types.Node{Filesystem: "root", Path: crioConfigPath}, FileEmbedded1: ignv2_2types.FileEmbedded1{Mode: &mode, Contents: ignv2_2types.FileContents{Source: criodu.String()}}}
		tempIgnConfig.Storage.Files = append(tempIgnConfig.Storage.Files, crioTempFile)
	}
	return tempIgnConfig
}
func createNewRegistriesConfigIgnition(registriesTOMLConfig []byte) ignv2_2types.Config {
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
	tempIgnConfig := ctrlcommon.NewIgnConfig()
	mode := 0644
	if registriesTOMLConfig != nil {
		regdu := dataurl.New(registriesTOMLConfig, "text/plain")
		regdu.Encoding = dataurl.EncodingASCII
		regTempFile := ignv2_2types.File{Node: ignv2_2types.Node{Filesystem: "root", Path: registriesConfigPath}, FileEmbedded1: ignv2_2types.FileEmbedded1{Mode: &mode, Contents: ignv2_2types.FileContents{Source: regdu.String()}}}
		tempIgnConfig.Storage.Files = append(tempIgnConfig.Storage.Files, regTempFile)
	}
	return tempIgnConfig
}
func findStorageConfig(mc *mcfgv1.MachineConfig) (*ignv2_2types.File, error) {
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
	for _, c := range mc.Spec.Config.Storage.Files {
		if c.Path == storageConfigPath {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("could not find Storage Config")
}
func findCRIOConfig(mc *mcfgv1.MachineConfig) (*ignv2_2types.File, error) {
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
	for _, c := range mc.Spec.Config.Storage.Files {
		if c.Path == crioConfigPath {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("could not find CRI-O Config")
}
func findRegistriesConfig(mc *mcfgv1.MachineConfig) (*ignv2_2types.File, error) {
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
	for _, c := range mc.Spec.Config.Storage.Files {
		if c.Path == registriesConfigPath {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("could not find Registries Config")
}
func getManagedKeyCtrCfg(pool *mcfgv1.MachineConfigPool, config *mcfgv1.ContainerRuntimeConfig) string {
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
	return fmt.Sprintf("99-%s-%s-containerruntime", pool.Name, pool.ObjectMeta.UID)
}
func getManagedKeyReg(pool *mcfgv1.MachineConfigPool, config *apicfgv1.Image) string {
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
	return fmt.Sprintf("99-%s-%s-registries", pool.Name, pool.ObjectMeta.UID)
}
func wrapErrorWithCondition(err error, args ...interface{}) mcfgv1.ContainerRuntimeConfigCondition {
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
	var condition *mcfgv1.ContainerRuntimeConfigCondition
	if err != nil {
		condition = mcfgv1.NewContainerRuntimeConfigCondition(mcfgv1.ContainerRuntimeConfigFailure, v1.ConditionFalse, fmt.Sprintf("Error: %v", err))
	} else {
		condition = mcfgv1.NewContainerRuntimeConfigCondition(mcfgv1.ContainerRuntimeConfigSuccess, v1.ConditionTrue, "Success")
	}
	if len(args) > 0 {
		format, ok := args[0].(string)
		if ok {
			condition.Message = fmt.Sprintf(format, args[:1]...)
		}
	}
	return *condition
}
func updateStorageConfig(data []byte, internal *mcfgv1.ContainerRuntimeConfiguration) ([]byte, error) {
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
	tomlConf := new(tomlConfigStorage)
	if _, err := toml.DecodeReader(bytes.NewBuffer(data), tomlConf); err != nil {
		return nil, fmt.Errorf("error decoding crio config: %v", err)
	}
	if internal.OverlaySize != (resource.Quantity{}) {
		tomlConf.Storage.Options.Size = internal.OverlaySize.String()
	}
	var newData bytes.Buffer
	encoder := toml.NewEncoder(&newData)
	if err := encoder.Encode(*tomlConf); err != nil {
		return nil, err
	}
	return newData.Bytes(), nil
}
func updateCRIOConfig(data []byte, internal *mcfgv1.ContainerRuntimeConfiguration) ([]byte, error) {
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
	tomlConf := new(tomlConfigCRIO)
	if _, err := toml.DecodeReader(bytes.NewBuffer(data), tomlConf); err != nil {
		return nil, fmt.Errorf("error decoding crio config: %v", err)
	}
	if internal.PidsLimit > 0 {
		tomlConf.Crio.Runtime.PidsLimit = internal.PidsLimit
	}
	if internal.LogSizeMax != (resource.Quantity{}) {
		tomlConf.Crio.Runtime.LogSizeMax = internal.LogSizeMax.Value()
	}
	if internal.LogLevel != "" {
		tomlConf.Crio.Runtime.LogLevel = internal.LogLevel
	}
	tomlConf.Crio.StorageOptions = []string{}
	var newData bytes.Buffer
	encoder := toml.NewEncoder(&newData)
	if err := encoder.Encode(*tomlConf); err != nil {
		return nil, err
	}
	return newData.Bytes(), nil
}
func updateRegistriesConfig(data []byte, internalInsecure, internalBlocked []string) ([]byte, error) {
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
	tomlConf := new(tomlConfigRegistries)
	if _, err := toml.Decode(string(data), tomlConf); err != nil {
		return nil, fmt.Errorf("error unmarshalling registries config: %v", err)
	}
	if internalInsecure != nil {
		tomlConf.Insecure = sysregistriesv2.V1TOMLregistries{Registries: internalInsecure}
	}
	if internalBlocked != nil {
		tomlConf.Block = sysregistriesv2.V1TOMLregistries{Registries: internalBlocked}
	}
	var newData bytes.Buffer
	encoder := toml.NewEncoder(&newData)
	if err := encoder.Encode(*tomlConf); err != nil {
		return nil, err
	}
	return newData.Bytes(), nil
}
func validateUserContainerRuntimeConfig(cfg *mcfgv1.ContainerRuntimeConfig) error {
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
	if cfg.Spec.ContainerRuntimeConfig == nil {
		return nil
	}
	ctrcfgValues := reflect.ValueOf(*cfg.Spec.ContainerRuntimeConfig)
	if !ctrcfgValues.IsValid() {
		return fmt.Errorf("containerRuntimeConfig is not valid")
	}
	ctrcfg := cfg.Spec.ContainerRuntimeConfig
	if ctrcfg.PidsLimit > 0 && ctrcfg.PidsLimit < minPidsLimit {
		return fmt.Errorf("invalid PidsLimit %q, cannot be less than 20", ctrcfg.PidsLimit)
	}
	if ctrcfg.LogSizeMax.Value() > 0 && ctrcfg.LogSizeMax.Value() <= minLogSize {
		return fmt.Errorf("invalid LogSizeMax %q, cannot be less than 8kB", ctrcfg.LogSizeMax.String())
	}
	if ctrcfg.LogLevel != "" {
		validLogLevels := map[string]bool{"error": true, "fatal": true, "panic": true, "warn": true, "info": true, "debug": true}
		if !validLogLevels[ctrcfg.LogLevel] {
			return fmt.Errorf("invalid LogLevel %q, must be one of error, fatal, panic, warn, info, or debug", ctrcfg.LogLevel)
		}
	}
	return nil
}
func getValidRegistries(clusterVersionStatus *apicfgv1.ClusterVersionStatus, imgSpec *apicfgv1.ImageSpec) ([]string, []string, error) {
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
	if clusterVersionStatus == nil || imgSpec == nil {
		return nil, nil, nil
	}
	var blockedRegs []string
	insecureRegs := imgSpec.RegistrySources.InsecureRegistries
	ref, err := reference.ParseNamed(clusterVersionStatus.Desired.Image)
	if err != nil {
		return nil, nil, errParsingReference
	}
	payloadReg := reference.Domain(ref)
	for i, reg := range imgSpec.RegistrySources.BlockedRegistries {
		if reg == payloadReg {
			blockedRegs = append(blockedRegs, imgSpec.RegistrySources.BlockedRegistries[i+1:]...)
			return insecureRegs, blockedRegs, fmt.Errorf("error adding %q to blocked registries, cannot block the registry being used by the payload", payloadReg)
		}
		blockedRegs = append(blockedRegs, reg)
	}
	return insecureRegs, blockedRegs, nil
}

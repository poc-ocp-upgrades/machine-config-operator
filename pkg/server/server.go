package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	daemonconsts "github.com/openshift/machine-config-operator/pkg/daemon/constants"
	"github.com/vincent-petithory/dataurl"
)

const (
	defaultMachineKubeConfPath	= "/etc/kubernetes/kubeconfig"
	pivotRebootNeeded		= "/run/pivot/reboot-needed"
	defaultFileSystem		= "root"
)

type kubeconfigFunc func() (kubeconfigData []byte, rootCAData []byte, err error)
type appenderFunc func(*ignv2_2types.Config) error
type Server interface {
	GetConfig(poolRequest) (*ignv2_2types.Config, error)
}

func getAppenders(cr poolRequest, currMachineConfig string, f kubeconfigFunc, osimageurl string) []appenderFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	appenders := []appenderFunc{func(config *ignv2_2types.Config) error {
		return appendNodeAnnotations(config, currMachineConfig)
	}, func(config *ignv2_2types.Config) error {
		return appendInitialPivot(config, osimageurl)
	}, func(config *ignv2_2types.Config) error {
		return appendKubeConfig(config, f)
	}}
	return appenders
}
func boolToPtr(b bool) *bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &b
}
func appendInitialPivot(conf *ignv2_2types.Config, osimageurl string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if osimageurl == "" {
		return nil
	}
	appendFileToIgnition(conf, daemonconsts.EtcPivotFile, osimageurl+"\n")
	if len(conf.Systemd.Units) == 0 {
		conf.Systemd.Units = make([]ignv2_2types.Unit, 0)
	}
	unit := ignv2_2types.Unit{Name: "mcd-write-pivot-reboot.service", Enabled: boolToPtr(true), Contents: `[Unit]
Before=pivot.service
ConditionFirstBoot=true
[Service]
ExecStart=/bin/sh -c 'mkdir /run/pivot && touch /run/pivot/reboot-needed'
[Install]
WantedBy=multi-user.target
`}
	conf.Systemd.Units = append(conf.Systemd.Units, unit)
	return nil
}
func appendKubeConfig(conf *ignv2_2types.Config, f kubeconfigFunc) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kcData, _, err := f()
	if err != nil {
		return err
	}
	appendFileToIgnition(conf, defaultMachineKubeConfPath, string(kcData))
	return nil
}
func appendNodeAnnotations(conf *ignv2_2types.Config, currConf string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	anno, err := getNodeAnnotation(currConf)
	if err != nil {
		return err
	}
	appendFileToIgnition(conf, daemonconsts.InitialNodeAnnotationsFilePath, string(anno))
	return nil
}
func getNodeAnnotation(conf string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nodeAnnotations := map[string]string{daemonconsts.CurrentMachineConfigAnnotationKey: conf, daemonconsts.DesiredMachineConfigAnnotationKey: conf, daemonconsts.MachineConfigDaemonStateAnnotationKey: daemonconsts.MachineConfigDaemonStateDone}
	contents, err := json.Marshal(nodeAnnotations)
	if err != nil {
		return "", fmt.Errorf("could not marshal node annotations, err: %v", err)
	}
	return string(contents), nil
}
func copyFileToIgnition(conf *ignv2_2types.Config, outPath, srcPath string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	contents, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("could not read file from: %s, err: %v", srcPath, err)
	}
	appendFileToIgnition(conf, outPath, string(contents))
	return nil
}
func appendFileToIgnition(conf *ignv2_2types.Config, outPath, contents string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileMode := int(420)
	file := ignv2_2types.File{Node: ignv2_2types.Node{Filesystem: defaultFileSystem, Path: outPath}, FileEmbedded1: ignv2_2types.FileEmbedded1{Contents: ignv2_2types.FileContents{Source: getEncodedContent(contents)}, Mode: &fileMode}}
	if len(conf.Storage.Files) == 0 {
		conf.Storage.Files = make([]ignv2_2types.File, 0)
	}
	conf.Storage.Files = append(conf.Storage.Files, file)
}
func getDecodedContent(inp string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d, err := dataurl.DecodeString(inp)
	if err != nil {
		return "", err
	}
	return string(d.Data), nil
}
func getEncodedContent(inp string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return (&url.URL{Scheme: "data", Opaque: "," + dataurl.Escape([]byte(inp))}).String()
}

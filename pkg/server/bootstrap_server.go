package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	yaml "github.com/ghodss/yaml"
	"github.com/golang/glog"
	clientcmd "k8s.io/client-go/tools/clientcmd/api/v1"
	"github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
)

var _ = Server(&bootstrapServer{})

type bootstrapServer struct {
	serverBaseDir	string
	kubeconfigFunc	kubeconfigFunc
}

func NewBootstrapServer(dir, kubeconfig string) (Server, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, err := os.Stat(kubeconfig); err != nil {
		return nil, fmt.Errorf("kubeconfig not found at location: %s", kubeconfig)
	}
	return &bootstrapServer{serverBaseDir: dir, kubeconfigFunc: func() ([]byte, []byte, error) {
		return kubeconfigFromFile(kubeconfig)
	}}, nil
}
func (bsc *bootstrapServer) GetConfig(cr poolRequest) (*ignv2_2types.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fileName := path.Join(bsc.serverBaseDir, "machine-pools", cr.machineConfigPool+".yaml")
	glog.Infof("reading file %q", fileName)
	data, err := ioutil.ReadFile(fileName)
	if os.IsNotExist(err) {
		glog.Errorf("could not find file: %s", fileName)
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("server: could not read file %s, err: %v", fileName, err)
	}
	mp := new(v1.MachineConfigPool)
	err = yaml.Unmarshal(data, mp)
	if err != nil {
		return nil, fmt.Errorf("server: could not unmarshal file %s, err: %v", fileName, err)
	}
	currConf := mp.Status.Configuration.Name
	fileName = path.Join(bsc.serverBaseDir, "machine-configs", currConf+".yaml")
	glog.Infof("reading file %q", fileName)
	data, err = ioutil.ReadFile(fileName)
	if os.IsNotExist(err) {
		glog.Errorf("could not find file: %s", fileName)
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("server: could not read file %s, err: %v", fileName, err)
	}
	mc := new(v1.MachineConfig)
	err = yaml.Unmarshal(data, mc)
	if err != nil {
		return nil, fmt.Errorf("server: could not unmarshal file %s, err: %v", fileName, err)
	}
	appenders := getAppenders(cr, currConf, bsc.kubeconfigFunc, mc.Spec.OSImageURL)
	for _, a := range appenders {
		if err := a(&mc.Spec.Config); err != nil {
			return nil, err
		}
	}
	return &mc.Spec.Config, nil
}
func kubeconfigFromFile(path string) ([]byte, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kcData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting kubeconfig from disk: %v", err)
	}
	kc := clientcmd.Config{}
	if err := yaml.Unmarshal(kcData, &kc); err != nil {
		return nil, nil, err
	}
	return kcData, kc.Clusters[0].Cluster.CertificateAuthorityData, nil
}

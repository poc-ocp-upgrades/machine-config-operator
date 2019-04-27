package server

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
	yaml "github.com/ghodss/yaml"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	clientcmdv1 "k8s.io/client-go/tools/clientcmd/api/v1"
	"github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/typed/machineconfiguration.openshift.io/v1"
)

const (
	inClusterConfig		= ""
	bootstrapTokenDir	= "/etc/mcs/bootstrap-token"
)

var _ = Server(&clusterServer{})

type clusterServer struct {
	machineClient	v1.MachineconfigurationV1Interface
	kubeconfigFunc	kubeconfigFunc
}

func NewClusterServer(kubeConfig, apiserverURL string) (Server, error) {
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
	restConfig, err := getClientConfig(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Kubernetes rest client: %v", err)
	}
	mc := v1.NewForConfigOrDie(restConfig)
	return &clusterServer{machineClient: mc, kubeconfigFunc: func() ([]byte, []byte, error) {
		return kubeconfigFromSecret(bootstrapTokenDir, apiserverURL)
	}}, nil
}
func (cs *clusterServer) GetConfig(cr poolRequest) (*ignv2_2types.Config, error) {
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
	mp, err := cs.machineClient.MachineConfigPools().Get(cr.machineConfigPool, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not fetch pool. err: %v", err)
	}
	currConf := mp.Status.Configuration.Name
	mc, err := cs.machineClient.MachineConfigs().Get(currConf, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not fetch config %s, err: %v", currConf, err)
	}
	appenders := getAppenders(cr, currConf, cs.kubeconfigFunc, mc.Spec.OSImageURL)
	for _, a := range appenders {
		if err := a(&mc.Spec.Config); err != nil {
			return nil, err
		}
	}
	return &mc.Spec.Config, nil
}
func getClientConfig(path string) (*rest.Config, error) {
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
	if path != inClusterConfig {
		return clientcmd.BuildConfigFromFlags("", path)
	}
	return rest.InClusterConfig()
}
func kubeconfigFromSecret(secertDir string, apiserverURL string) ([]byte, []byte, error) {
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
	caFile := filepath.Join(secertDir, corev1.ServiceAccountRootCAKey)
	tokenFile := filepath.Join(secertDir, corev1.ServiceAccountTokenKey)
	caData, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to read %s: %v", caFile, err)
	}
	token, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to read %s: %v", tokenFile, err)
	}
	kubeconfig := clientcmdv1.Config{Clusters: []clientcmdv1.NamedCluster{{Name: "local", Cluster: clientcmdv1.Cluster{Server: apiserverURL, CertificateAuthorityData: caData}}}, AuthInfos: []clientcmdv1.NamedAuthInfo{{Name: "kubelet", AuthInfo: clientcmdv1.AuthInfo{Token: string(token)}}}, Contexts: []clientcmdv1.NamedContext{{Name: "kubelet", Context: clientcmdv1.Context{Cluster: "local", AuthInfo: "kubelet"}}}, CurrentContext: "kubelet"}
	kcData, err := yaml.Marshal(kubeconfig)
	if err != nil {
		return nil, nil, err
	}
	return kcData, caData, nil
}

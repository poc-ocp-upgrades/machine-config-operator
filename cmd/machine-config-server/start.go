package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/openshift/machine-config-operator/pkg/server"
	"github.com/openshift/machine-config-operator/pkg/version"
	"github.com/spf13/cobra"
)

var (
	startCmd	= &cobra.Command{Use: "start", Short: "Starts Machine Config Server", Long: "", Run: runStartCmd}
	startOpts	struct {
		kubeconfig		string
		apiserverURL	string
	}
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringVar(&startOpts.kubeconfig, "kubeconfig", "", "Kubeconfig file to access a remote cluster (testing only)")
	startCmd.PersistentFlags().StringVar(&startOpts.apiserverURL, "apiserver-url", "", "URL for apiserver; Used to generate kubeconfig")
}
func runStartCmd(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flag.Set("logtostderr", "true")
	flag.Parse()
	glog.Infof("Version: %+v", version.Version)
	if startOpts.apiserverURL == "" {
		glog.Exitf("--apiserver-url cannot be empty")
	}
	cs, err := server.NewClusterServer(startOpts.kubeconfig, startOpts.apiserverURL)
	if err != nil {
		glog.Exitf("Machine Config Server exited with error: %v", err)
	}
	apiHandler := server.NewServerAPIHandler(cs)
	secureServer := server.NewAPIServer(apiHandler, rootOpts.sport, false, rootOpts.cert, rootOpts.key)
	insecureServer := server.NewAPIServer(apiHandler, rootOpts.isport, true, "", "")
	stopCh := make(chan struct{})
	go secureServer.Serve()
	go insecureServer.Serve()
	<-stopCh
	panic("not possible")
}

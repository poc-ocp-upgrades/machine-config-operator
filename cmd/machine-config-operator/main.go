package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

const (
	componentName		= "machine-config"
	componentNamespace	= "openshift-machine-config-operator"
)

var (
	rootCmd		= &cobra.Command{Use: componentName, Short: "Run Machine Config Operator", Long: ""}
	rootOpts	struct{}
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
		glog.Exitf("Error executing mcc: %v", err)
	}
}

package main

import (
	"flag"
	"fmt"
	"github.com/openshift/machine-config-operator/pkg/version"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{Use: "version", Short: "Print the version number of Machine Config Controller", Long: `All software has versions. This is Machine Config Controller's.`, Run: runVersionCmd}
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rootCmd.AddCommand(versionCmd)
}
func runVersionCmd(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flag.Set("logtostderr", "true")
	flag.Parse()
	program := "MachineConfigController"
	version := "v" + version.Version.String()
	fmt.Println(program, version)
}

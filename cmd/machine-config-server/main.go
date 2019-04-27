package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

const (
	componentName = "machine-config-server"
)

var (
	rootCmd		= &cobra.Command{Use: componentName, Short: "Run Machine Config Server", Long: ""}
	rootOpts	struct {
		sport	int
		isport	int
		cert	string
		key	string
	}
)

func init() {
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
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	rootCmd.PersistentFlags().IntVar(&rootOpts.sport, "secure-port", 22623, "secure port to serve ignition configs")
	rootCmd.PersistentFlags().StringVar(&rootOpts.cert, "cert", "/etc/ssl/mcs/tls.crt", "cert file for TLS")
	rootCmd.PersistentFlags().StringVar(&rootOpts.key, "key", "/etc/ssl/mcs/tls.key", "key file for TLS")
	rootCmd.PersistentFlags().IntVar(&rootOpts.isport, "insecure-port", 22624, "insecure port to serve ignition configs")
}
func main() {
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
	if err := rootCmd.Execute(); err != nil {
		glog.Exitf("Error executing mcs: %v", err)
	}
}

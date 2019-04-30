package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"syscall"
	"github.com/golang/glog"
	"github.com/openshift/machine-config-operator/internal/clients"
	controllercommon "github.com/openshift/machine-config-operator/pkg/controller/common"
	"github.com/openshift/machine-config-operator/pkg/daemon"
	mcfgclientset "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned"
	"github.com/openshift/machine-config-operator/pkg/version"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
)

var (
	startCmd	= &cobra.Command{Use: "start", Short: "Starts Machine Config Daemon", Long: "", Run: runStartCmd}
	startOpts	struct {
		kubeconfig		string
		nodeName		string
		rootMount		string
		onceFrom		string
		skipReboot		bool
		fromIgnition		bool
		kubeletHealthzEnabled	bool
		kubeletHealthzEndpoint	string
	}
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringVar(&startOpts.kubeconfig, "kubeconfig", "", "Kubeconfig file to access a remote cluster (testing only)")
	startCmd.PersistentFlags().StringVar(&startOpts.nodeName, "node-name", "", "kubernetes node name daemon is managing.")
	startCmd.PersistentFlags().StringVar(&startOpts.rootMount, "root-mount", "/rootfs", "where the nodes root filesystem is mounted for chroot and file manipulation.")
	startCmd.PersistentFlags().StringVar(&startOpts.onceFrom, "once-from", "", "Runs the daemon once using a provided file path or URL endpoint as its machine config or ignition (.ign) file source")
	startCmd.PersistentFlags().BoolVar(&startOpts.skipReboot, "skip-reboot", false, "Skips reboot after a sync, applies only in once-from")
	startCmd.PersistentFlags().BoolVar(&startOpts.kubeletHealthzEnabled, "kubelet-healthz-enabled", true, "kubelet healthz endpoint monitoring")
	startCmd.PersistentFlags().StringVar(&startOpts.kubeletHealthzEndpoint, "kubelet-healthz-endpoint", "http://localhost:10248/healthz", "healthz endpoint to check health")
}
func getBootID() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	currentBootIDBytes, err := ioutil.ReadFile("/proc/sys/kernel/random/boot_id")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(currentBootIDBytes)), nil
}
func runStartCmd(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flag.Set("logtostderr", "true")
	flag.Parse()
	glog.V(2).Infof("Options parsed: %+v", startOpts)
	glog.Infof("Version: %+v", version.Version)
	operatingSystem, err := daemon.GetHostRunningOS(startOpts.rootMount)
	if err != nil {
		glog.Fatalf("Error found when checking operating system: %s", err)
	}
	if startOpts.nodeName == "" {
		name, ok := os.LookupEnv("NODE_NAME")
		if !ok || name == "" {
			glog.Fatalf("node-name is required")
		}
		startOpts.nodeName = name
	}
	if _, err := os.Stat(startOpts.rootMount); err != nil {
		if os.IsNotExist(err) {
			glog.Fatalf("rootMount %s does not exist", startOpts.rootMount)
		}
		glog.Fatalf("Unable to verify rootMount %s exists: %s", startOpts.rootMount, err)
	}
	stopCh := make(chan struct{})
	defer close(stopCh)
	exitCh := make(chan error)
	defer close(exitCh)
	glog.Info("Starting node writer")
	nodeWriter := daemon.NewNodeWriter()
	go nodeWriter.Run(stopCh)
	cb, err := clients.NewBuilder(startOpts.kubeconfig)
	if err != nil {
		if startOpts.onceFrom != "" {
			glog.Info("Cannot initialize ClientBuilder, likely in onceFrom mode with Ignition")
		} else {
			glog.Fatalf("Failed to initialize ClientBuilder: %v", err)
		}
	}
	var kubeClient kubernetes.Interface
	if cb != nil {
		kubeClient, err = cb.KubeClient(componentName)
		if err != nil {
			glog.Info("Cannot initialize kubeClient, likely in onceFrom mode with Ignition")
		}
	}
	var dn *daemon.Daemon
	bootID, err := getBootID()
	if err != nil {
		glog.Fatalf("Cannot get boot ID: %v", err)
	}
	if startOpts.onceFrom != "" {
		var mcClient mcfgclientset.Interface
		if cb != nil {
			mcClient, err = cb.MachineConfigClient(componentName)
			if err != nil {
				glog.Info("Cannot initialize MC client, likely in onceFrom mode with Ignition")
			}
		}
		dn, err = daemon.New(startOpts.rootMount, startOpts.nodeName, operatingSystem, daemon.NewNodeUpdaterClient(), bootID, startOpts.onceFrom, startOpts.skipReboot, mcClient, kubeClient, startOpts.kubeletHealthzEnabled, startOpts.kubeletHealthzEndpoint, nodeWriter, exitCh, stopCh)
		if err != nil {
			glog.Fatalf("Failed to initialize single run daemon: %v", err)
		}
	} else {
		if kubeClient == nil {
			panic("Running in cluster mode without a kubeClient")
		}
		ctx := controllercommon.CreateControllerContext(cb, stopCh, componentName)
		dn, err = daemon.NewClusterDrivenDaemon(startOpts.rootMount, startOpts.nodeName, operatingSystem, daemon.NewNodeUpdaterClient(), ctx.InformerFactory.Machineconfiguration().V1().MachineConfigs(), kubeClient, bootID, startOpts.onceFrom, startOpts.skipReboot, ctx.KubeInformerFactory.Core().V1().Nodes(), startOpts.kubeletHealthzEnabled, startOpts.kubeletHealthzEndpoint, nodeWriter, exitCh, stopCh)
		if err != nil {
			glog.Fatalf("Failed to initialize daemon: %v", err)
		}
		if err := dn.BindPodMounts(); err != nil {
			glog.Fatalf("Binding pod mounts: %s", err)
		}
		ctx.KubeInformerFactory.Start(stopCh)
		ctx.InformerFactory.Start(stopCh)
		close(ctx.InformersStarted)
	}
	glog.Infof(`Calling chroot("%s")`, startOpts.rootMount)
	if err := syscall.Chroot(startOpts.rootMount); err != nil {
		glog.Fatalf("Unable to chroot to %s: %s", startOpts.rootMount, err)
	}
	glog.V(2).Infof("Moving to / inside the chroot")
	if err := os.Chdir("/"); err != nil {
		glog.Fatalf("Unable to change directory to /: %s", err)
	}
	glog.Info("Starting MachineConfigDaemon")
	defer glog.Info("Shutting down MachineConfigDaemon")
	if err := dn.Run(stopCh, exitCh); err != nil {
		glog.Fatalf("Failed to run: %v", err)
	}
}

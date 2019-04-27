package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/openshift/machine-config-operator/cmd/common"
	"github.com/openshift/machine-config-operator/internal/clients"
	controllercommon "github.com/openshift/machine-config-operator/pkg/controller/common"
	containerruntimeconfig "github.com/openshift/machine-config-operator/pkg/controller/container-runtime-config"
	kubeletconfig "github.com/openshift/machine-config-operator/pkg/controller/kubelet-config"
	"github.com/openshift/machine-config-operator/pkg/controller/node"
	"github.com/openshift/machine-config-operator/pkg/controller/render"
	"github.com/openshift/machine-config-operator/pkg/controller/template"
	"github.com/openshift/machine-config-operator/pkg/version"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/leaderelection"
)

var (
	startCmd	= &cobra.Command{Use: "start", Short: "Starts Machine Config Controller", Long: "", Run: runStartCmd}
	startOpts	struct {
		kubeconfig		string
		templates		string
		resourceLockNamespace	string
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
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringVar(&startOpts.kubeconfig, "kubeconfig", "", "Kubeconfig file to access a remote cluster (testing only)")
	startCmd.PersistentFlags().StringVar(&startOpts.resourceLockNamespace, "resourcelock-namespace", metav1.NamespaceSystem, "Path to the template files used for creating MachineConfig objects")
}
func runStartCmd(cmd *cobra.Command, args []string) {
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
	flag.Set("logtostderr", "true")
	flag.Parse()
	glog.Infof("Version: %+v", version.Version)
	cb, err := clients.NewBuilder(startOpts.kubeconfig)
	if err != nil {
		glog.Fatalf("error creating clients: %v", err)
	}
	run := func(ctx context.Context) {
		ctrlctx := controllercommon.CreateControllerContext(cb, ctx.Done(), componentName)
		controllers := createControllers(ctrlctx)
		ctrlctx.InformerFactory.Start(ctrlctx.Stop)
		ctrlctx.KubeInformerFactory.Start(ctrlctx.Stop)
		ctrlctx.OpenShiftConfigKubeNamespacedInformerFactory.Start(ctrlctx.Stop)
		ctrlctx.ConfigInformerFactory.Start(ctrlctx.Stop)
		close(ctrlctx.InformersStarted)
		for _, c := range controllers {
			go c.Run(2, ctrlctx.Stop)
		}
		select {}
	}
	leaderelection.RunOrDie(context.TODO(), leaderelection.LeaderElectionConfig{Lock: common.CreateResourceLock(cb, startOpts.resourceLockNamespace, componentName), LeaseDuration: common.LeaseDuration, RenewDeadline: common.RenewDeadline, RetryPeriod: common.RetryPeriod, Callbacks: leaderelection.LeaderCallbacks{OnStartedLeading: run, OnStoppedLeading: func() {
		glog.Fatalf("leaderelection lost")
	}}})
	panic("unreachable")
}
func createControllers(ctx *controllercommon.ControllerContext) []controllercommon.Controller {
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
	var controllers []controllercommon.Controller
	controllers = append(controllers, template.New(rootOpts.templates, ctx.InformerFactory.Machineconfiguration().V1().ControllerConfigs(), ctx.InformerFactory.Machineconfiguration().V1().MachineConfigs(), ctx.OpenShiftConfigKubeNamespacedInformerFactory.Core().V1().Secrets(), ctx.ClientBuilder.KubeClientOrDie("template-controller"), ctx.ClientBuilder.MachineConfigClientOrDie("template-controller")), kubeletconfig.New(rootOpts.templates, ctx.InformerFactory.Machineconfiguration().V1().MachineConfigPools(), ctx.InformerFactory.Machineconfiguration().V1().ControllerConfigs(), ctx.InformerFactory.Machineconfiguration().V1().KubeletConfigs(), ctx.ConfigInformerFactory.Config().V1().FeatureGates(), ctx.ClientBuilder.KubeClientOrDie("kubelet-config-controller"), ctx.ClientBuilder.MachineConfigClientOrDie("kubelet-config-controller")), containerruntimeconfig.New(rootOpts.templates, ctx.InformerFactory.Machineconfiguration().V1().MachineConfigPools(), ctx.InformerFactory.Machineconfiguration().V1().ControllerConfigs(), ctx.InformerFactory.Machineconfiguration().V1().ContainerRuntimeConfigs(), ctx.ConfigInformerFactory.Config().V1().Images(), ctx.ConfigInformerFactory.Config().V1().ClusterVersions(), ctx.ClientBuilder.KubeClientOrDie("container-runtime-config-controller"), ctx.ClientBuilder.MachineConfigClientOrDie("container-runtime-config-controller"), ctx.ClientBuilder.ConfigClientOrDie("container-runtime-config-controller")), render.New(ctx.InformerFactory.Machineconfiguration().V1().MachineConfigPools(), ctx.InformerFactory.Machineconfiguration().V1().MachineConfigs(), ctx.InformerFactory.Machineconfiguration().V1().ControllerConfigs(), ctx.ClientBuilder.KubeClientOrDie("render-controller"), ctx.ClientBuilder.MachineConfigClientOrDie("render-controller")), node.New(ctx.InformerFactory.Machineconfiguration().V1().MachineConfigPools(), ctx.KubeInformerFactory.Core().V1().Nodes(), ctx.ClientBuilder.KubeClientOrDie("node-update-controller"), ctx.ClientBuilder.MachineConfigClientOrDie("node-update-controller")))
	return controllers
}

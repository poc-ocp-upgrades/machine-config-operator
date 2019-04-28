package e2e_test

import (
	"testing"
	"github.com/openshift/machine-config-operator/test/e2e/framework"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestOSImageURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cs := framework.NewClientSet("")
	mcp, err := cs.MachineConfigPools().Get("worker", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("%#v", err)
	}
	mc, err := cs.MachineConfigs().Get(mcp.Status.Configuration.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("%#v", err)
	}
	if mc.Spec.OSImageURL == "" {
		t.Fatalf("Empty OSImageURL for %s", mc.Name)
	}
	mcp, err = cs.MachineConfigPools().Get("master", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("%#v", err)
	}
	mc, err = cs.MachineConfigs().Get(mcp.Status.Configuration.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("%#v", err)
	}
	if mc.Spec.OSImageURL == "" {
		t.Fatalf("Empty OSImageURL for %s", mc.Name)
	}
}

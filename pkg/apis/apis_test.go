package apis

import (
	"testing"
)

func TestGroupName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if got, want := GroupName, "machineconfiguration.openshift.io"; got != want {
		t.Fatalf("mismatch group name, got: %s want: %s", got, want)
	}
}

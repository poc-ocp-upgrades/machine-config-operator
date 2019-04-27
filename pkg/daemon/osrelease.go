package daemon

import (
	"fmt"
	"path"
	"github.com/ashcrow/osrelease"
)

const (
	machineConfigDaemonOSRHCOS	= "RHCOS"
	machineConfigDaemonOSRHEL	= "RHEL"
	machineConfigDaemonOSCENTOS	= "CENTOS"
)

func GetHostRunningOS(rootFs string) (string, error) {
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
	libPath := path.Join(rootFs, "usr", "lib", "os-release")
	etcPath := path.Join(rootFs, "etc", "os-release")
	or, err := osrelease.NewWithOverrides(etcPath, libPath)
	if err != nil {
		return "", err
	}
	switch or.ID {
	case "rhcos":
		return machineConfigDaemonOSRHCOS, nil
	case "rhel":
		return machineConfigDaemonOSRHEL, nil
	case "centos":
		return machineConfigDaemonOSCENTOS, nil
	default:
		return "", fmt.Errorf("an unsupported OS is being used: %s:%s", or.ID, or.VARIANT_ID)
	}
}

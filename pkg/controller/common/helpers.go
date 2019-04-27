package common

import (
	ignv2_2types "github.com/coreos/ignition/config/v2_2/types"
)

func NewIgnConfig() ignv2_2types.Config {
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
	return ignv2_2types.Config{Ignition: ignv2_2types.Ignition{Version: "2.2.0"}}
}

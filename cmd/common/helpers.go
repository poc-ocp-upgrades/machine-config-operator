package common

import (
	"os"
	"time"
	"github.com/openshift/machine-config-operator/internal/clients"
	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/record"
)

const (
	LeaseDuration	= 90 * time.Second
	RenewDeadline	= 60 * time.Second
	RetryPeriod	= 30 * time.Second
)

func CreateResourceLock(cb *clients.Builder, componentNamespace, componentName string) resourcelock.Interface {
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
	recorder := record.NewBroadcaster().NewRecorder(runtime.NewScheme(), v1.EventSource{Component: componentName})
	id, err := os.Hostname()
	if err != nil {
		glog.Fatalf("error creating lock: %v", err)
	}
	id = id + "_" + string(uuid.NewUUID())
	return &resourcelock.ConfigMapLock{ConfigMapMeta: metav1.ObjectMeta{Namespace: componentNamespace, Name: componentName}, Client: cb.KubeClientOrDie("leader-election").CoreV1(), LockConfig: resourcelock.ResourceLockConfig{Identity: id, EventRecorder: recorder}}
}

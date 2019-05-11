package daemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/openshift/machine-config-operator/pkg/daemon/constants"
)

const (
	pivotUnit		= "pivot.service"
	rpmostreedUnit	= "rpm-ostreed.service"
)

type RpmOstreeState struct{ Deployments []RpmOstreeDeployment }
type RpmOstreeDeployment struct {
	ID				string		`json:"id"`
	OSName			string		`json:"osname"`
	Serial			int32		`json:"serial"`
	Checksum		string		`json:"checksum"`
	Version			string		`json:"version"`
	Timestamp		uint64		`json:"timestamp"`
	Booted			bool		`json:"booted"`
	Origin			string		`json:"origin"`
	CustomOrigin	[]string	`json:"custom-origin"`
}
type NodeUpdaterClient interface {
	GetStatus() (string, error)
	GetBootedOSImageURL(string) (string, string, error)
	RunPivot(string) error
}
type RpmOstreeClient struct{}

func NewNodeUpdaterClient() NodeUpdaterClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &RpmOstreeClient{}
}
func (r *RpmOstreeClient) getBootedDeployment(rootMount string) (*RpmOstreeDeployment, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var rosState RpmOstreeState
	output, err := RunGetOut("chroot", rootMount, "rpm-ostree", "status", "--json")
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(output, &rosState); err != nil {
		return nil, fmt.Errorf("failed to parse `rpm-ostree status --json` output: %v", err)
	}
	for _, deployment := range rosState.Deployments {
		if deployment.Booted {
			return &deployment, nil
		}
	}
	return nil, fmt.Errorf("not currently booted in a deployment")
}
func (r *RpmOstreeClient) GetStatus() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	output, err := RunGetOut("rpm-ostree", "status")
	if err != nil {
		return "", err
	}
	return string(output), nil
}
func (r *RpmOstreeClient) GetBootedOSImageURL(rootMount string) (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bootedDeployment, err := r.getBootedDeployment(rootMount)
	if err != nil {
		return "", "", err
	}
	osImageURL := "<not pivoted>"
	if len(bootedDeployment.CustomOrigin) > 0 {
		if strings.HasPrefix(bootedDeployment.CustomOrigin[0], "pivot://") {
			osImageURL = bootedDeployment.CustomOrigin[0][len("pivot://"):]
		}
	}
	return osImageURL, bootedDeployment.Version, nil
}
func (r *RpmOstreeClient) RunPivot(osImageURL string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := os.MkdirAll(filepath.Dir(constants.EtcPivotFile), os.FileMode(0755)); err != nil {
		return fmt.Errorf("error creating leading dirs for %s: %v", constants.EtcPivotFile, err)
	}
	if err := ioutil.WriteFile(constants.EtcPivotFile, []byte(osImageURL), 0644); err != nil {
		return fmt.Errorf("error writing to %s: %v", constants.EtcPivotFile, err)
	}
	journalStopCh := make(chan time.Time)
	defer close(journalStopCh)
	go followPivotJournalLogs(journalStopCh)
	err := exec.Command("systemctl", "start", "pivot.service").Run()
	if err != nil {
		return errors.Wrapf(err, "failed to start pivot.service")
	}
	return nil
}
func followPivotJournalLogs(stopCh <-chan time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := exec.Command("journalctl", "-f", "-b", "-o", "cat", "-u", "rpm-ostreed", "-u", "pivot")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		glog.Fatal(err)
	}
	go func() {
		<-stopCh
		cmd.Process.Kill()
	}()
}

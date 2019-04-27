package template

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/client-go/util/retry"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	mcfgclientv1 "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/typed/machineconfiguration.openshift.io/v1"
	"github.com/openshift/machine-config-operator/pkg/version"
)

func (ctrl *Controller) syncRunningStatus(ctrlconfig *mcfgv1.ControllerConfig) error {
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
	updateFunc := func(cfg *mcfgv1.ControllerConfig) error {
		reason := fmt.Sprintf("syncing towards (%d) generation using controller version %s", cfg.GetGeneration(), version.Version)
		rcond := mcfgv1.NewControllerConfigStatusCondition(mcfgv1.TemplateContollerRunning, corev1.ConditionTrue, reason, "")
		mcfgv1.SetControllerConfigStatusCondition(&cfg.Status, *rcond)
		if cfg.GetGeneration() != cfg.Status.ObservedGeneration && mcfgv1.IsControllerConfigStatusConditionPresentAndEqual(cfg.Status.Conditions, mcfgv1.TemplateContollerCompleted, corev1.ConditionTrue) {
			acond := mcfgv1.NewControllerConfigStatusCondition(mcfgv1.TemplateContollerCompleted, corev1.ConditionFalse, fmt.Sprintf("%s due to change in Generation", reason), "")
			mcfgv1.SetControllerConfigStatusCondition(&cfg.Status, *acond)
		}
		cfg.Status.ObservedGeneration = ctrlconfig.GetGeneration()
		return nil
	}
	return updateControllerConfigStatus(ctrlconfig.GetName(), ctrl.ccLister.Get, ctrl.client.MachineconfigurationV1().ControllerConfigs(), updateFunc)
}
func (ctrl *Controller) syncFailingStatus(ctrlconfig *mcfgv1.ControllerConfig, oerr error) error {
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
	if oerr == nil {
		return oerr
	}
	updateFunc := func(cfg *mcfgv1.ControllerConfig) error {
		reason := oerr.Error()
		message := fmt.Sprintf("failed to syncing towards (%d) generation using controller version %s: %v", cfg.GetGeneration(), version.Version, oerr)
		fcond := mcfgv1.NewControllerConfigStatusCondition(mcfgv1.TemplateContollerFailing, corev1.ConditionTrue, reason, message)
		mcfgv1.SetControllerConfigStatusCondition(&cfg.Status, *fcond)
		acond := mcfgv1.NewControllerConfigStatusCondition(mcfgv1.TemplateContollerCompleted, corev1.ConditionFalse, "", "")
		mcfgv1.SetControllerConfigStatusCondition(&cfg.Status, *acond)
		rcond := mcfgv1.NewControllerConfigStatusCondition(mcfgv1.TemplateContollerRunning, corev1.ConditionFalse, "", "")
		mcfgv1.SetControllerConfigStatusCondition(&cfg.Status, *rcond)
		cfg.Status.ObservedGeneration = ctrlconfig.GetGeneration()
		return nil
	}
	if err := updateControllerConfigStatus(ctrlconfig.GetName(), ctrl.ccLister.Get, ctrl.client.MachineconfigurationV1().ControllerConfigs(), updateFunc); err != nil {
		return fmt.Errorf("failed to sync status for %v", oerr)
	}
	return oerr
}
func (ctrl *Controller) syncCompletedStatus(ctrlconfig *mcfgv1.ControllerConfig) error {
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
	updateFunc := func(cfg *mcfgv1.ControllerConfig) error {
		reason := fmt.Sprintf("sync completed towards (%d) generation using controller version %s", cfg.GetGeneration(), version.Version)
		acond := mcfgv1.NewControllerConfigStatusCondition(mcfgv1.TemplateContollerCompleted, corev1.ConditionTrue, reason, "")
		mcfgv1.SetControllerConfigStatusCondition(&cfg.Status, *acond)
		rcond := mcfgv1.NewControllerConfigStatusCondition(mcfgv1.TemplateContollerRunning, corev1.ConditionFalse, "", "")
		mcfgv1.SetControllerConfigStatusCondition(&cfg.Status, *rcond)
		fcond := mcfgv1.NewControllerConfigStatusCondition(mcfgv1.TemplateContollerFailing, corev1.ConditionFalse, "", "")
		mcfgv1.SetControllerConfigStatusCondition(&cfg.Status, *fcond)
		cfg.Status.ObservedGeneration = ctrlconfig.GetGeneration()
		return nil
	}
	return updateControllerConfigStatus(ctrlconfig.GetName(), ctrl.ccLister.Get, ctrl.client.MachineconfigurationV1().ControllerConfigs(), updateFunc)
}

type updateControllerConfigStatusFunc func(*mcfgv1.ControllerConfig) error

func updateControllerConfigStatus(name string, controllerConfigGetter func(name string) (*mcfgv1.ControllerConfig, error), client mcfgclientv1.ControllerConfigInterface, updateFuncs ...updateControllerConfigStatusFunc) error {
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
	return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		old, err := controllerConfigGetter(name)
		if err != nil {
			return err
		}
		new := old.DeepCopy()
		for _, update := range updateFuncs {
			if err := update(new); err != nil {
				return err
			}
		}
		if equality.Semantic.DeepEqual(old, new) {
			return nil
		}
		_, err = client.UpdateStatus(new)
		return err
	})
}

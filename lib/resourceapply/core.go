package resourceapply

import (
	"github.com/openshift/machine-config-operator/lib/resourcemerge"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	coreclientv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func ApplyServiceAccount(client coreclientv1.ServiceAccountsGetter, required *corev1.ServiceAccount) (*corev1.ServiceAccount, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	existing, err := client.ServiceAccounts(required.Namespace).Get(required.Name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		actual, err := client.ServiceAccounts(required.Namespace).Create(required)
		return actual, true, err
	}
	if err != nil {
		return nil, false, err
	}
	modified := resourcemerge.BoolPtr(false)
	resourcemerge.EnsureObjectMeta(modified, &existing.ObjectMeta, required.ObjectMeta)
	if !*modified {
		return existing, false, nil
	}
	actual, err := client.ServiceAccounts(required.Namespace).Update(existing)
	return actual, true, err
}
func ApplySecret(client coreclientv1.SecretsGetter, required *corev1.Secret) (*corev1.Secret, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	existing, err := client.Secrets(required.Namespace).Get(required.Name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		actual, err := client.Secrets(required.Namespace).Create(required)
		return actual, true, err
	}
	if err != nil {
		return nil, false, err
	}
	modified := resourcemerge.BoolPtr(false)
	resourcemerge.EnsureObjectMeta(modified, &existing.ObjectMeta, required.ObjectMeta)
	actual, err := client.Secrets(required.Namespace).Update(existing)
	return actual, true, err
}

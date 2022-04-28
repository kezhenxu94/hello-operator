package v1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var hellooperatorlog = logf.Log.WithName("hellooperator-resource")

func (r *HelloOperator) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-webapp-io-github-kezhenxu94-v1-hellooperator,mutating=true,failurePolicy=fail,sideEffects=None,groups=webapp.io.github.kezhenxu94,resources=hellooperators,verbs=create;update,versions=v1,name=mhellooperator.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &HelloOperator{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *HelloOperator) Default() {
	hellooperatorlog.Info("default", "name", r.Name)

	if r.Spec.Name == "" {
		r.Spec.Name = "hello-operator"
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-webapp-io-github-kezhenxu94-v1-hellooperator,mutating=false,failurePolicy=fail,sideEffects=None,groups=webapp.io.github.kezhenxu94,resources=hellooperators,verbs=create;update,versions=v1,name=vhellooperator.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &HelloOperator{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *HelloOperator) ValidateCreate() error {
	hellooperatorlog.Info("validate create", "name", r.Name)

	if r.Spec.Name == "USA" {
		return fmt.Errorf("name cannot be USA")
	}

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *HelloOperator) ValidateUpdate(old runtime.Object) error {
	hellooperatorlog.Info("validate update", "name", r.Name)

	if r.Spec.Name == "USA" {
		return fmt.Errorf("name cannot be updated to USA")
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *HelloOperator) ValidateDelete() error {
	hellooperatorlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

package controllers

import (
	"context"
	"reflect"

	webappv1 "github.com/kezhenxu94/hello-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// HelloOperatorReconciler reconciles a HelloOperator object
type HelloOperatorReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=webapp.io.github.kezhenxu94,resources=hellooperators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=webapp.io.github.kezhenxu94,resources=hellooperators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=webapp.io.github.kezhenxu94,resources=hellooperators/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the HelloOperator object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *HelloOperatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	instance := &webappv1.HelloOperator{}
	if err := r.Get(context.TODO(), req.NamespacedName, instance); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name + "-deployment",
			Namespace: instance.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"deployment": instance.Name + "-deployment"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"deployment": instance.Name + "-deployment"}},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "hello",
							Image: "kezhenxu94/hello:2", // Built from Dockerfile.hello
							Env: []corev1.EnvVar{
								{
									Name:  "NAME",
									Value: instance.Spec.Name,
								},
							},
						},
					},
				},
			},
		},
	}
	if err := controllerutil.SetControllerReference(instance, deploy, r.Scheme); err != nil {
		return reconcile.Result{}, err
	}
	found := &appsv1.Deployment{}
	err := r.Get(context.TODO(), types.NamespacedName{Name: deploy.Name, Namespace: deploy.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		err = r.Create(context.TODO(), deploy)
		if err != nil {
			return reconcile.Result{}, err
		}
		r.Recorder.Eventf(instance, "Normal", "DeploymentCreated", "The Deployment %s has been created", deploy.Name)
	} else if err != nil {
		return reconcile.Result{}, err
	}

	if !reflect.DeepEqual(deploy.Spec, found.Spec) {
		found.Spec = deploy.Spec
		err = r.Update(context.TODO(), found)
		if err != nil {
			return reconcile.Result{}, err
		}
	}
	return reconcile.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *HelloOperatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.HelloOperator{}).
		Complete(r)
}

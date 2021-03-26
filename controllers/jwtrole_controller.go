package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	vaultv1 "github.com/ibm/vault-go/api/v1"
)

// JWTRoleReconciler reconciles a JWTRole object
type JWTRoleReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=vault.vault-go.ibm.com,resources=jwtroles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=vault.vault-go.ibm.com,resources=jwtroles/status,verbs=get;update;patch

func (r *JWTRoleReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("jwtrole", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *JWTRoleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&vaultv1.JWTRole{}).
		Complete(r)
}

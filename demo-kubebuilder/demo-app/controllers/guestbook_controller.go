/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	bookv1 "my.domain/guestbook/api/v1"
	webappv1 "my.domain/guestbook/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// GuestbookReconciler reconciles a Guestbook object
type GuestbookReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=webapp.my.domain,resources=guestbooks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=webapp.my.domain,resources=guestbooks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=webapp.my.domain,resources=guestbooks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Guestbook object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.1/pkg/reconcile
var c controller.Controller
var ctrlLog = ctrl.Log.WithName("controllerlog")

func (r *GuestbookReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	// reqLogger := r.log.WithValues("at", req.NamespacedName)
	// reqLogger.Info("=== Reconciling At")

	// err := controller.Controller.Watch(
	// 	&source.Kind{Type: &appsv1.ReplicaSet{}},
	// 	&handler.EnqueueRequestForOwner{},
	// 	// OwnerType:    &appsv1.Deployment{},
	// 	// IsController: false, // See here that we are setting it as false.
	// 	//singleObjectPredicate{Namespace: meta.GetNamespace(), Name: meta.GetName()},
	// )
	ctrlLog.Info("Going to watch deployments...!")
	found := &bookv1.Guestbook{}
	//err := r.Get(context.TODO(), types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, found)
	// if err != nil {
	// 	ctrlLog.Info("Error occured:", err.Error())
	// }
	if err := r.Get(ctx, types.NamespacedName{Namespace: req.Namespace, Name: req.Name}, found); err != nil {
		ctrlLog.Error(err, "get deployment instance")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GuestbookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Guestbook{}).
		Complete(r)
}

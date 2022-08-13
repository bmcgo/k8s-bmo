/*
Copyright 2022 The BMCGO Authors.

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
	"github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	bmov1alpha1 "github.com/bmcgo/k8s-bmo/api/v1alpha1"
)

// SystemReconciler reconciles a System object
type SystemReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=bmo.bmcgo.dev,resources=systems,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=bmo.bmcgo.dev,resources=systems/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=bmo.bmcgo.dev,resources=systems/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the System object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *SystemReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	system := bmov1alpha1.System{}
	err := r.Get(ctx, req.NamespacedName, &system)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			l.Info("Deleted System")
			return ctrl.Result{Requeue: false}, nil
		}
		return r.requeue(err)
	}
	if system.Spec.State == bmov1alpha1.DesiredStateNotManaged {
		if system.Status.State != bmov1alpha1.ActualStateNotManaged {
			system.Status.State = bmov1alpha1.ActualStateNotManaged
			err = r.Status().Update(ctx, &system)
			if err != nil {
				return r.requeue(errors.Wrap(err, "failed to update status"))
			}
		}
		l.Info("System not managed")
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

func (r *SystemReconciler) requeue(err error) (ctrl.Result, error) {
	return ctrl.Result{Requeue: true}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *SystemReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&bmov1alpha1.System{}).
		Complete(r)
}

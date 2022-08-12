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
	"github.com/bmcgo/k8s-bmo/redfish"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	bmov1alpha1 "github.com/bmcgo/k8s-bmo/api/v1alpha1"
)

// RedfishEndpointReconciler reconciles a RedfishEndpoint object
type RedfishEndpointReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=bmo.bmcgo.dev,resources=redfishendpoints,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=bmo.bmcgo.dev,resources=redfishendpoints/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=bmo.bmcgo.dev,resources=redfishendpoints/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RedfishEndpoint object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *RedfishEndpointReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	endpoint := bmov1alpha1.RedfishEndpoint{}
	err := r.Get(ctx, req.NamespacedName, &endpoint)
	if err != nil {
		if errors.IsNotFound(err) {
			l.Info("Deleted RedfishEndpoint", "NamespacedName", req.NamespacedName)
			return ctrl.Result{Requeue: false}, nil
		}
		return r.requeue(err)
	}
	if !endpoint.DeletionTimestamp.IsZero() {
		return r.handleDelete(ctx, endpoint, l)
	}

	err = r.ensureFinalizer(ctx, endpoint)
	if err != nil {
		return r.requeue(err)
	}

	if len(endpoint.Status.SystemsDiscovered) == 0 {
		rc := redfish.NewClient(redfish.ClientConfig{URL: endpoint.Spec.EndpointURL})
		systemsDiscovered, err := rc.GetSystems()
		if err != nil {
			return r.requeue(err)
		}
		for _, s := range systemsDiscovered {
			endpoint.Status.SystemsDiscovered = append(endpoint.Status.SystemsDiscovered, bmov1alpha1.System{
				Name: s.Name,
				UUID: s.UUID,
			})
		}
		endpoint.Status.LastUpdated = metav1.Now()
		err = r.Status().Update(ctx, &endpoint)
		if err != nil {
			return r.requeue(err)
		}
	}

	return ctrl.Result{}, nil
}

func (r *RedfishEndpointReconciler) requeue(err error) (ctrl.Result, error) {
	return ctrl.Result{Requeue: true}, err
}

func (r *RedfishEndpointReconciler) handleDelete(
	ctx context.Context,
	endpoint bmov1alpha1.RedfishEndpoint,
	l logr.Logger) (ctrl.Result, error) {
	//TODO:
	return ctrl.Result{Requeue: false, RequeueAfter: time.Second * 10}, nil
}

func (r *RedfishEndpointReconciler) ensureFinalizer(ctx context.Context, endpoint bmov1alpha1.RedfishEndpoint) error {
	//TODO:
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RedfishEndpointReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&bmov1alpha1.RedfishEndpoint{}).
		Complete(r)
}

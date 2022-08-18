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
	"github.com/bmcgo/k8s-bmo/ipmitool"
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

// IPMIEndpointReconciler reconciles a IPMIEndpoint object
type IPMIEndpointReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=bmo.bmcgo.dev,resources=ipmiendpoints,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=bmo.bmcgo.dev,resources=ipmiendpoints/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=bmo.bmcgo.dev,resources=ipmiendpoints/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the IPMIEndpoint object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *IPMIEndpointReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	endpoint := bmov1alpha1.IPMIEndpoint{}
	err := r.Get(ctx, req.NamespacedName, &endpoint)
	if err != nil {
		if errors.IsNotFound(err) {
			l.Info("Deleted IPMIEndpoint")
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

	bmn := bmov1alpha1.BareMetalNode{}
	err = r.Get(ctx, client.ObjectKey{Namespace: endpoint.Namespace, Name: endpoint.Name}, &bmn)
	if err != nil {
		if errors.IsNotFound(err) {
			return r.createNode(ctx, endpoint, l)
		}
	}
	return ctrl.Result{}, nil
}

func (r *IPMIEndpointReconciler) createNode(ctx context.Context, endpoint bmov1alpha1.IPMIEndpoint, l logr.Logger) (ctrl.Result, error) {
	it := ipmitool.New(endpoint.Spec.Host, int(endpoint.Spec.Port), endpoint.Spec.Username, endpoint.Spec.Password)
	_, err := it.GetChassisStatus()
	if err != nil {
		l.Error(err, "failed to check new ipmi endpoint")
		endpoint.Status.ErrorMessage = err.Error()
		updateErr := r.Status().Update(ctx, &endpoint)
		if updateErr != nil {
			l.Error(err, "failed to update ipmi endpoint status")
			return ctrl.Result{Requeue: true, RequeueAfter: time.Minute * 1}, updateErr
		}
		return ctrl.Result{Requeue: true, RequeueAfter: time.Minute * 1}, err
	} else {
		bareMetalHost := bmov1alpha1.BareMetalNode{
			ObjectMeta: metav1.ObjectMeta{
				Name:      endpoint.Name,
				Namespace: endpoint.Namespace,
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion: endpoint.APIVersion,
						Kind:       endpoint.Kind,
						Name:       endpoint.Name,
						UID:        endpoint.UID,
					},
				},
			},
			Spec: bmov1alpha1.BareMetalNodeSpec{
				State: bmov1alpha1.DesiredStateNotManaged,
			},
			Status: bmov1alpha1.BareMetalNodeStatus{
				State: bmov1alpha1.ActualStateNotManaged,
			},
		}
		err = r.Create(ctx, &bareMetalHost)
		if err != nil {
			return ctrl.Result{Requeue: true, RequeueAfter: time.Second * 30}, err
		}
		endpoint.Status.ErrorMessage = ""
		err = r.Status().Update(ctx, &endpoint)
		if err != nil {
			return ctrl.Result{Requeue: true, RequeueAfter: time.Second * 30}, err
		}
	}
	return ctrl.Result{}, nil
}

func (r *IPMIEndpointReconciler) ensureFinalizer(ctx context.Context, endpoint bmov1alpha1.IPMIEndpoint) error {
	//TODO:
	return nil
}

func (r *IPMIEndpointReconciler) requeue(err error) (ctrl.Result, error) {
	return ctrl.Result{Requeue: true}, err
}

func (r *IPMIEndpointReconciler) handleDelete(
	ctx context.Context,
	endpoint bmov1alpha1.IPMIEndpoint,
	l logr.Logger) (ctrl.Result, error) {
	//TODO:
	return ctrl.Result{Requeue: false, RequeueAfter: time.Second * 10}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *IPMIEndpointReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&bmov1alpha1.IPMIEndpoint{}).
		Complete(r)
}

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
	"errors"
	"github.com/bmcgo/k8s-bmo/ipmitool"
	"github.com/go-logr/logr"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	bmov1alpha1 "github.com/bmcgo/k8s-bmo/api/v1alpha1"
)

// BareMetalNodeReconciler reconciles a BareMetalNode object
type BareMetalNodeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=bmo.bmcgo.dev,resources=baremetalnodes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=bmo.bmcgo.dev,resources=baremetalnodes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=bmo.bmcgo.dev,resources=baremetalnodes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the BareMetalNode object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *BareMetalNodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	bmn := bmov1alpha1.BareMetalNode{}
	err := r.Get(ctx, req.NamespacedName, &bmn)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			l.Info("Deleted BareMetalNode")
			return ctrl.Result{Requeue: false}, nil
		}
		return r.requeueIfError(err)
	}
	if bmn.Spec.State == bmov1alpha1.DesiredState(bmn.Status.State) {
		l.Info("BareMetalNode has consistent state. No action.", "state", bmn.Status.State)
		return ctrl.Result{}, nil
	}

	l.Info("Action required", "desired state", bmn.Spec.State, "actual state", bmn.Status.State)

	switch bmn.Spec.State {
	case bmov1alpha1.DesiredStateNotManaged:
		bmn.Status.State = bmov1alpha1.ActualStateNotManaged
		return r.requeueIfError(r.Status().Update(ctx, &bmn))
	case bmov1alpha1.DesiredStatePowerOff:
		return r.handlePowerOff(ctx, bmn)
	case bmov1alpha1.DesiredStateInspected:
		return r.handleInspect(ctx, bmn, l)
	default:
		l.Error(errors.New("not implemented"), "not implemented")
		return ctrl.Result{}, nil
	}
}

func (r *BareMetalNodeReconciler) handleInspect(
	ctx context.Context,
	bmn bmov1alpha1.BareMetalNode,
	l logr.Logger) (ctrl.Result, error) {
	ic, err := bmn.GetIPMIClient(ctx, r.Client)
	if err != nil {
		return r.requeueIfError(err)
	}

	if bmn.Status.State == bmov1alpha1.ActualStateInspecting {
		bootDev, err := ic.GetBootDev()
		if err != nil {
			return r.requeueIfError(err)
		}
		if bootDev == ipmitool.BootDevPxe {
			status, err := ic.GetChassisStatus()
			if err != nil {
				return r.requeueIfError(err)
			}
			if status.SystemPower == ipmitool.ChassisStatusPowerOn {
				l.Info("Node is inspecting.")
				return ctrl.Result{}, nil
			} else {
				l.Info("Node was powered off. Turning on.")
				err = ic.Power(ipmitool.ChassisPowerOn)
				return r.requeueIfError(err)
			}
		} else {
			l.Info("Invalid node boot device", "boot device", bootDev)
		}
	}

	err = ic.Power(ipmitool.ChassisPowerOff)
	if err != nil {
		return r.requeueIfError(err)
	}
	l.Info("Node powered off.")
	err = ic.SetBootDev(ipmitool.BootDevPxe)
	if err != nil {
		return r.requeueIfError(err)
	}
	l.Info("Set node boot device to PXE.")
	err = ic.Power(ipmitool.ChassisPowerOn)
	if err != nil {
		return r.requeueIfError(err)
	}
	l.Info("Node powered on.")
	bmn.Status.State = bmov1alpha1.ActualStateInspecting
	bmn.Status.LastUpdate = metav1.Now()
	return r.requeueIfError(r.Status().Update(ctx, &bmn))
}

func (r *BareMetalNodeReconciler) handlePowerOff(ctx context.Context, bmn bmov1alpha1.BareMetalNode) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

func (r *BareMetalNodeReconciler) requeueIfError(err error) (ctrl.Result, error) {
	if err != nil {
		return ctrl.Result{Requeue: true}, err
	}
	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *BareMetalNodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&bmov1alpha1.BareMetalNode{}).
		Complete(r)
}

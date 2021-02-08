/*
Copyright 2021 Kong, Inc.

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

// Code generated by Kong; DO NOT EDIT.

package inputs

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	netv1 "k8s.io/api/networking/v1"
	netv1beta1 "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// -----------------------------------------------------------------------------
// NetV1 Ingress
// -----------------------------------------------------------------------------

// NetV1Ingress reconciles a Ingress object
type NetV1IngressReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// SetupWithManager sets up the controller with the Manager.
func (r *NetV1IngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).For(&netv1.Ingress{}).Complete(r)
}

//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/finalizers,verbs=update

// Reconcile processes the watched objects
func (r *NetV1IngressReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("NetV1Ingress", req.NamespacedName)

	obj := new(netv1.Ingress)
	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if !obj.DeletionTimestamp.IsZero() && time.Now().After(obj.DeletionTimestamp.Time) {
		log.Info("resource is being deleted, its configuration will be removed", "type", "Ingress", "namespace", req.Namespace, "name", req.Name)
		return cleanupObj(ctx, r.Client, log, req.NamespacedName, obj)
	}

	return storeObjUpdates(ctx, r.Client, log, req.NamespacedName, obj)
}

// -----------------------------------------------------------------------------
// NetV1Beta1 Ingress
// -----------------------------------------------------------------------------

// NetV1Beta1Ingress reconciles a Ingress object
type NetV1Beta1IngressReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// SetupWithManager sets up the controller with the Manager.
func (r *NetV1Beta1IngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).For(&netv1beta1.Ingress{}).Complete(r)
}

//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/finalizers,verbs=update

// Reconcile processes the watched objects
func (r *NetV1Beta1IngressReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("NetV1Beta1Ingress", req.NamespacedName)

	obj := new(netv1beta1.Ingress)
	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if !obj.DeletionTimestamp.IsZero() && time.Now().After(obj.DeletionTimestamp.Time) {
		log.Info("resource is being deleted, its configuration will be removed", "type", "Ingress", "namespace", req.Namespace, "name", req.Name)
		return cleanupObj(ctx, r.Client, log, req.NamespacedName, obj)
	}

	return storeObjUpdates(ctx, r.Client, log, req.NamespacedName, obj)
}

// -----------------------------------------------------------------------------
// ExtV1Beta1 Ingress
// -----------------------------------------------------------------------------

// ExtV1Beta1Ingress reconciles a Ingress object
type ExtV1Beta1IngressReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// SetupWithManager sets up the controller with the Manager.
func (r *ExtV1Beta1IngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).For(&extv1beta1.Ingress{}).Complete(r)
}

//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=ingresses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=ingresses/finalizers,verbs=update

// Reconcile processes the watched objects
func (r *ExtV1Beta1IngressReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("ExtV1Beta1Ingress", req.NamespacedName)

	obj := new(extv1beta1.Ingress)
	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if !obj.DeletionTimestamp.IsZero() && time.Now().After(obj.DeletionTimestamp.Time) {
		log.Info("resource is being deleted, its configuration will be removed", "type", "Ingress", "namespace", req.Namespace, "name", req.Name)
		return cleanupObj(ctx, r.Client, log, req.NamespacedName, obj)
	}

	return storeObjUpdates(ctx, r.Client, log, req.NamespacedName, obj)
}

// -----------------------------------------------------------------------------
// KongV1 KongIngress
// -----------------------------------------------------------------------------

// KongV1KongIngress reconciles a Ingress object
type KongV1KongIngressReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// SetupWithManager sets up the controller with the Manager.
func (r *KongV1KongIngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).For(&kongv1.KongIngress{}).Complete(r)
}

//+kubebuilder:rbac:groups=networking.konghq.com,resources=kongingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.konghq.com,resources=kongingresses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=networking.konghq.com,resources=kongingresses/finalizers,verbs=update

// Reconcile processes the watched objects
func (r *KongV1KongIngressReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("KongV1KongIngress", req.NamespacedName)

	obj := new(kongv1.KongIngress)
	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if !obj.DeletionTimestamp.IsZero() && time.Now().After(obj.DeletionTimestamp.Time) {
		log.Info("resource is being deleted, its configuration will be removed", "type", "KongIngress", "namespace", req.Namespace, "name", req.Name)
		return cleanupObj(ctx, r.Client, log, req.NamespacedName, obj)
	}

	return storeObjUpdates(ctx, r.Client, log, req.NamespacedName, obj)
}

/*


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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	jxv4beta1 "github.com/chrismellard/crd/api/v4beta1"
)

// PlatformFeatureReconciler reconciles a PlatformFeature object
type PlatformFeatureReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=jx.jenkins-x.io,resources=platformfeatures,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=jx.jenkins-x.io,resources=platformfeatures/status,verbs=get;update;patch

func (r *PlatformFeatureReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("platformfeature", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *PlatformFeatureReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&jxv4beta1.PlatformFeature{}).
		Complete(r)
}

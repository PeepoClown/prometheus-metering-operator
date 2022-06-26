package controllers

import (
	"context"

	"gopkg.in/yaml.v3"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1alpha1 "github.com/peepoclown/prometheus-metering-operator/api/v1alpha1"
)

const (
	PrometheusImage string = "prom/prometheus:latest"
)

const (
	PrometheusMemoryRequest = "256Mi"
	PrometheusCpuRequest    = "200m"
	PrometheusMemoryLimit   = "1024Mi"
	PrometheusCpuLimit      = "500m"
)

// PrometheusMeteringReconciler reconciles a PrometheusMetering object
type PrometheusMeteringReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps.ru.prometheus-metering-operator,resources=prometheusmeterings,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.ru.prometheus-metering-operator,resources=prometheusmeterings/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.ru.prometheus-metering-operator,resources=prometheusmeterings/finalizers,verbs=update
func (r *PrometheusMeteringReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	prometheusMetering := &appsv1alpha1.PrometheusMetering{}
	err := r.Client.Get(ctx, req.NamespacedName, prometheusMetering)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Log.Info("PrometheusMetering resource not found")
			return ctrl.Result{}, nil
		}
		log.Log.Error(err, "Failed to get PrometheusMetering")
		return ctrl.Result{}, err
	}

	foundPromentheusDeployment := &appsv1.Deployment{}
	err = r.Client.Get(ctx, types.NamespacedName{Name: prometheusMetering.Spec.Name, Namespace: prometheusMetering.Spec.Namespace},
		foundPromentheusDeployment)
	if err == nil {
		return ctrl.Result{}, nil
	}

	if err != nil && errors.IsNotFound(err) {
		promentheusDeployment := r.deployePrometheus(ctx, prometheusMetering)
		log.Log.Info("Creating a new Deployment", "Deployment.Namespace", promentheusDeployment.Namespace, "Deployment.Name", promentheusDeployment.Name)
		err = r.Client.Create(ctx, promentheusDeployment)
		if err != nil {
			log.Log.Error(err, "Failed to create new Deployment", "Deployment.Namespace", promentheusDeployment.Namespace, "Deployment.Name", promentheusDeployment.Name)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		log.Log.Error(err, "Failed to get Deployment")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *PrometheusMeteringReconciler) deployePrometheus(ctx context.Context, prometheusMetering *appsv1alpha1.PrometheusMetering) *appsv1.Deployment {
	name := prometheusMetering.Spec.Name
	labels := map[string]string{"app": name}
	prometheusConfigMap := r.createPrometheusConfigMap(ctx, prometheusMetering)

	promentheusDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: prometheusMetering.Spec.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &prometheusMetering.Spec.ReplicasCount,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: PrometheusImage,
						Name:  name,
						Args:  []string{"--config.file=/etc/prometheus/prometheus.yml"},
						VolumeMounts: []corev1.VolumeMount{{
							Name:      "prometheus-config",
							MountPath: "/etc/prometheus",
						}, {
							Name:      "prometheus-data",
							MountPath: "/prometheus",
						}},
					}},
					Volumes: []corev1.Volume{{
						Name: "prometheus-config",
						VolumeSource: corev1.VolumeSource{
							ConfigMap: &corev1.ConfigMapVolumeSource{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: prometheusConfigMap.Name,
								},
							},
						},
					}, {
						Name: "prometheus-data",
						VolumeSource: corev1.VolumeSource{
							EmptyDir: &corev1.EmptyDirVolumeSource{},
						},
					}},
				},
			},
		},
	}
	ctrl.SetControllerReference(prometheusMetering, promentheusDeployment, r.Scheme)
	return promentheusDeployment
}

func (r *PrometheusMeteringReconciler) createPrometheusConfigMap(ctx context.Context, prometheusMetering *appsv1alpha1.PrometheusMetering) *corev1.ConfigMap {
	cmName := "prometheus-" + prometheusMetering.Spec.Name + "-config"

	data, err := yaml.Marshal(prometheusMetering.Spec.PrometheusConfig)
	if err != nil {
		log.Log.Error(err, "Failed to serialize prometheus config")
		return nil
	}

	prometheusConfigMap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cmName,
			Namespace: prometheusMetering.Spec.Namespace,
		},
		Data: map[string]string{
			"prometheus.yml": string(data),
		},
	}
	r.Client.Create(ctx, prometheusConfigMap)
	return prometheusConfigMap
}

// SetupWithManager sets up the controller with the Manager.
func (r *PrometheusMeteringReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1alpha1.PrometheusMetering{}).
		Complete(r)
}

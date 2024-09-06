package manifests

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewAlertmanagerDeployment creates a Deployment for Alertmanager
func NewAlertmanagerDeployment(opts Options) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "alertmanager",
			Labels: map[string]string{
				"app": "alertmanager",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "alertmanager",
				},
			},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RollingUpdateDeploymentStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "alertmanager",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "alertmanager",
							Image: "quay.io/prometheus/alertmanager:latest",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 9093,
								},
							},
							ImagePullPolicy: corev1.PullAlways,
						},
					},
					RestartPolicy: corev1.RestartPolicyAlways,
					TerminationGracePeriodSeconds: int64Ptr(30),
				},
			},
		},
	}
}

// NewAlertmanagerService creates a Service for Alertmanager
func NewAlertmanagerService(opts Options) *corev1.Service {
	serviceName := "alertmanager"
	labels := map[string]string{
		"app": "alertmanager",
	}

	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: corev1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   serviceName,
			Labels: labels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "web",
					Port:       9093,
					Protocol:   corev1.ProtocolTCP,
					TargetPort: intstr.FromInt(9093),
				},
			},
			Selector: labels,
			Type:     corev1.ServiceTypeNodePort,
			ExternalTrafficPolicy: corev1.ServiceExternalTrafficPolicyTypeCluster,
		},
	}
}

// NewAlertmanagerPodDisruptionBudget returns a PodDisruptionBudget for Alertmanager pods.
func NewAlertmanagerPodDisruptionBudget(opts Options) *policyv1.PodDisruptionBudget {
	l := map[string]string{
		"app": "alertmanager",
	}

	ma := intstr.FromInt(1)

	return &policyv1.PodDisruptionBudget{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PodDisruptionBudget",
			APIVersion: policyv1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels:    l,
			Name:      "alertmanager",
			Namespace: opts.Namespace,
		},
		Spec: policyv1.PodDisruptionBudgetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: l,
			},
			MinAvailable: &ma,
		},
	}
}

// BuildAlertmanager constructs Alertmanager Deployment, Service, and PDB
func BuildAlertmanager(opts Options) ([]client.Object, error) {
	deployment := NewAlertmanagerDeployment(opts)
	service := NewAlertmanagerService(opts)
	pdb := NewAlertmanagerPodDisruptionBudget(opts)

	return []client.Object{deployment, service, pdb}, nil
}

func int32Ptr(i int32) *int32 { return &i }
func int64Ptr(i int64) *int64 { return &i }

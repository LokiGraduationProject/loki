package manifests

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func BuildAlertmanager(opts Options) ([]client.Object, error) {
    // Alertmanager Deployment 생성
    alertmanagerDeployment := &appsv1.Deployment{
        ObjectMeta: metav1.ObjectMeta{
            Name:      "alertmanager",
            Namespace: "default",
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

    // Alertmanager Service 생성
    alertmanagerService := &corev1.Service{
        ObjectMeta: metav1.ObjectMeta{
            Name:      "alertmanager",
            Namespace: "default",
            Labels: map[string]string{
                "app": "alertmanager",
            },
        },
        Spec: corev1.ServiceSpec{
            Type: corev1.ServiceTypeNodePort,
            Selector: map[string]string{
                "app": "alertmanager",
            },
            Ports: []corev1.ServicePort{
                {
                    Name:       "web",
                    Port:       9093,
                    TargetPort: intstr.FromInt(9093),
                    NodePort:   30093,
                    Protocol:   corev1.ProtocolTCP,
                },
            },
            ExternalTrafficPolicy: corev1.ServiceExternalTrafficPolicyTypeCluster,
            IPFamilies:            []corev1.IPFamily{corev1.IPv4Protocol},
        },
    }

    return []client.Object{alertmanagerDeployment, alertmanagerService}, nil
}

func int32Ptr(i int32) *int32 { return &i }
func int64Ptr(i int64) *int64 { return &i }

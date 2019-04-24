package configmaps

import (
	v1alpha1 "github.com/aneeshkp/smartgateway-operator/pkg/apis/app/v1alpha1"
	"github.com/aneeshkp/smartgateway-operator/pkg/utils/configs"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Create NewConfigMapForCR method to create configmap
func NewConfigMapForCR(m *v1alpha1.SmartGateway) *corev1.ConfigMap {
	config := configs.ConfigForSmartGateway(m)
	configMap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "ConfigMap",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name,
			Namespace: m.Namespace,
		},
		Data: map[string]string{
			"smartgateway.conf.template": config,
		},
	}

	return configMap
}

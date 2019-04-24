package selectors

import (
	"k8s.io/apimachinery/pkg/labels"
)

const (
	LabelAppKey = "application"

	LabelResourceKey = "smartgateway_cr"
)

// Set labels in a map
func LabelsForSmartGateway(name string) map[string]string {
	return map[string]string{
		LabelAppKey:      name,
		LabelResourceKey: name,
	}
}

// return a selector that matches resources for a smartgateway resource
func ResourcesBySmartGatewayName(name string) labels.Selector {
	set := map[string]string{
		LabelAppKey:      name,
		LabelResourceKey: name,
	}
	return labels.SelectorFromSet(set)
}

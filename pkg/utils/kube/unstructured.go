package kube

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ToUnstructured(obj any) unstructured.Unstructured {
	_ = "STUB: not implemented"
	return *new(unstructured.Unstructured)
}

package kube

import (
	corev1 "k8s.io/api/core/v1"
)

func Namespace(name string) corev1.Namespace {
	_ = "STUB: not implemented"
	return *new(corev1.Namespace)
}

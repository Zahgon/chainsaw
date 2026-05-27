package testing

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type FakeLoader struct {
	LoadFn   func(int, []byte) (schema.GroupVersionKind, unstructured.Unstructured, error)
	numCalls int
}

func (f *FakeLoader) Load(data []byte) (schema.GroupVersionKind, unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return *new(schema.GroupVersionKind), *new(unstructured.Unstructured), nil
}

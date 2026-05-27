package config

import (
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha2"
	"github.com/kyverno/pkg/ext/resource/loader"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	DefaultFileName = ".chainsaw.yaml"
)

type (
	splitter      = func([]byte) ([][]byte, error)
	loaderFactory = func() (loader.Loader, error)
	converter     = func(schema.GroupVersionKind, unstructured.Unstructured) (*v1alpha2.Configuration, error)
)

var (
	configuration_v1alpha1 = schema.GroupVersion(v1alpha1.GroupVersion).WithKind("Configuration")
	configuration_v1alpha2 = schema.GroupVersion(v1alpha2.GroupVersion).WithKind("Configuration")
)

func Load(path string) (*v1alpha2.Configuration, error) { _ = "STUB: not implemented"; return nil, nil }

func LoadBytes(content []byte) (*v1alpha2.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Parse(content []byte) ([]*v1alpha2.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parse(content []byte, splitter splitter, loaderFactory loaderFactory, converter converter) ([]*v1alpha2.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func defaultConverter(gvk schema.GroupVersionKind, untyped unstructured.Unstructured) (*v1alpha2.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

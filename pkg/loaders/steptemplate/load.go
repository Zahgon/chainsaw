package steptemplate

import (
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/pkg/ext/resource/loader"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type (
	splitter      = func([]byte) ([][]byte, error)
	loaderFactory = func() (loader.Loader, error)
	converter     = func(unstructured.Unstructured) (*v1alpha1.StepTemplate, error)
)

var stepTemplate_v1alpha1 = schema.GroupVersion(v1alpha1.GroupVersion).WithKind("StepTemplate")

func Load(path string, remarshal bool) ([]*v1alpha1.StepTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Parse(content []byte, remarshal bool) ([]*v1alpha1.StepTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parse(content []byte, remarshal bool, splitter splitter, loaderFactory loaderFactory, converter converter) ([]*v1alpha1.StepTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

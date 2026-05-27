package errors

import (
	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

type resourceError struct {
	compilers compilers.Compilers
	expected  unstructured.Unstructured
	actual    unstructured.Unstructured
	template  bool
	bindings  apis.Bindings
	errs      field.ErrorList
}

func ResourceError(
	compilers compilers.Compilers,
	expected unstructured.Unstructured,
	actual unstructured.Unstructured,
	template bool,
	bindings apis.Bindings,
	errs field.ErrorList,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (e resourceError) Error() string { _ = "STUB: not implemented"; return "" }

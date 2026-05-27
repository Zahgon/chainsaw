package templating

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func Template(ctx context.Context, compilers compilers.Compilers, tpl v1alpha1.Projection, value any, bindings apis.Bindings) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func TemplateAndMerge(ctx context.Context, compilers compilers.Compilers, obj unstructured.Unstructured, bindings apis.Bindings, templates ...v1alpha1.Projection) (unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return *new(unstructured.Unstructured), nil
}

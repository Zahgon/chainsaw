package templating

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ResourceRef(ctx context.Context, c compilers.Compilers, obj *unstructured.Unstructured, bindings apis.Bindings) error {
	_ = "STUB: not implemented"
	return nil
}

// this is not a valid resource (non resource assertion maybe ?)

func getStringLabels(obj *unstructured.Unstructured) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

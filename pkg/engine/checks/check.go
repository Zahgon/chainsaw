package checks

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func Check(ctx context.Context, compilers compilers.Compilers, obj any, bindings apis.Bindings, check *v1alpha1.Check) (field.ErrorList, error) {
	_ = "STUB: not implemented"
	return *new(field.ErrorList), nil
}

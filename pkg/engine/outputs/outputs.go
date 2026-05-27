package outputs

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
)

type Outputs = map[string]any

func Process(ctx context.Context, compilers compilers.Compilers, tc apis.Bindings, input any, outputs ...v1alpha1.Output) (Outputs, error) {
	_ = "STUB: not implemented"
	return *new(Outputs), nil
}

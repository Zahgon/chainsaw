package expressions

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
)

func String(ctx context.Context, c compilers.Compilers, in string, bindings apis.Bindings) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

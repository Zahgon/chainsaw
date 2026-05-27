package kubectl

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/client"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
)

func Describe(ctx context.Context, compilers compilers.Compilers, client client.Client, tc apis.Bindings, collector *v1alpha1.Describe) (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

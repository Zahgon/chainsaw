package internal

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
)

func RegisterEnvs(ctx context.Context, compilers compilers.Compilers, namespace string, bindings apis.Bindings, envs ...v1alpha1.Binding) (map[string]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

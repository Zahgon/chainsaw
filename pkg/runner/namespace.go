package runner

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	corev1 "k8s.io/api/core/v1"
)

func buildNamespace(ctx context.Context, compilers compilers.Compilers, name string, template *v1alpha1.Projection, tc enginecontext.TestContext) (*corev1.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

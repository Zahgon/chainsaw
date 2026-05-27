package kubectl

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
)

func Logs(ctx context.Context, compilers compilers.Compilers, tc apis.Bindings, collector *v1alpha1.PodLogs) (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

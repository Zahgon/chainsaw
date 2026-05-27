package bindings

import (
	"context"
	"regexp"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
)

var identifier = regexp.MustCompile(`^\w+$`)

func checkBindingName(name string) error { _ = "STUB: not implemented"; return nil }

func RegisterBinding(bindings apis.Bindings, name string, value any) apis.Bindings {
	_ = "STUB: not implemented"
	return *new(apis.Bindings)
}

func ResolveBinding(ctx context.Context, compilers compilers.Compilers, bindings apis.Bindings, input any, variable v1alpha1.Binding) (string, any, error) {
	_ = "STUB: not implemented"
	return "", *new(any), nil
}

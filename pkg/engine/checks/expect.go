package checks

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func Expect(ctx context.Context, compilers compilers.Compilers, obj unstructured.Unstructured, bindings apis.Bindings, expect ...v1alpha1.Expectation) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// if a match is specified, skip the check if the resource doesn't match

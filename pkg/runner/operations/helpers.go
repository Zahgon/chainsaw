package operations

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/cleanup/cleaner"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
	"github.com/kyverno/kyverno-json/pkg/core/compilers"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func fileRefOrResource(ctx context.Context, ref v1alpha1.ActionResourceRef, basePath string, compilers compilers.Compilers, bindings apis.Bindings) ([]unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fileRefOrCheck(ctx context.Context, ref v1alpha1.ActionCheckRef, basePath string, compilers compilers.Compilers, bindings apis.Bindings) ([]unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepareResource(resource unstructured.Unstructured, tc enginecontext.TestContext) error {
	_ = "STUB: not implemented"
	return nil
}

func getCleanerOrNil(cleaner cleaner.CleanerCollector, tc enginecontext.TestContext) cleaner.CleanerCollector {
	_ = "STUB: not implemented"
	return *new(cleaner.CleanerCollector)
}

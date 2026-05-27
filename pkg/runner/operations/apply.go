package operations

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/cleanup/cleaner"
	"github.com/kyverno/chainsaw/pkg/engine/outputs"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type applyAction struct {
	op       v1alpha1.Apply
	resource unstructured.Unstructured
	cleaner  cleaner.CleanerCollector
}

func (o applyAction) Execute(ctx context.Context, tc enginecontext.TestContext) (outputs.Outputs, error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func applyOperation(ctx context.Context, tc enginecontext.TestContext, cleaner cleaner.CleanerCollector, op v1alpha1.Apply) ([]Operation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

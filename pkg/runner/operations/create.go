package operations

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/cleanup/cleaner"
	"github.com/kyverno/chainsaw/pkg/engine/outputs"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type createAction struct {
	op       v1alpha1.Create
	resource unstructured.Unstructured
	cleaner  cleaner.CleanerCollector
}

func (o createAction) Execute(ctx context.Context, tc enginecontext.TestContext) (outputs.Outputs, error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func createOperation(ctx context.Context, tc enginecontext.TestContext, cleaner cleaner.CleanerCollector, op v1alpha1.Create) ([]Operation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

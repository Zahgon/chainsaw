package operations

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/engine/outputs"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type updateAction struct {
	op       v1alpha1.Update
	resource unstructured.Unstructured
}

func (o updateAction) Execute(ctx context.Context, tc enginecontext.TestContext) (outputs.Outputs, error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func updateOperation(ctx context.Context, tc enginecontext.TestContext, op v1alpha1.Update) ([]Operation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

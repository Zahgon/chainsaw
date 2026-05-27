package operations

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/engine/outputs"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
)

type waitAction struct {
	op v1alpha1.Wait
}

func (o waitAction) Execute(ctx context.Context, tc enginecontext.TestContext) (outputs.Outputs, error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

// make sure timeout is set to populate the command flag

// shift operation timeout

func waitOperation(op v1alpha1.Wait) Operation { _ = "STUB: not implemented"; return *new(Operation) }

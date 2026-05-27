package operations

import (
	"context"
	"time"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/engine/outputs"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
)

type sleepAction struct {
	duration time.Duration
}

func (o sleepAction) Execute(ctx context.Context, tc enginecontext.TestContext) (outputs.Outputs, error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func sleepOperation(op v1alpha1.Sleep) Operation { _ = "STUB: not implemented"; return *new(Operation) }

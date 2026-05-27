package operations

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/engine/outputs"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
)

type podLogsAction struct {
	op v1alpha1.PodLogs
}

func (o podLogsAction) Execute(ctx context.Context, tc enginecontext.TestContext) (outputs.Outputs, error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func logsOperation(op v1alpha1.PodLogs) Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

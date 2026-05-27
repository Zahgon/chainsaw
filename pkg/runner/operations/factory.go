package operations

import (
	"context"

	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/cleanup/cleaner"
	"github.com/kyverno/chainsaw/pkg/model"
	enginecontext "github.com/kyverno/chainsaw/pkg/runner/context"
)

func TryOperation(
	ctx context.Context,
	tc enginecontext.TestContext,
	handler v1alpha1.Operation,
	cleaner cleaner.CleanerCollector,
) (model.OperationType, []Operation, error) {
	_ = "STUB: not implemented"
	return *new(model.OperationType), nil, nil
}

func CatchOperation(
	ctx context.Context,
	tc enginecontext.TestContext,
	handler v1alpha1.CatchFinally,
) ([]Operation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

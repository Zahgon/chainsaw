package sleep

import (
	"context"
	"time"

	"github.com/kyverno/chainsaw/pkg/apis"
	"github.com/kyverno/chainsaw/pkg/engine/operations"
	"github.com/kyverno/chainsaw/pkg/engine/outputs"
)

type operation struct {
	duration time.Duration
}

func New(duration time.Duration) operations.Operation {
	_ = "STUB: not implemented"
	return *new(operations.Operation)
}

func (o *operation) Exec(ctx context.Context, _ apis.Bindings) (_ outputs.Outputs, _err error) {
	_ = "STUB: not implemented"
	return *new(outputs.Outputs), nil
}

func (o *operation) execute() error { _ = "STUB: not implemented"; return nil }

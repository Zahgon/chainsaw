package mocks

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/kyverno/chainsaw/pkg/client"
	"github.com/kyverno/chainsaw/pkg/logging"
)

type Logger struct {
	Logs     []string
	numCalls int
}

func (f *Logger) Log(_ context.Context, operation logging.Operation, status logging.Status, obj client.Object, color *color.Color, args ...fmt.Stringer) {
	_ = "STUB: not implemented"
	return
}

func (f *Logger) NumCalls() int { _ = "STUB: not implemented"; return 0 }

package logging

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/kyverno/chainsaw/pkg/client"
)

func Log(ctx context.Context, operation Operation, status Status, obj client.Object, color *color.Color, args ...fmt.Stringer) {
	_ = "STUB: not implemented"
	return
}

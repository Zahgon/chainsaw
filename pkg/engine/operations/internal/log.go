package internal

import (
	"context"
	"fmt"

	"github.com/kyverno/chainsaw/pkg/client"
	"github.com/kyverno/chainsaw/pkg/logging"
)

func LogStart(ctx context.Context, op logging.Operation, obj client.Object, args ...fmt.Stringer) {
	_ = "STUB: not implemented"
	return
}

func LogEnd(ctx context.Context, op logging.Operation, obj client.Object, err error) {
	_ = "STUB: not implemented"
	return
}

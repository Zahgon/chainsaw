package logging

import (
	"context"
)

type (
	sinkKey   struct{}
	loggerKey struct{}
)

func WithSink(ctx context.Context, sink Sink) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getSink(ctx context.Context) Sink { _ = "STUB: not implemented"; return *new(Sink) }

func WithLogger(ctx context.Context, logger Logger) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getLogger(ctx context.Context) Logger { _ = "STUB: not implemented"; return *new(Logger) }

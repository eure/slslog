package slslog

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func ExampleInfof() {
	parent := context.Background()
	s := trace.SpanFromContext(parent)
	ctx := trace.ContextWithSpan(parent, s)
	Infof(ctx, "%s", "test")

	// Output:
	// {"severity":"INFO","message":"test","trace":"service/slslog/trace/00000000000000000000000000000000","span":"service/slslog/span/0000000000000000"}
}

func ExampleWarningf() {
	parent := context.Background()
	s := trace.SpanFromContext(parent)
	ctx := trace.ContextWithSpan(parent, s)
	Warningf(ctx, "%s", "test")

	// Output:
	// {"severity":"WARNING","message":"test","trace":"service/slslog/trace/00000000000000000000000000000000","span":"service/slslog/span/0000000000000000"}
}

func ExampleErrorf() {
	parent := context.Background()
	s := trace.SpanFromContext(parent)
	ctx := trace.ContextWithSpan(parent, s)
	Errorf(ctx, "%s", "test")

	// Output:
	// {"severity":"ERROR","message":"test","trace":"service/slslog/trace/00000000000000000000000000000000","span":"service/slslog/span/0000000000000000"}
}

func ExampleCriticalf() {
	parent := context.Background()
	s := trace.SpanFromContext(parent)
	ctx := trace.ContextWithSpan(parent, s)
	Criticalf(ctx, "%s", "test")

	// Output:
	// {"severity":"CRITICAL","message":"test","trace":"service/slslog/trace/00000000000000000000000000000000","span":"service/slslog/span/0000000000000000"}
}

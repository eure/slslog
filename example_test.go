package slslog

import (
	"context"

	"go.opencensus.io/trace"
)

func ExampleInfof() {
	ctx := trace.NewContext(context.Background(), nil)
	Infof(ctx, "%s", "test")

	// Output:
	// {"severity":"INFO","message":"test","trace":"service/slslog/trace/00000000000000000000000000000000","span":"service/slslog/span/0000000000000000"}
}

func ExampleWarningf() {
	ctx := trace.NewContext(context.Background(), nil)
	Warningf(ctx, "%s", "test")

	// Output:
	// {"severity":"WARNING","message":"test","trace":"service/slslog/trace/00000000000000000000000000000000","span":"service/slslog/span/0000000000000000"}
}

func ExampleErrorf() {
	ctx := trace.NewContext(context.Background(), nil)
	Errorf(ctx, "%s", "test")

	// Output:
	// {"severity":"ERROR","message":"test","trace":"service/slslog/trace/00000000000000000000000000000000","span":"service/slslog/span/0000000000000000"}
}

func ExampleCriticalf() {
	ctx := trace.NewContext(context.Background(), nil)
	Criticalf(ctx, "%s", "test")

	// Output:
	// {"severity":"CRITICAL","message":"test","trace":"service/slslog/trace/00000000000000000000000000000000","span":"service/slslog/span/0000000000000000"}
}

package spancontext

import (
	"context"

	"go.opencensus.io/trace"
)

type LogSpanContext struct {
	SpanID  string
	TraceID string
}

func Get(ctx context.Context) *LogSpanContext {
	sc := trace.FromContext(ctx).SpanContext()
	return &LogSpanContext{
		SpanID:  sc.SpanID.String(),
		TraceID: sc.TraceID.String(),
	}
}

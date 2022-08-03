package spancontext

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type LogSpanContext struct {
	SpanID  string
	TraceID string
}

func Get(ctx context.Context) *LogSpanContext {
	sc := trace.SpanContextFromContext(ctx)

	return &LogSpanContext{
		SpanID:  sc.SpanID().String(),
		TraceID: sc.TraceID().String(),
	}
}

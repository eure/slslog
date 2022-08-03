package spancontext

import (
	"context"

	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type LogSpanContext struct {
	SpanID  string
	TraceID string
}

func Get(ctx context.Context) *LogSpanContext {
	tp := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(tp)
	_, span := otel.Tracer("github.com/slslog").Start(ctx, "logging")
	sc := span.SpanContext()

	return &LogSpanContext{
		SpanID:  sc.SpanID().String(),
		TraceID: sc.TraceID().String(),
	}
}

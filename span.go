package slslog

import (
	"context"

	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// Span wraps go.opentelemetry.io/otel/trace.Span.
type Span struct {
	ctx  context.Context
	span trace.Span
}

// StartSpan starts a new span from the current span in the given context
// and returns it as Span.
// This span can be propagated to the subsequent process by using span's
// context.
func StartSpan(ctx context.Context, label string) *Span {
	tp := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(tp)
	ctx, span := otel.Tracer("github.com/slslog").Start(ctx, "slslog")
	return &Span{
		ctx:  ctx,
		span: span,
	}
}

// Context returns the span context attached to Span.
func (s *Span) Context() context.Context {
	return s.ctx
}

// End ends the span.
func (s *Span) End() {
	s.span.End()
}

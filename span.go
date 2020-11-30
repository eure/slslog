package slslog

import (
	"context"

	"go.opencensus.io/trace"
)

// Span wraps go.opencensus.io/trace.Span.
type Span struct {
	ctx  context.Context
	span *trace.Span
}

// StartSpan starts a new span from the current span in the given context
// and returns it as Span.
// This span can be propagated to the subsequent process by using span's
// context.
func StartSpan(ctx context.Context, label string) *Span {
	ctx, span := trace.StartSpan(ctx, label)
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

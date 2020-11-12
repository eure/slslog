package slslog

import (
	"context"

	"go.opencensus.io/trace"
)

// slslog.Span is wrapper of go.opencensus.io/trace.Span.
type Span struct {
	ctx  context.Context
	span *trace.Span
}

// StartSpan starts a new span of the current span in the context.
//
// This returns the new span which has a new child span and the context.
// which contains the newly created span in the span, and you are able
// to use this context to propagate the returned span in process per lambda
// call.
func StartSpan(ctx context.Context, label string) *Span {
	ctx, span := trace.StartSpan(ctx, label)
	return &Span{
		ctx:  ctx,
		span: span,
	}
}

// Context returns the span context with the given Span attached.
func (s *Span) Context() context.Context {
	return s.ctx
}

// End ends the span.
func (s *Span) End() {
	s.span.End()
}

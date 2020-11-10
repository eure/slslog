package slslog

import (
	"context"

	"go.opencensus.io/trace"
)

type Span struct {
	ctx  context.Context
	span *trace.Span
}

func StartSpan(ctx context.Context, label string) *Span {
	ctx, span := trace.StartSpan(ctx, label)
	return &Span{
		ctx:  ctx,
		span: span,
	}
}

func (s *Span) Context() context.Context {
	return s.ctx
}

func (s *Span) End() {
	s.span.End()
}

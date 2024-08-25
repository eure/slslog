package slslog

import (
	"context"
	"log/slog"
)

type slsLogHandler struct {
	handler slog.Handler
	attrs   []slog.Attr
	groups  []string
}

func (h *slsLogHandler) Handle(ctx context.Context, r slog.Record) error {
	switch r.Level {
	case levelCritical.Level():
		r.AddAttrs(slog.String("severity", "CRITICAL"))
	default:
		r.AddAttrs(slog.String("severity", r.Level.String()))
	}
	a := make([]slog.Attr, 0, r.NumAttrs())
	for attr := range r.Attrs {
		if attr.Key == "msg" {
			continue
		}
		if attr.Key == "time" {
			continue
		}
		if attr.Key == "level" {
			continue
		}
		a = append(a, attr)
	}
	return h.handler.Handle(ctx, r)
}

func (h *slsLogHandler) Enabled(context.Context, slog.Level) bool {
	return true
}

func (h *slsLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &slsLogHandler{
		handler: h.handler,
		attrs:   attrs,
	}
}

func (h *slsLogHandler) WithGroup(group string) slog.Handler {
	return &slsLogHandler{
		handler: h.handler,
		groups:  append(h.groups, group),
	}
}

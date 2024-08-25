package slslog

import (
	"context"
	"log/slog"
	"time"
)

type slsLogHandler struct {
	handler slog.Handler
	attrs   []slog.Attr
	groups  []string
}

func (h *slsLogHandler) Handle(ctx context.Context, r slog.Record) error {
	newRecord := slog.NewRecord(time.Time{}, r.Level, "", r.PC)
	switch r.Level {
	case levelCritical.Level():
		newRecord.AddAttrs(slog.String("severity", "CRITICAL"))
	default:
		newRecord.AddAttrs(slog.String("severity", r.Level.String()))
	}
	for attr := range r.Attrs {
		if attr.Key == "time" {
			continue
		}
		if attr.Key == "msg" {
			continue
		}
		if attr.Key == "level" {
			continue
		}
		newRecord.AddAttrs(attr)
	}
	return h.handler.Handle(ctx, newRecord)
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

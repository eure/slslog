package slslog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
)

type slsLogHandler struct {
	attrs  []slog.Attr
	groups []string
	w      io.Writer
}

func (h *slsLogHandler) Handle(ctx context.Context, r slog.Record) error {
	switch r.Level {
	case levelCritical.Level():
		r.AddAttrs(slog.String("severity", "CRITICAL"))
	default:
		r.AddAttrs(slog.String("severity", r.Level.String()))
	}

	var out string
	r.Attrs(func(attr slog.Attr) bool {
		if attr.Key != "level" && attr.Key != "msg" && attr.Key != "time" {
			out += fmt.Sprintf(`"%s":"%v",`, attr.Key, attr.Value)
		}
		return true
	})

	if len(out) > 0 {
		out = out[:len(out)-1]
	}

	fmt.Printf("{%+v}", out)
	return nil
}

func (h *slsLogHandler) Enabled(context.Context, slog.Level) bool {
	return true
}

func (h *slsLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &slsLogHandler{
		attrs: attrs,
	}
}

func (h *slsLogHandler) WithGroup(group string) slog.Handler {
	return &slsLogHandler{
		groups: append(h.groups, group),
	}
}

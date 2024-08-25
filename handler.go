package slslog

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"slices"
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

	num := r.NumAttrs()
	buf := bytes.NewBuffer(nil)
	buf.WriteString("{")
	for i, attr := range slices.Collect(r.Attrs) {
		if attr.Key == "msg" {
			continue
		}
		if attr.Key == "time" {
			continue
		}
		if attr.Key == "level" {
			continue
		}
		buf.WriteString("\"")
		buf.WriteString(attr.Key)
		buf.WriteString("\":")
		buf.WriteString("\"")
		buf.WriteString(attr.Value.String())
		if i == num-1 {
			buf.WriteString("\"")
			break
		}
		buf.WriteString("\",")
	}
	buf.WriteString("}\n")
	_, err := h.w.Write(buf.Bytes())
	return err
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

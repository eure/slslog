package slslog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type slsLogHandler struct {
	attrs  []slog.Attr
	groups []string
}

func (h *slsLogHandler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(h.attrs...)

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

	fmt.Fprintf(os.Stdout, "{%s}\n", out)
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

/* Package slslog provides the trackable log output with tracing on AWS CloudWatch Logs.

Example:

	slslog.SetLogLabel("log label")                              // e.g. program name
	span := slslog.StartSpan(context.Background(), "span label") // e.g. func name
	defer span.End()

	ctx := span.Context()
	Infof(ctx, "this is slslog output")
*/

package slslog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/eure/slslog/internal/spancontext"
)

type logger struct {
	mu    sync.Mutex
	label string
	c     *slog.Logger
}

var std = &logger{
	label: "slslog",
	c: slog.New(&slsLogHandler{
		handler: slog.NewJSONHandler(os.Stderr, nil),
	}),
}

func SetLogLabel(label string) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.label = label
}

func (l *logger) output(ctx context.Context, level slog.Level, format string, a ...interface{}) {
	std.mu.Lock()
	defer std.mu.Unlock()

	msg := fmt.Sprintf(format, a...)
	sc := spancontext.Get(ctx)
	l.c.LogAttrs(ctx, level, msg,
		slog.String("trace", fmt.Sprintf("service/%s/trace/%s", l.label, sc.TraceID)),
		slog.String("span", fmt.Sprintf("service/%s/span/%s", l.label, sc.SpanID)),
		slog.String("message", msg),
	)
}

// Infof formats its arguments according to the format like fmt.Printf,
// and records the text as log message at Info level.
func Infof(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, slog.LevelInfo, format, a...)
}

// Warningf is like Infof, but the severity is warning level.
func Warningf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, slog.LevelWarn, format, a...)
}

// Errorf is like Infof, but the severity is error level.
func Errorf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, slog.LevelError, format, a...)
}

// Criticalf is like Infof, but the severity is critical level.
func Criticalf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, levelCritical.Level(), format, a...)
}

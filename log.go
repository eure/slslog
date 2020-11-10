package slslog

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/eure/slslog/internal/spancontext"
)

type logger struct {
	mu    sync.Mutex
	label string
}

var std = &logger{
	label: "slslog",
}

func SetLogLabel(label string) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.label = label
}

func (l *logger) output(ctx context.Context, severity, format string, a ...interface{}) {
	std.mu.Lock()
	defer std.mu.Unlock()

	sc := spancontext.Get(ctx)
	payload := struct {
		Severity string `json:"severity"`
		Message  string `json:"message"`
		Trace    string `json:"trace"`
		Span     string `json:"span"`
	}{
		Severity: severity,
		Message:  fmt.Sprintf(format, a...),
		Trace:    fmt.Sprintf("service/%s/trace/%s", l.label, sc.TraceID),
		Span:     fmt.Sprintf("service/%s/span/%s", l.label, sc.SpanID),
	}

	b, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

func Infof(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "INFO", format, a...)
}

func Warningf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "WARNING", format, a...)
}

func Errorf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "ERROR", format, a...)
}

func Criticalf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "CRITICAL", format, a...)
}
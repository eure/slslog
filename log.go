/* Package slslog provides the trackable log output with using trace on AWS CloudWatch Logs.

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

// Infof formats its arguments according to the format like fmt.Printf,
// and records the text as log message at Info level.
func Infof(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "INFO", format, a...)
}

// Warningf is like Infof, but the severity is warning level.
func Warningf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "WARNING", format, a...)
}

// Errorf is like Infof, but the severity is error level.
func Errorf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "ERROR", format, a...)
}

// Criticalf is like Infof, but the severity is critical level.
func Criticalf(ctx context.Context, format string, a ...interface{}) {
	std.output(ctx, "CRITICAL", format, a...)
}

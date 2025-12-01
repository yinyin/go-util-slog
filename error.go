package utilslog

import (
	"context"
	"log/slog"
	"runtime"
	"time"
)

// ErrorContext logs an error message with context.
// The `err` parameter is REQUIRED, it will be included in the log attributes with "error" as its key.
// Use the `.ErrorContext()` in original slog package if you don't want to include an error.
func ErrorContext(ctx context.Context, msg string, err error, attrs ...slog.Attr) {
	logger := slog.Default()
	if ctx == nil {
		ctx = context.Background()
	}
	if !logger.Enabled(ctx, slog.LevelError) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, ErrorContext]
	r := slog.NewRecord(time.Now(), slog.LevelError, msg, pcs[0])
	attrs = append(attrs, slog.Any("error", err))
	if len(attrs) > 0 {
		r.AddAttrs(attrs...)
	}
	_ = logger.Handler().Handle(ctx, r)
}

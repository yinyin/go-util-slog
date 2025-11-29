package utilslog

import (
	"context"
	"log/slog"
	"runtime"
	"time"
)

// WarnContext logs a warning message with context.
// The `err` parameter is REQUIRED, it will be included in the log attributes with "error" as its key.
// Use the `.WarnContext()` in original slog package if you don't want to include an error.
func WarnContext(ctx context.Context, msg string, err error, attrs ...slog.Attr) {
	logger := slog.Default()
	if ctx == nil {
		ctx = context.Background()
	}
	if !logger.Enabled(ctx, slog.LevelWarn) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelWarn, msg, pcs[0])
	attrs = append(attrs, slog.Any("error", err))
	if len(attrs) > 0 {
		r.AddAttrs(attrs...)
	}
	_ = logger.Handler().Handle(context.Background(), r)
}

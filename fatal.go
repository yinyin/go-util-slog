package utilslog

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"
)

// FatalContext logs a fatal error message with context and exits the program with `os.Exit(1)`.
// The `err` parameter is optional; if provided, it will be included in the log attributes with "error" as its key.
func FatalContext(ctx context.Context, msg string, err error, attrs ...slog.Attr) {
	logger := slog.Default()
	if ctx == nil {
		ctx = context.Background()
	}
	if !logger.Enabled(ctx, slog.LevelError) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelError, msg, pcs[0])
	if err != nil {
		attrs = append(attrs, slog.Any("error", err))
	}
	if len(attrs) > 0 {
		r.AddAttrs(attrs...)
	}
	_ = logger.Handler().Handle(context.Background(), r)
	os.Exit(1)
}

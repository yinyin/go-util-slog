package utilslog

import (
	"context"
	"log/slog"
)

// ErrorContext logs an error message with context.
// The `err` parameter is REQUIRED, it will be included in the log attributes with "error" as its key.
// Use the `.ErrorContext()` in original slog package if you don't want to include an error.
func ErrorContext(ctx context.Context, msg string, err error, attrs ...slog.Attr) {
	logErrAttrs(ctx, slog.LevelError, msg, err, attrs...)
}

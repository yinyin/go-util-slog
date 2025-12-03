package utilslog

import (
	"context"
	"log/slog"
)

// WarnContext logs a warning message with context.
// The `err` parameter is REQUIRED, it will be included in the log attributes with "error" as its key.
// Use the `.WarnContext()` in original slog package if you don't want to include an error.
func WarnContext(ctx context.Context, msg string, err error, attrs ...slog.Attr) {
	logErrAttrs(ctx, slog.LevelWarn, msg, err, attrs...)
}

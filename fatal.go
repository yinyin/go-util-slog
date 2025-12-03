package utilslog

import (
	"context"
	"log/slog"
	"os"
)

// FatalContext logs a fatal error message with context and exits the program with `os.Exit(1)`.
// The `err` parameter is optional; if provided, it will be included in the log attributes with "error" as its key.
func FatalContext(ctx context.Context, msg string, err error, attrs ...slog.Attr) {
	logErrAttrs(ctx, slog.LevelError, msg, err, attrs...)
	os.Exit(1)
}

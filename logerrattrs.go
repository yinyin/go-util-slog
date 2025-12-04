package utilslog

import (
	"context"
	"log/slog"
	"runtime"
	"time"
)

func logErrAttrs(
	ctx context.Context,
	level slog.Level,
	msg string,
	err error,
	attrs ...slog.Attr) {
	logger := slog.Default()
	if ctx == nil {
		ctx = context.Background()
	}
	if !logger.Enabled(ctx, level) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:]) // skip [Callers, *Context, logErrAttrs]
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	if err != nil {
		attrs = append(attrs, slog.Any("error", err))
	}
	if len(attrs) > 0 {
		r.AddAttrs(attrs...)
	}
	_ = logger.Handler().Handle(ctx, r)
}

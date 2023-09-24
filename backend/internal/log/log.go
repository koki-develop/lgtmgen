package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

var (
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
)

func Warn(ctx context.Context, msg string, args ...interface{}) {
	logger.WarnContext(ctx, msg, args...)
}

func Info(ctx context.Context, msg string, args ...interface{}) {
	logger.InfoContext(ctx, msg, args...)
}

func Error(ctx context.Context, msg string, err error, args ...interface{}) {
	logger.ErrorContext(ctx, msg, append(args, "error", err, "trace", fmt.Sprintf("%+v", err))...)
}

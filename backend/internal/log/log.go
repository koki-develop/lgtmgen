package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
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

func Middleware(ctx *gin.Context) {
	p := ctx.Request.URL.Path
	q := ctx.Request.URL.RawQuery

	ctx.Next()

	ip := ctx.ClientIP()
	method := ctx.Request.Method
	statusCode := ctx.Writer.Status()

	Info(
		ctx, "gin request",
		"status_code", statusCode,
		"method", method,
		"path", p,
		"query", q,
		"ip", ip,
	)
}

package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/awslabs/aws-lambda-go-api-proxy/core"
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
	ip := ctx.ClientIP()

	apictx, ok := core.GetAPIGatewayContextFromContext(ctx.Request.Context())
	if ok {
		ip = apictx.Identity.SourceIP
	} else {
		Info(ctx, "no api gateway context")
	}

	p := ctx.Request.URL.Path
	q := ctx.Request.URL.RawQuery

	ctx.Next()

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

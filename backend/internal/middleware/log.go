package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Apply() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := ctx.Request.URL.Path
		q := ctx.Request.URL.RawQuery

		ctx.Next()

		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()

		log.Info(
			ctx, "gin request",
			"status_code", statusCode,
			"method", method,
			"path", p,
			"query", q,
			"ip", util.GetClientIPFromContext(ctx),
		)
	}
}

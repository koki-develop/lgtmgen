package util

import (
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
)

func GetClientIPFromContext(ctx *gin.Context) string {
	var ip string

	apictx, ok := core.GetAPIGatewayContextFromContext(ctx.Request.Context())
	if ok {
		log.Info(ctx, "api gateway context found")
		ip = apictx.Identity.SourceIP
	} else {
		log.Info(ctx, "no api gateway context")
		ip = ctx.ClientIP()
	}

	return ip
}

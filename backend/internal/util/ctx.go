package util

import (
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/gin-gonic/gin"
)

func GetClientIPFromContext(ctx *gin.Context) string {
	var ip string

	apictx, ok := core.GetAPIGatewayContextFromContext(ctx.Request.Context())
	if ok {
		ip = apictx.Identity.SourceIP
	} else {
		ip = ctx.ClientIP()
	}

	return ip
}

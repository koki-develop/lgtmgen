package api

import (
	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/service"
)

func NewEngine() *gin.Engine {
	r := gin.Default()
	svc := service.New()

	r.GET("/h", svc.HealthCheck)

	return r
}

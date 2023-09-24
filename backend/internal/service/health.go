package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthService struct{}

func newHealthService() *healthService {
	return &healthService{}
}

func (s *healthService) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]any{"status": "ok"})
}

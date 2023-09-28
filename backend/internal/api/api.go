package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/service"
)

func NewEngine(ctx context.Context) (*gin.Engine, error) {
	svc, err := service.New(ctx)
	if err != nil {
		return nil, err
	}

	e := gin.New()
	e.Use(log.Middleware)
	e.Use(gin.Recovery())

	e.GET("/h", svc.HealthCheck)

	v1 := e.Group("/v1")
	{
		v1.GET("/lgtms", svc.ListLGTMs)
		v1.POST("/lgtms", svc.CreateLGTM)

		v1.POST("/reports", svc.CreateReport)

		v1.GET("/images", svc.SearchImages)
	}

	return e, nil
}

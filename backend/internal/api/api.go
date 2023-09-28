package api

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/middleware"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
	"github.com/koki-develop/lgtmgen/backend/internal/service"
)

func NewEngine(ctx context.Context) (*gin.Engine, error) {
	r, err := repo.New(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize repository")
	}

	svc, err := service.New(ctx, r)
	if err != nil {
		return nil, err
	}

	e := gin.New()
	e.Use(log.Middleware)
	e.Use(gin.Recovery())
	rl := middleware.NewRateLimitter(r)

	e.GET("/h", svc.HealthCheck)

	v1 := e.Group("/v1")
	{
		v1.GET("/lgtms", svc.ListLGTMs)
		v1.POST("/lgtms", rl.Apply("post/lgtms", 5), svc.CreateLGTM) // TODO: 5 -> 100

		v1.POST("/reports", svc.CreateReport)

		v1.GET("/images", rl.Apply("get/images", 5), svc.SearchImages) // TODO: 5 -> 30
	}

	return e, nil
}

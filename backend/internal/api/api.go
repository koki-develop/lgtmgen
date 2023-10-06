package api

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
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
	e.Use(middleware.NewLogger().Apply())
	e.Use(gin.Recovery())
	e.Use(middleware.NewCORS().Apply(env.Vars.FrontendOrigin))
	rl := middleware.NewRateLimitter(r)

	e.GET("/h", svc.HealthCheck)

	v1 := e.Group("/v1")
	{
		v1.GET("/lgtms", svc.ListLGTMs)
		v1.POST("/lgtms", rl.Apply("post/lgtms", 100), svc.CreateLGTM)

		v1.POST("/reports", svc.CreateReport)

		v1.GET("/images", rl.Apply("get/images", 30), svc.SearchImages)

		v1.GET("/news", svc.ListNews)
	}

	return e, nil
}

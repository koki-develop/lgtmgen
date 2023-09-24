package api

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
	"github.com/koki-develop/lgtmgen/backend/internal/service"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
)

func NewEngine(ctx context.Context) (*gin.Engine, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load aws config")
	}

	dbOpts := []func(*dynamodb.Options){}
	storageOpts := []func(*s3.Options){}
	if env.Vars.Stage == "local" {
		dbOpts = append(dbOpts, func(o *dynamodb.Options) {
			o.BaseEndpoint = util.Ptr("http://localhost:4566")
		})
		storageOpts = append(storageOpts, func(o *s3.Options) {
			o.BaseEndpoint = util.Ptr("http://localhost:4566")
			o.UsePathStyle = true
		})
	}
	dbClient := dynamodb.NewFromConfig(cfg, dbOpts...)
	storageClient := s3.NewFromConfig(cfg, storageOpts...)

	r := repo.New(dbClient, storageClient)
	svc := service.New(r)
	e := gin.New()
	e.Use(log.Middleware)
	e.Use(gin.Recovery())

	e.GET("/h", svc.HealthCheck)

	v1 := e.Group("/v1")
	{
		v1.GET("/lgtms", svc.ListLGTMs)
		v1.POST("/lgtms", svc.CreateLGTM)
	}

	return e, nil
}

package api

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
	"github.com/koki-develop/lgtmgen/backend/internal/service"
)

func NewEngine(ctx context.Context) (*gin.Engine, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	dbClient := dynamodb.NewFromConfig(cfg)
	storageClient := s3.NewFromConfig(cfg)
	r := repo.New(dbClient, storageClient)
	svc := service.New(r)
	e := gin.Default()

	e.GET("/h", svc.HealthCheck)

	v1 := e.Group("/v1")
	{
		v1.GET("/lgtms", svc.ListLGTMs)
	}

	return e, nil
}

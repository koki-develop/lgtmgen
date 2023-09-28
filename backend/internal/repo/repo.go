package repo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
	"github.com/pkg/errors"
	"github.com/slack-go/slack"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

type Repository struct {
	*lgtmRepository
	*reportRepository
	*notificationsRepository
	*imageRepository
	*rateRepository
}

func New(ctx context.Context) (*Repository, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load aws config")
	}

	dbOpts := []func(*dynamodb.Options){}
	storageOpts := []func(*s3.Options){}
	queueOpts := []func(*sqs.Options){}
	if env.Vars.Stage == "local" {
		dbOpts = append(dbOpts, func(o *dynamodb.Options) {
			o.BaseEndpoint = util.Ptr("http://localhost:4566")
		})
		storageOpts = append(storageOpts, func(o *s3.Options) {
			o.BaseEndpoint = util.Ptr("http://localhost:4566")
			o.UsePathStyle = true
		})
		queueOpts = append(queueOpts, func(o *sqs.Options) {
			o.BaseEndpoint = util.Ptr("http://localhost:4566")
		})
	}
	dbClient := dynamodb.NewFromConfig(cfg, dbOpts...)
	storageClient := s3.NewFromConfig(cfg, storageOpts...)
	queueClient := sqs.NewFromConfig(cfg, queueOpts...)

	slackClient := slack.New(env.Vars.SlackOAuthToken)

	search, err := customsearch.NewService(ctx, option.WithAPIKey(env.Vars.GoogleAPIKey))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create search service")
	}

	return &Repository{
		lgtmRepository:          newLGTMRepository(dbClient, storageClient),
		reportRepository:        newReportRepository(dbClient, queueClient),
		imageRepository:         newImageRepository(env.Vars.SearchEngineID, search),
		rateRepository:          newRateRepository(dbClient),
		notificationsRepository: newNotificationsRepository(queueClient, slackClient),
	}, nil
}

package repo

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/slack-go/slack"
	"google.golang.org/api/customsearch/v1"
)

type Repository struct {
	*lgtmRepository
	*reportRepository
	*notificationsRepository
	*imageRepository
}

type Config struct {
	DBClient       *dynamodb.Client
	StorageClient  *s3.Client
	QueueClient    *sqs.Client
	SlackClient    *slack.Client
	SearchEngine   *customsearch.Service
	SearchEngineID string
}

func New(cfg *Config) *Repository {
	return &Repository{
		lgtmRepository:          newLGTMRepository(cfg.DBClient, cfg.StorageClient),
		reportRepository:        newReportRepository(cfg.DBClient, cfg.QueueClient),
		notificationsRepository: newNotificationsRepository(cfg.QueueClient, cfg.SlackClient),
		imageRepository:         newImageRepository(cfg.SearchEngineID, cfg.SearchEngine),
	}
}

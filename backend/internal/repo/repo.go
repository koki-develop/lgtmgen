package repo

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Repository struct {
	*lgtmRepository
	*reportRepository
	*notificationsRepository
}

func New(dbClient *dynamodb.Client, storageClient *s3.Client, queueClient *sqs.Client) *Repository {
	return &Repository{
		lgtmRepository:          newLGTMRepository(dbClient, storageClient),
		reportRepository:        newReportRepository(dbClient, queueClient),
		notificationsRepository: newNotificationsRepository(queueClient),
	}
}

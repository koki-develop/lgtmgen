package repo

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Repository struct {
	*lgtmRepository
}

func New(dbClient *dynamodb.Client, storageClient *s3.Client) *Repository {
	return &Repository{
		lgtmRepository: newLGTMRepository(dbClient, storageClient),
	}
}

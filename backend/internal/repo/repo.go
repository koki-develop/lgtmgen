package repo

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type Repository struct {
	*lgtmRepository
}

func New(dbClient *dynamodb.Client) *Repository {
	return &Repository{
		lgtmRepository: newLGTMRepository(dbClient),
	}
}

package repo

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type tagRepository struct {
	dbClient *dynamodb.Client
}

func newTagRepository(dbClient *dynamodb.Client) *tagRepository {
	return &tagRepository{
		dbClient: dbClient,
	}
}

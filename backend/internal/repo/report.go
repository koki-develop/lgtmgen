package repo

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type reportRepository struct {
	dbClient *dynamodb.Client
}

func newReportRepository(dbClient *dynamodb.Client) *reportRepository {
	return &reportRepository{dbClient: dbClient}
}

package repo

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
)

type reportRepository struct {
	dbClient    *dynamodb.Client
	queueClient *sqs.Client
}

func newReportRepository(dbClient *dynamodb.Client, queueClient *sqs.Client) *reportRepository {
	return &reportRepository{
		dbClient:    dbClient,
		queueClient: queueClient,
	}
}

func (r *reportRepository) CreateReport(ctx context.Context, lgtmID string, t models.ReportType, text string) (*models.Report, error) {
	rp := &models.Report{
		ID:        models.NewID(),
		LGTMID:    lgtmID,
		Type:      t,
		Text:      text,
		CreatedAt: time.Now(),
	}

	item, err := attributevalue.MarshalMap(rp)
	if err != nil {
		return nil, err
	}

	_, err = r.dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: util.Ptr(env.Vars.DynamoDBTableReports),
		Item:      item,
	})
	if err != nil {
		return nil, err
	}

	return rp, nil
}

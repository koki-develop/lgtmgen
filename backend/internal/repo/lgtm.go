package repo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
)

type lgtmRepository struct {
	dbClient      *dynamodb.Client
	storageClient *s3.Client
}

func newLGTMRepository(db *dynamodb.Client, storageClient *s3.Client) *lgtmRepository {
	return &lgtmRepository{
		dbClient:      db,
		storageClient: storageClient,
	}
}

type lgtmListOptions struct {
	Limit int
}

type LGTMListOption func(*lgtmListOptions)

func WithLGTMLimit(limit int) LGTMListOption {
	return func(o *lgtmListOptions) {
		o.Limit = limit
	}
}

func (r *lgtmRepository) ListLGTMs(ctx context.Context, opts ...LGTMListOption) (models.LGTMs, error) {
	o := &lgtmListOptions{}
	for _, opt := range opts {
		opt(o)
	}

	expr, err := expression.NewBuilder().
		WithKeyCondition(expression.KeyEqual(expression.Key("status"), expression.Value("ok"))).
		Build()
	if err != nil {
		return nil, err
	}

	resp, err := r.dbClient.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName:                 util.Ptr(env.Vars.DynamoDBTableLGTMs),
			IndexName:                 util.Ptr("index_by_status"),
			KeyConditionExpression:    expr.KeyCondition(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),

			Limit:            util.Ptr(int32(o.Limit)),
			ScanIndexForward: util.Ptr(false),
		},
	)
	if err != nil {
		return nil, err
	}

	lgtms := make(models.LGTMs, len(resp.Items))
	for _, item := range resp.Items {
		var lgtm models.LGTM
		err := attributevalue.UnmarshalMap(item, &lgtm)
		if err != nil {
			return nil, err
		}
		lgtms = append(lgtms, &lgtm)
	}

	return lgtms, nil
}

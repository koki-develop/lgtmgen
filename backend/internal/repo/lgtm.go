package repo

import (
	"bytes"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/cockroachdb/errors"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/lgtmgen"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
)

type lgtmRepository struct {
	dbClient      *dynamodb.Client
	storageClient *s3.Client
	queueClient   *sqs.Client
}

func newLGTMRepository(db *dynamodb.Client, storageClient *s3.Client, queueClient *sqs.Client) *lgtmRepository {
	return &lgtmRepository{
		dbClient:      db,
		storageClient: storageClient,
		queueClient:   queueClient,
	}
}

func (r *lgtmRepository) FindLGTM(ctx context.Context, id string) (*models.LGTM, error) {
	expr, err := expression.NewBuilder().
		WithKeyCondition(expression.KeyEqual(expression.Key("id"), expression.Value(id))).
		Build()
	resp, err := r.dbClient.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName:                 util.Ptr(env.Vars.DynamoDBTableLGTMs),
			KeyConditionExpression:    expr.KeyCondition(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query lgtm")
	}

	if len(resp.Items) == 0 {
		return nil, nil
	}

	l := &models.LGTM{}
	if err := attributevalue.UnmarshalMap(resp.Items[0], l); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal lgtm")
	}

	return l, nil
}

type lgtmListOptions struct {
	Limit int
	After *models.LGTM
}

type LGTMListOption func(*lgtmListOptions)

func WithLGTMLimit(limit int) LGTMListOption {
	return func(o *lgtmListOptions) {
		o.Limit = limit
	}
}

func WithLGTMAfter(l *models.LGTM) LGTMListOption {
	return func(o *lgtmListOptions) {
		o.After = l
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
		return nil, errors.Wrap(err, "failed to build expression")
	}

	var startKey map[string]types.AttributeValue
	if o.After != nil {
		startKey, err = attributevalue.MarshalMap(o.After)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal")
		}
	}

	resp, err := r.dbClient.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName:                 util.Ptr(env.Vars.DynamoDBTableLGTMs),
			IndexName:                 util.Ptr("index_by_status"),
			KeyConditionExpression:    expr.KeyCondition(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			Limit:                     util.Ptr(int32(o.Limit)),
			ScanIndexForward:          util.Ptr(false),
			ExclusiveStartKey:         startKey,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query")
	}

	lgtms := make(models.LGTMs, len(resp.Items))
	for i, item := range resp.Items {
		var lgtm models.LGTM
		if err := attributevalue.UnmarshalMap(item, &lgtm); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal")
		}
		lgtms[i] = &lgtm
	}

	return lgtms, nil
}

func (r *lgtmRepository) CreateLGTM(ctx context.Context, data []byte) (*models.LGTM, error) {
	t := http.DetectContentType(data)
	log.Info(ctx, "detected content type", "type", t)

	if !strings.HasPrefix(t, "image/") {
		return nil, errors.Wrap(lgtmgen.ErrUnsupportImageFormat, "not image")
	}

	img, err := lgtmgen.Generate(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate lgtm")
	}

	lgtm := &models.LGTM{
		ID:        models.NewID(),
		Status:    models.LGTMStatusPending,
		CreatedAt: time.Now(),
	}

	item, err := attributevalue.MarshalMap(lgtm)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal")
	}

	_, err = r.dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: util.Ptr(env.Vars.DynamoDBTableLGTMs),
		Item:      item,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to put item")
	}

	uploader := manager.NewUploader(r.storageClient)
	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      util.Ptr(env.Vars.S3BucketImages),
		Key:         util.Ptr(lgtm.ID),
		Body:        bytes.NewReader(img),
		ContentType: util.Ptr(t),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to upload image")
	}

	k, err := attributevalue.MarshalMap(map[string]interface{}{"id": lgtm.ID, "created_at": lgtm.CreatedAt})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal")
	}
	expr, err := expression.NewBuilder().
		WithUpdate(expression.Set(expression.Name("status"), expression.Value(models.LGTMStatusOK))).
		Build()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build expression")
	}

	_, err = r.dbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 util.Ptr(env.Vars.DynamoDBTableLGTMs),
		Key:                       k,
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to update item")
	}

	lgtm.Status = models.LGTMStatusOK
	return lgtm, nil
}

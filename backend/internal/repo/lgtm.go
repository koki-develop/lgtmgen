package repo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.1/computervision"
	_ "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.1/computervision"
	"github.com/Azure/go-autorest/autorest"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
}

func newLGTMRepository(db *dynamodb.Client, storageClient *s3.Client) *lgtmRepository {
	return &lgtmRepository{
		dbClient:      db,
		storageClient: storageClient,
	}
}

func (r *lgtmRepository) FindLGTM(ctx context.Context, id string) (*models.LGTM, error) {
	expr, err := expression.NewBuilder().
		WithKeyCondition(expression.KeyEqual(expression.Key("id"), expression.Value(id))).
		Build()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build expression")
	}

	resp, err := r.dbClient.Query(
		ctx,
		&dynamodb.QueryInput{
			TableName:                 util.Ptr(r.table()),
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
	Tag    string
	Limit  int
	After  *models.LGTM
	Random bool
}

type LGTMListOption func(*lgtmListOptions)

func WithLGTMTag(tag string) LGTMListOption {
	return func(o *lgtmListOptions) {
		o.Tag = tag
	}
}

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

func WithLGTMRandom() LGTMListOption {
	return func(o *lgtmListOptions) {
		o.Random = true
	}
}

func (r *lgtmRepository) ListLGTMs(ctx context.Context, opts ...LGTMListOption) (models.LGTMs, error) {
	o := &lgtmListOptions{}
	for _, opt := range opts {
		opt(o)
	}

	if o.Random {
		return r.listLGTMsRandomly(ctx, o)
	} else {
		return r.listLGTMs(ctx, o)
	}
}

func (r *lgtmRepository) listLGTMs(ctx context.Context, o *lgtmListOptions) (models.LGTMs, error) {
	exb := expression.NewBuilder().
		WithKeyCondition(expression.KeyEqual(expression.Key("status"), expression.Value("ok")))

	if o.Tag != "" {
		exb = exb.WithFilter(expression.Contains(expression.Name("tags_ja"), o.Tag).Or(expression.Contains(expression.Name("tags_en"), o.Tag)))
	}

	expr, err := exb.Build()
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
			TableName:                 util.Ptr(r.table()),
			IndexName:                 util.Ptr("index_by_status"),
			FilterExpression:          expr.Filter(),
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

func (r *lgtmRepository) listLGTMsRandomly(ctx context.Context, o *lgtmListOptions) (models.LGTMs, error) {
	resp, err := r.storageClient.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:  util.Ptr(r.bucket()),
		MaxKeys: 500,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list objects")
	}
	rand.Shuffle(len(resp.Contents), func(i, j int) {
		resp.Contents[i], resp.Contents[j] = resp.Contents[j], resp.Contents[i]
	})
	limit := min(o.Limit, len(resp.Contents))

	keys := make([]string, limit)
	for i, obj := range resp.Contents[:limit] {
		keys[i] = *obj.Key
	}

	lgtms := make(models.LGTMs, limit)
	for i, key := range keys {
		lgtms[i] = &models.LGTM{ID: key}
	}

	return lgtms, nil
}

func (r *lgtmRepository) CreateLGTM(ctx context.Context, data []byte, opts ...lgtmgen.GenerateOption) (*models.LGTM, error) {
	t := http.DetectContentType(data)
	log.Info(ctx, "detected content type", "type", t)

	if !strings.HasPrefix(t, "image/") {
		return nil, errors.Wrap(lgtmgen.ErrUnsupportImageFormat, "not image")
	}

	img, err := lgtmgen.Generate(data, opts...)
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
		TableName: util.Ptr(r.table()),
		Item:      item,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to put item")
	}

	uploader := manager.NewUploader(r.storageClient)
	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      util.Ptr(r.bucket()),
		Key:         util.Ptr(lgtm.ID),
		Body:        bytes.NewReader(img),
		ContentType: util.Ptr(t),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to upload image")
	}

	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      util.Ptr(r.originalBucket()),
		Key:         util.Ptr(lgtm.ID),
		Body:        bytes.NewReader(data),
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
		TableName:                 util.Ptr(r.table()),
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

func (r *lgtmRepository) DeleteLGTM(ctx context.Context, id string) error {
	lgtm, err := r.FindLGTM(ctx, id)
	if err != nil {
		return errors.Wrap(err, "failed to find lgtm")
	}

	k, err := attributevalue.MarshalMap(map[string]interface{}{"id": lgtm.ID, "created_at": lgtm.CreatedAt})
	if err != nil {
		return errors.Wrap(err, "failed to marshal")
	}

	_, err = r.dbClient.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: util.Ptr(r.table()),
		Key:       k,
	})
	if err != nil {
		return errors.Wrap(err, "failed to delete item")
	}

	_, err = r.storageClient.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: util.Ptr(r.bucket()),
		Key:    util.Ptr(lgtm.ID),
	})
	if err != nil {
		return errors.Wrap(err, "failed to delete object")
	}

	return nil
}

func (r *lgtmRepository) TagLGTM(ctx context.Context) (*models.LGTM, error) {
	resp, err := r.storageClient.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:  util.Ptr(r.originalBucket()),
		MaxKeys: 1,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list objects")
	}
	if len(resp.Contents) == 0 {
		return nil, nil
	}

	id := resp.Contents[0].Key

	lgtm, err := r.FindLGTM(ctx, *id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find lgtm")
	}

	if lgtm == nil {
		_, err = r.storageClient.DeleteObject(ctx, &s3.DeleteObjectInput{
			Bucket: util.Ptr(r.originalBucket()),
			Key:    id,
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to delete object")
		}
		return nil, nil
	}

	obj, err := r.storageClient.GetObject(ctx, &s3.GetObjectInput{
		Bucket: util.Ptr(r.originalBucket()),
		Key:    id,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to download image")
	}
	defer obj.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(obj.Body)

	k, err := attributevalue.MarshalMap(map[string]interface{}{"id": lgtm.ID, "created_at": lgtm.CreatedAt})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal")
	}

	cv := computervision.New(env.Vars.AzureEndpoint)
	cv.Authorizer = autorest.NewCognitiveServicesAuthorizer(env.Vars.AzureAPIKey)
	for _, lang := range []string{"ja", "en"} {
		b := io.NopCloser(bytes.NewReader(buf.Bytes()))
		defer b.Close()

		rslt, err := cv.TagImageInStream(ctx, b, lang)
		if err != nil {
			return nil, errors.Wrap(err, "failed to tag image")
		}
		defer rslt.Body.Close()

		var tags []string
		for _, tag := range *rslt.Tags {
			if *tag.Confidence > 0.90 {
				tags = append(tags, *tag.Name)
			}
		}

		expr, err := expression.NewBuilder().
			WithUpdate(expression.Set(expression.Name(fmt.Sprintf("tags_%s", lang)), expression.Value(tags))).
			Build()
		if err != nil {
			return nil, errors.Wrap(err, "failed to build expression")
		}

		_, err = r.dbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
			TableName:                 util.Ptr(r.table()),
			Key:                       k,
			UpdateExpression:          expr.Update(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to update item")
		}
	}

	_, err = r.storageClient.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: util.Ptr(r.originalBucket()),
		Key:    id,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete object")
	}

	return lgtm, nil
}

func (*lgtmRepository) table() string {
	return fmt.Sprintf("lgtmgen-%s-lgtms", env.Vars.Stage)
}

func (*lgtmRepository) bucket() string {
	return fmt.Sprintf("lgtmgen-%s-images", env.Vars.Stage)
}

func (*lgtmRepository) originalBucket() string {
	return fmt.Sprintf("lgtmgen-%s-original-images", env.Vars.Stage)
}

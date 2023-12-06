package repo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/cockroachdb/errors"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
)

type tagRepository struct {
	dbClient *dynamodb.Client
}

func newTagRepository(dbClient *dynamodb.Client) *tagRepository {
	return &tagRepository{
		dbClient: dbClient,
	}
}

func (r *tagRepository) ListTags(ctx context.Context, lang string) (models.Tags, error) {
	expr, err := expression.NewBuilder().
		WithKeyCondition(expression.KeyEqual(expression.Key("lang"), expression.Value(lang))).
		Build()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build expression")
	}

	res, err := r.dbClient.Query(ctx, &dynamodb.QueryInput{
		TableName:                 util.Ptr(r.table()),
		IndexName:                 util.Ptr("index_by_lang"),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		Limit:                     util.Ptr(int32(20)),
		ScanIndexForward:          util.Ptr(false),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to query tags")
	}

	var tags models.Tags
	if err := attributevalue.UnmarshalListOfMaps(res.Items, &tags); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal tags")
	}

	return tags, nil
}

func (r *tagRepository) IncrementTagByName(ctx context.Context, name string, lang string) error {
	expr, err := expression.NewBuilder().
		WithUpdate(expression.Add(expression.Name("count"), expression.Value(1))).
		Build()
	if err != nil {
		return errors.Wrap(err, "failed to build expression")
	}

	k, err := attributevalue.MarshalMap(map[string]interface{}{"name": name, "lang": lang})
	if err != nil {
		return errors.Wrap(err, "failed to marshal key")
	}

	_, err = r.dbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 util.Ptr(r.table()),
		Key:                       k,
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to update tag")
	}

	return nil
}

func (*tagRepository) table() string {
	return fmt.Sprintf("lgtmgen-%s-tags", env.Vars.Stage)
}

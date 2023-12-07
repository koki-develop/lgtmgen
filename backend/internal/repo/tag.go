package repo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cockroachdb/errors"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
)

type categoryRepository struct {
	dbClient *dynamodb.Client
}

func newCategoryRepository(dbClient *dynamodb.Client) *categoryRepository {
	return &categoryRepository{
		dbClient: dbClient,
	}
}

func (r *categoryRepository) ListCategories(ctx context.Context, lang string) (models.Categories, error) {
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
		return nil, errors.Wrap(err, "failed to query categories")
	}

	var cs models.Categories
	if err := attributevalue.UnmarshalListOfMaps(res.Items, &cs); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal categories")
	}

	return cs, nil
}

func (r *categoryRepository) IncrementCategoryByName(ctx context.Context, name string, lang string) (*models.Category, error) {
	expr, err := expression.NewBuilder().
		WithUpdate(expression.Add(expression.Name("count"), expression.Value(1))).
		Build()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build expression")
	}

	k, err := attributevalue.MarshalMap(map[string]interface{}{"name": name, "lang": lang})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal key")
	}

	resp, err := r.dbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 util.Ptr(r.table()),
		Key:                       k,
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ReturnValues:              types.ReturnValueAllNew,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to update category")
	}

	var c models.Category
	if err := attributevalue.UnmarshalMap(resp.Attributes, &c); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal category")
	}

	return &c, nil
}

func (*categoryRepository) table() string {
	return fmt.Sprintf("lgtmgen-%s-categories", env.Vars.Stage)
}

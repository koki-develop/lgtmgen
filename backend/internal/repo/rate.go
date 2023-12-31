package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/cockroachdb/errors"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
)

type rateRepository struct {
	dbClient *dynamodb.Client
}

func newRateRepository(dbClient *dynamodb.Client) *rateRepository {
	return &rateRepository{dbClient: dbClient}
}

func (r *rateRepository) FindRate(ctx context.Context, ip, tier string) (*models.Rate, error) {
	key, err := attributevalue.MarshalMap(map[string]string{
		"ip":   ip,
		"tier": tier,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal map")
	}

	res, err := r.dbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: util.Ptr(r.table()),
		Key:       key,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get item")
	}

	if res.Item == nil {
		return nil, nil
	}

	var rate models.Rate
	if err := attributevalue.UnmarshalMap(res.Item, &rate); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal map")
	}

	if rate.Expired() {
		return nil, nil
	}

	return &rate, nil
}

func (r *rateRepository) IncrementRate(ctx context.Context, ip, tier string) error {
	rate, err := r.FindRate(ctx, ip, tier)
	if err != nil {
		return errors.Wrap(err, "failed to find rate")
	}

	if rate == nil {
		rate = &models.Rate{
			IP:      ip,
			Tier:    tier,
			Count:   1,
			ResetAt: time.Now().Add(1 * time.Hour),
		}
		item, err := attributevalue.MarshalMap(rate)
		if err != nil {
			return errors.Wrap(err, "failed to marshal map")
		}
		_, err = r.dbClient.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: util.Ptr(r.table()),
			Item:      item,
		})
		if err != nil {
			return errors.Wrap(err, "failed to put item")
		}
	} else {
		expr, err := expression.NewBuilder().
			WithUpdate(expression.Add(expression.Name("count"), expression.Value(1))).
			Build()
		if err != nil {
			return errors.Wrap(err, "failed to build expression")
		}

		key, err := attributevalue.MarshalMap(map[string]string{
			"ip":   ip,
			"tier": tier,
		})
		if err != nil {
			return errors.Wrap(err, "failed to marshal map")
		}

		_, err = r.dbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
			TableName:                 util.Ptr(r.table()),
			Key:                       key,
			UpdateExpression:          expr.Update(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
		})
		if err != nil {
			return errors.Wrap(err, "failed to update item")
		}
	}

	return nil
}

func (*rateRepository) table() string {
	return fmt.Sprintf("lgtmgen-%s-rates", env.Vars.Stage)
}

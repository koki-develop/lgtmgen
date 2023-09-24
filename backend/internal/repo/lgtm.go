package repo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
)

type lgtmRepository struct {
	dbClient *dynamodb.Client
}

func newLGTMRepository(db *dynamodb.Client) *lgtmRepository {
	return &lgtmRepository{
		dbClient: db,
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

	// FIXME: implement
	resp, err := r.dbClient.ListTables(
		ctx,
		&dynamodb.ListTablesInput{
			Limit: util.Ptr(int32(o.Limit)),
		},
	)
	if err != nil {
		return nil, err
	}
	lgtms := make(models.LGTMs, len(resp.TableNames))
	for i, tableName := range resp.TableNames {
		lgtms[i] = &models.LGTM{ID: tableName}
	}
	return lgtms, nil
}

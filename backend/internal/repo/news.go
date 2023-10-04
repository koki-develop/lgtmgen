package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
	"gopkg.in/yaml.v3"
)

type newsRepository struct {
	storageClient *s3.Client
}

func newNewsRepository(storageClient *s3.Client) *newsRepository {
	return &newsRepository{
		storageClient: storageClient,
	}
}

func (r *newsRepository) ListNews(ctx context.Context, locale string) (models.NewsList, error) {
	resp, err := r.storageClient.GetObject(ctx, &s3.GetObjectInput{
		Bucket: util.Ptr(r.bucket()),
		Key:    util.Ptr(fmt.Sprintf("news.%s.yml", locale)),
	})
	if err != nil {
		var aerr *types.NoSuchKey
		if errors.As(err, &aerr) {
			return models.NewsList{}, nil
		}

		return nil, fmt.Errorf("failed to get object: %w", err)
	}
	defer resp.Body.Close()

	var news models.NewsList
	if err := yaml.NewDecoder(resp.Body).Decode(&news); err != nil {
		return nil, fmt.Errorf("failed to decode yaml: %w", err)
	}

	return news, nil
}

func (r *newsRepository) bucket() string {
	return fmt.Sprintf("lgtmgen-%s-news", env.Vars.Stage)
}

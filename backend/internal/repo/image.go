package repo

import (
	"context"
	"net/url"

	"github.com/cockroachdb/errors"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
	"google.golang.org/api/customsearch/v1"
)

type imageRepository struct {
	searchEngineID string
	searchEngine   *customsearch.Service
}

func newImageRepository(engineID string, searchEngine *customsearch.Service) *imageRepository {
	return &imageRepository{
		searchEngineID: engineID,
		searchEngine:   searchEngine,
	}
}

func (r *imageRepository) SearchImages(ctx context.Context, q string) (models.Images, error) {
	call := r.searchEngine.Cse.List().
		Cx(r.searchEngineID).
		Q(q).
		SearchType("image").
		Safe("active").
		Num(10).
		Start(1)

	resp, err := call.Do()
	if err != nil {
		return nil, errors.Wrap(err, "failed to search")
	}

	imgs := models.Images{}
	for _, item := range resp.Items {
		u, err := url.ParseRequestURI(item.Link)
		if err != nil {
			continue
		}
		if u.Scheme != "https" {
			continue
		}

		exclude := []string{"image/svg+xml"}
		if util.Contains(exclude, item.Mime) {
			continue
		}

		imgs = append(imgs, &models.Image{
			URL:   item.Link,
			Title: item.Title,
		})
	}

	return imgs, nil
}

package service

import (
	"context"

	"github.com/koki-develop/lgtmgen/backend/internal/repo"
)

type Service struct {
	*lgtmService
	*categoryService
	*reportService
	*notificationService
	*imageService
	*newsService
	*healthService
}

func New(ctx context.Context, repo *repo.Repository) (*Service, error) {
	return &Service{
		lgtmService:         newLGTMService(repo),
		categoryService:     newCategoryService(repo),
		reportService:       newReportService(repo),
		notificationService: newNotificationService(repo),
		imageService:        newImageService(repo),
		newsService:         newNewsService(repo),
		healthService:       newHealthService(),
	}, nil
}

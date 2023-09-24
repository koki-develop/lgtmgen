package service

import "github.com/koki-develop/lgtmgen/backend/internal/repo"

type Service struct {
	*lgtmService
	*healthService
}

func New(repo *repo.Repository) *Service {
	return &Service{
		lgtmService:   newLGTMService(repo),
		healthService: newHealthService(),
	}
}

package service

import "github.com/koki-develop/lgtmgen/backend/internal/repo"

type reportService struct {
	repo *repo.Repository
}

func newReportService(repo *repo.Repository) *reportService {
	return &reportService{repo: repo}
}

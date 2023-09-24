package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
)

type lgtmService struct {
	repo *repo.Repository
}

func newLGTMService(repo *repo.Repository) *lgtmService {
	return &lgtmService{
		repo: repo,
	}
}

func (svc *lgtmService) ListLGTMs(ctx *gin.Context) {
	lgtms, err := svc.repo.ListLGTMs(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lgtms)
}

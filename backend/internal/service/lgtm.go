package service

import (
	"net/http"
	"strconv"

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
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	lgtms, err := svc.repo.ListLGTMs(ctx, repo.WithLGTMLimit(limit))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lgtms)
}

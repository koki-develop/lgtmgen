package service

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
)

type imageService struct {
	repo *repo.Repository
}

func newImageService(repo *repo.Repository) *imageService {
	return &imageService{
		repo: repo,
	}
}

func (s *imageService) SearchImages(ctx *gin.Context) {
	q := ctx.Query("q")
	if strings.TrimSpace(q) == "" {
		log.Info(ctx, "query is empty")
		ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeBadRequest})
		return
	}

	imgs, err := s.repo.SearchImages(ctx, q)
	if err != nil {
		log.Error(ctx, "failed to search images", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": ErrCodeInternalServerError})
		return
	}

	ctx.JSON(http.StatusOK, imgs)
}

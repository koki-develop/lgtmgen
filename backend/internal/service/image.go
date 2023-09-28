package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
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
	imgs := models.Images{}

	ctx.JSON(http.StatusOK, imgs)
}

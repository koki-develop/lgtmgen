package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
)

type imageService struct{}

func newImageService() *imageService {
	return &imageService{}
}

func (s *imageService) SearchImages(ctx *gin.Context) {
	imgs := models.Images{}

	ctx.JSON(http.StatusOK, imgs)
}

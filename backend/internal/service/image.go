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

//	@Router		/v1/images [get]
//	@Param		q	query		string	true	"query"
//	@Success	200	{array}		models.Images
//	@Failure	400	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
func (s *imageService) SearchImages(ctx *gin.Context) {
	q := ctx.Query("q")
	if strings.TrimSpace(q) == "" {
		log.Info(ctx, "query is empty")
		renderError(ctx, http.StatusBadRequest, ErrCodeBadRequest)
		return
	}

	imgs, err := s.repo.SearchImages(ctx, q)
	if err != nil {
		log.Error(ctx, "failed to search images", err)
		renderError(ctx, http.StatusInternalServerError, ErrCodeInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, imgs)
}

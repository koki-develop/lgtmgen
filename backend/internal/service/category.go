package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
)

type categoryService struct {
	repo *repo.Repository
}

func newCategoryService(repo *repo.Repository) *categoryService {
	return &categoryService{
		repo: repo,
	}
}

// @Router		/v1/categories [get]
// @Param		lang	query		string	false	"lang"
// @Success	200		{array}		models.Category
// @Failure	400		{object}	ErrorResponse
// @Failure	500		{object}	ErrorResponse
func (svc *categoryService) ListCategories(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", "ja")

	cs, err := svc.repo.ListCategories(ctx, lang)
	if err != nil {
		log.Error(ctx, "failed to list categories", err)
		renderError(ctx, http.StatusInternalServerError, ErrCodeInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, cs)
}

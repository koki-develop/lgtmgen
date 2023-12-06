package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
)

type tagService struct {
	repo *repo.Repository
}

func newTagService(repo *repo.Repository) *tagService {
	return &tagService{
		repo: repo,
	}
}

//	@Router		/v1/tags [get]
//	@Param		lang	query		string	false	"lang"
//	@Success	200		{array}		models.Tag
//	@Failure	400		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
func (svc *tagService) ListTags(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", "ja")
	tags, err := svc.repo.ListTags(ctx, lang)
	if err != nil {
		log.Error(ctx, "failed to list tags", err)
		renderError(ctx, http.StatusInternalServerError, ErrCodeInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, tags)
}

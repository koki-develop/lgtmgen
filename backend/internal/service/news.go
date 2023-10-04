package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
)

type newsService struct {
	repo *repo.Repository
}

func newNewsService(repo *repo.Repository) *newsService {
	return &newsService{
		repo: repo,
	}
}

//		@Router		/v1/news [get]
//	 @Param locale query string false "locale"
//		@Success	200	{array}		models.News
//		@Failure	500	{object}	ErrorResponse
func (svc *newsService) ListNews(ctx *gin.Context) {
	locale := ctx.DefaultQuery("locale", "ja")
	l, err := svc.repo.ListNews(ctx, locale)
	if err != nil {
		log.Error(ctx, "failed to list news", err, "locale", locale)
		renderError(ctx, http.StatusInternalServerError, ErrCodeInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, l)
}

package service

import (
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/lgtmgen"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
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
	qlimit := ctx.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(qlimit)
	if err != nil {
		log.Info(ctx, "failed to parse limit", "limit", qlimit, "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeBadRequest})
		return
	}
	if limit < 1 || limit > 100 {
		log.Info(ctx, "limit is out of range", "limit", limit)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeBadRequest})
		return
	}

	lgtms, err := svc.repo.ListLGTMs(ctx, repo.WithLGTMLimit(limit))
	if err != nil {
		log.Error(ctx, "failed to list lgtms", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": ErrCodeInternalServerError})
		return
	}

	ctx.JSON(http.StatusOK, lgtms)
}

type CreateLGTMInput struct {
	Base64 string `json:"base64"`
}

func (ipt *CreateLGTMInput) Validate() error {
	if ipt.Base64 == "" {
		return errors.New("base64 is required")
	}

	return nil
}

func (svc *lgtmService) CreateLGTM(ctx *gin.Context) {
	var ipt CreateLGTMInput
	if err := ctx.ShouldBindJSON(&ipt); err != nil {
		log.Info(ctx, "failed to bind json", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeBadRequest})
		return
	}

	if err := ipt.Validate(); err != nil {
		log.Info(ctx, "failed to validate input", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeBadRequest})
		return
	}

	data, err := base64.StdEncoding.DecodeString(ipt.Base64)
	if err != nil {
		log.Info(ctx, "failed to decode base64", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeBadRequest})
		return
	}

	lgtm, err := svc.repo.Create(ctx, data)
	if err != nil {
		if errors.Is(err, lgtmgen.ErrUnsupportImageFormat) {
			log.Info(ctx, "unsupported image format", "error", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeUnsupportedImageFormat})
			return
		}

		log.Error(ctx, "failed to create lgtm", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": ErrCodeInternalServerError})
		return
	}

	ctx.JSON(http.StatusOK, lgtm)
}

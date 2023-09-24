package service

import (
	"encoding/base64"
	"errors"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ipt.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := base64.StdEncoding.DecodeString(ipt.Base64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	lgtm, err := svc.repo.Create(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lgtm)
}

package service

import (
	"bytes"
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
	repo       *repo.Repository
	httpClient *http.Client
}

func newLGTMService(repo *repo.Repository) *lgtmService {
	return &lgtmService{
		repo:       repo,
		httpClient: &http.Client{},
	}
}

func (svc *lgtmService) ListLGTMs(ctx *gin.Context) {
	opts := []repo.LGTMListOption{}

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
	opts = append(opts, repo.WithLGTMLimit(limit))

	after := ctx.Query("after")
	if after != "" {
		lgtm, err := svc.repo.FindLGTM(ctx, after)
		if err != nil {
			log.Error(ctx, "failed to find lgtm", err, "id", after)
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": ErrCodeInternalServerError})
			return
		}
		if lgtm == nil {
			log.Info(ctx, "lgtm not found", "id", after)
			ctx.JSON(http.StatusNotFound, gin.H{"code": ErrCodeNotFound})
			return
		}
		opts = append(opts, repo.WithLGTMAfter(lgtm))
	}

	lgtms, err := svc.repo.ListLGTMs(ctx, opts...)
	if err != nil {
		log.Error(ctx, "failed to list lgtms", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": ErrCodeInternalServerError})
		return
	}

	ctx.JSON(http.StatusOK, lgtms)
}

type CreateLGTMInput struct {
	URL    string `json:"url"`
	Base64 string `json:"base64"`
}

func (ipt *CreateLGTMInput) Validate() error {
	if ipt.URL == "" && ipt.Base64 == "" {
		return errors.New("url or base64 is required")
	}

	if ipt.URL != "" && ipt.Base64 != "" {
		return errors.New("url and base64 are exclusive")
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

	var data []byte

	if ipt.Base64 != "" {
		d, err := base64.StdEncoding.DecodeString(ipt.Base64)
		if err != nil {
			log.Info(ctx, "failed to decode base64", "error", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeBadRequest})
			return
		}
		data = d
	}

	if ipt.URL != "" {
		log.Info(ctx, "request", "url", ipt.URL)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, ipt.URL, nil)
		if err != nil {
			log.Info(ctx, "failed to create request", "error", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeFailedToGetImage})
			return
		}

		resp, err := svc.httpClient.Do(req)
		if err != nil {
			log.Info(ctx, "failed to get image", "error", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeFailedToGetImage})
			return
		}
		defer resp.Body.Close()
		log.Info(ctx, "response", "status", resp.StatusCode)

		if resp.StatusCode != http.StatusOK {
			log.Info(ctx, "failed to get image", "status", resp.StatusCode)
			ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeFailedToGetImage})
			return
		}

		buf := new(bytes.Buffer)
		if _, err = buf.ReadFrom(resp.Body); err != nil {
			log.Info(ctx, "failed to read image", "error", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeFailedToGetImage})
			return
		}
		data = buf.Bytes()
	}

	lgtm, err := svc.repo.CreateLGTM(ctx, data)
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

	ctx.JSON(http.StatusCreated, lgtm)
}

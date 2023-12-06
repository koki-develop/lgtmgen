package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/lgtmgen"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
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

// @Router		/v1/lgtms [get]
// @Param		limit	query		int		false	"limit"
// @Param		after	query		string	false	"after"
// @Param		random	query		bool	false	"random"
// @Param		tag		query		string	false	"tag"
// @Success	200		{array}		models.LGTM
// @Failure	400		{object}	ErrorResponse
// @Failure	500		{object}	ErrorResponse
func (svc *lgtmService) ListLGTMs(ctx *gin.Context) {
	opts := []repo.LGTMListOption{}

	if tag := ctx.Query("tag"); tag != "" {
		opts = append(opts, repo.WithLGTMTag(tag))
	}

	qlimit := ctx.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(qlimit)
	if err != nil {
		log.Info(ctx, "failed to parse limit", "limit", qlimit, "error", err)
		renderError(ctx, http.StatusBadRequest, ErrCodeBadRequest)
		return
	}
	if limit < 1 || limit > 100 {
		log.Info(ctx, "limit is out of range", "limit", limit)
		renderError(ctx, http.StatusBadRequest, ErrCodeBadRequest)
		return
	}
	opts = append(opts, repo.WithLGTMLimit(limit))

	after := ctx.Query("after")
	if after != "" {
		lgtm, err := svc.repo.FindLGTM(ctx, after)
		if err != nil {
			log.Error(ctx, "failed to find lgtm", err, "id", after)
			renderError(ctx, http.StatusInternalServerError, ErrCodeInternalServerError)
			return
		}
		if lgtm == nil {
			log.Info(ctx, "lgtm not found", "id", after)
			renderError(ctx, http.StatusBadRequest, ErrCodeBadRequest)
			return
		}
		opts = append(opts, repo.WithLGTMAfter(lgtm))
	}

	random := ctx.Query("random")
	if random == "true" {
		opts = append(opts, repo.WithLGTMRandom())
	}

	lgtms, err := svc.repo.ListLGTMs(ctx, opts...)
	if err != nil {
		log.Error(ctx, "failed to list lgtms", err)
		renderError(ctx, http.StatusInternalServerError, ErrCodeInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, lgtms)
}

type createLGTMInput struct {
	URL     string             `json:"url"`
	Base64  string             `json:"base64"`
	Options *createLGTMOptions `json:"options"`
}

type createLGTMOptions struct {
	TextColor string `json:"textColor"`
}

func (ipt *createLGTMInput) Validate() error {
	if ipt.URL == "" && ipt.Base64 == "" {
		return errors.New("url or base64 is required")
	}

	if ipt.URL != "" && ipt.Base64 != "" {
		return errors.New("url and base64 are exclusive")
	}

	return nil
}

// @Router		/v1/lgtms [post]
// @Accept		json
// @Param		body	body		createLGTMInput	true	"body"
// @Success	201		{object}	models.LGTM
// @Failure	400		{object}	ErrorResponse
// @Failure	500		{object}	ErrorResponse
func (svc *lgtmService) CreateLGTM(ctx *gin.Context) {
	var ipt createLGTMInput
	if err := ctx.ShouldBindJSON(&ipt); err != nil {
		log.Info(ctx, "failed to bind json", "error", err)
		renderError(ctx, http.StatusBadRequest, ErrCodeBadRequest)
		return
	}

	if err := ipt.Validate(); err != nil {
		log.Info(ctx, "failed to validate input", "error", err)
		renderError(ctx, http.StatusBadRequest, ErrCodeBadRequest)
		return
	}

	var src string
	var data []byte

	switch {
	case ipt.Base64 != "":
		d, err := svc.readFromBase64(ctx, ipt.Base64)
		if err != nil {
			log.Info(ctx, "failed to read from base64", "error", err)
			renderError(ctx, http.StatusBadRequest, ErrCodeBadRequest)
			return
		}
		data = d
		src = "Base64"
	case ipt.URL != "":
		d, err := svc.readFromURL(ctx, ipt.URL)
		if err != nil {
			log.Info(ctx, "failed to read from url", "error", err)
			renderError(ctx, http.StatusBadRequest, ErrCodeFailedToGetImage)
			return
		}
		data = d
		src = ipt.URL
	}

	genopts := []lgtmgen.GenerateOption{}
	if ipt.Options != nil {
		genopts = append(genopts, lgtmgen.WithTextColor(ipt.Options.TextColor))
	}

	lgtm, err := svc.repo.CreateLGTM(ctx, data, genopts...)
	if err != nil {
		if errors.Is(err, lgtmgen.ErrInvalidOption) {
			log.Info(ctx, "invalid option", "error", err)
			renderError(ctx, http.StatusBadRequest, ErrCodeBadRequest)
			return
		}
		if errors.Is(err, lgtmgen.ErrUnsupportImageFormat) {
			log.Info(ctx, "unsupported image format", "error", err)
			renderError(ctx, http.StatusBadRequest, ErrCodeUnsupportedImageFormat)
			return
		}

		log.Error(ctx, "failed to create lgtm", err)
		renderError(ctx, http.StatusInternalServerError, ErrCodeInternalServerError)
		return
	}

	err = svc.repo.SendLGTMCreatedMessage(ctx, &repo.LGTMCreatedMessage{
		LGTM:     lgtm,
		Source:   src,
		ClientIP: util.GetClientIPFromContext(ctx),
	})
	if err != nil {
		log.Error(ctx, "failed to send lgtm created message", err)
	}

	ctx.JSON(http.StatusCreated, lgtm)
}

func (svc *lgtmService) DeleteLGTM(ctx context.Context, id string) error {
	if err := svc.repo.DeleteLGTM(ctx, id); err != nil {
		return errors.Wrap(err, "failed to delete lgtm")
	}

	return nil
}

func (svc *lgtmService) TagLGTM(ctx context.Context) error {
	lgtm, err := svc.repo.TagLGTM(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to tag lgtm")
	}

	if lgtm != nil {
		log.Info(ctx, "tagged lgtm", "id", lgtm.ID)

		for lang, tags := range map[string][]string{
			"ja": lgtm.TagsJa,
			"en": lgtm.TagsEn,
		} {
			for _, tag := range tags {
				if err := svc.repo.IncrementTagByName(ctx, tag, lang); err != nil {
					return errors.Wrap(err, "failed to upsert tags")
				}

				// TODO: sync to algolia
			}
		}
	} else {
		log.Info(ctx, "no lgtm to tag")
	}

	return nil
}

func (svc *lgtmService) readFromBase64(ctx context.Context, b string) ([]byte, error) {
	d, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode base64")
	}
	return d, nil
}

func (svc *lgtmService) readFromURL(ctx context.Context, u string) ([]byte, error) {
	log.Info(ctx, "request", "url", u)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	resp, err := svc.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get image")
	}
	defer resp.Body.Close()
	log.Info(ctx, "response", "status", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "failed to get image")
	}

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return nil, errors.Wrap(err, "failed to read image")
	}

	return buf.Bytes(), nil
}

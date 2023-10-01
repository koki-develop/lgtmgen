package service

import (
	"net/http"
	"unicode/utf8"

	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
)

type reportService struct {
	repo *repo.Repository
}

func newReportService(repo *repo.Repository) *reportService {
	return &reportService{repo: repo}
}

type createReportInput struct {
	LGTMID string            `json:"lgtm_id"`
	Type   models.ReportType `json:"type"`
	Text   string            `json:"text"`
}

func (ipt createReportInput) Validate() error {
	if ipt.LGTMID == "" {
		return errors.New("lgtm_id is required")
	}

	if ipt.Type == "" {
		return errors.New("type is required")
	}
	switch ipt.Type {
	case models.ReportTypeInappropriate, models.ReportTypeIllegal, models.ReportTypeOther:
		// ok
	default:
		return errors.New("type is invalid")
	}

	if ipt.Text == "" {
		return errors.New("text is required")
	}
	if utf8.RuneCountInString(ipt.Text) > 1000 {
		return errors.New("text is too long")
	}

	return nil
}

//	@Router		/v1/reports [post]
//	@Accept		json
//	@Param		body	body		createReportInput	true	"body"
//	@Success	201		{object}	models.Report
//	@Failure	400		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
func (s *reportService) CreateReport(ctx *gin.Context) {
	var ipt createReportInput
	if err := ctx.ShouldBindJSON(&ipt); err != nil {
		log.Info(ctx, "failed to bind json", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeBadRequest})
		return
	}

	if err := ipt.Validate(); err != nil {
		log.Info(ctx, "failed to validate", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeBadRequest})
		return
	}

	lgtm, err := s.repo.FindLGTM(ctx, ipt.LGTMID)
	if err != nil {
		log.Error(ctx, "failed to find lgtm", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": ErrCodeInternalServerError})
		return
	}
	if lgtm == nil {
		log.Info(ctx, "lgtm not found", "lgtm_id", ipt.LGTMID)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": ErrCodeBadRequest})
		return
	}

	rp, err := s.repo.CreateReport(ctx, ipt.LGTMID, ipt.Type, ipt.Text)
	if err != nil {
		log.Error(ctx, "failed to create report", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": ErrCodeInternalServerError})
		return
	}

	err = s.repo.SendReportCreatedMessage(ctx, &repo.ReportCreatedMessage{
		Report: rp,
	})
	if err != nil {
		log.Error(ctx, "failed to send report created message", err)
	}

	ctx.JSON(http.StatusCreated, rp)
}

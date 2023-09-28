package service

import "github.com/gin-gonic/gin"

type ErrCode string

const (
	// 4xx
	ErrCodeBadRequest             ErrCode = "BAD_REQUEST"
	ErrCodeUnsupportedImageFormat ErrCode = "UNSUPPORTED_IMAGE_FORMAT"
	ErrCodeFailedToGetImage       ErrCode = "FAILED_TO_GET_IMAGE"
	ErrCodeNotFound               ErrCode = "NOT_FOUND"
	ErrCodeRateLimitReached       ErrCode = "RATE_LIMIT_REACHED"

	// 5xx
	ErrCodeInternalServerError ErrCode = "INTERNAL_SERVER_ERROR"
)

type ErrorResponse struct {
	Code ErrCode `json:"code"`
}

func renderError(ctx *gin.Context, status int, code ErrCode) {
	ctx.JSON(status, ErrorResponse{Code: code})
}

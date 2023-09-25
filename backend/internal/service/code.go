package service

type ErrCode string

const (
	// 4xx
	ErrCodeBadRequest             ErrCode = "BAD_REQUEST"
	ErrCodeUnsupportedImageFormat ErrCode = "UNSUPPORTED_IMAGE_FORMAT"
	ErrCodeFailedToGetImage       ErrCode = "FAILED_TO_GET_IMAGE"

	// 5xx
	ErrCodeInternalServerError ErrCode = "INTERNAL_SERVER_ERROR"
)

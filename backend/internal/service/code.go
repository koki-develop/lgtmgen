package service

type ErrCode string

const (
	// 4xx
	ErrCodeBadRequest ErrCode = "BAD_REQUEST"

	// 5xx
	ErrCodeInternalServerError ErrCode = "INTERNAL_SERVER_ERROR"
)

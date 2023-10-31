package lgtmgen

import (
	"github.com/cockroachdb/errors"
)

var (
	ErrUnsupportImageFormat = errors.New("unsupported image format")
	ErrInvalidOption        = errors.New("invalid option")
)

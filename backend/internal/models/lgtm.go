package models

import (
	"net/url"
	"time"

	"github.com/cockroachdb/errors"
)

type LGTM struct {
	ID        string     `json:"id" dynamodbav:"id"         validate:"required"`
	Status    LGTMStatus `json:"-"  dynamodbav:"status"                        `
	CreatedAt time.Time  `json:"-"  dynamodbav:"created_at"                    `
}

type LGTMs []*LGTM

func (l *LGTM) URL(base string) (string, error) {
	u, err := url.JoinPath(base, l.ID)
	if err != nil {
		return "", errors.Wrap(err, "failed to join path")
	}
	return u, nil
}

type LGTMStatus string

const (
	LGTMStatusOK      LGTMStatus = "ok"
	LGTMStatusPending LGTMStatus = "pending"
)

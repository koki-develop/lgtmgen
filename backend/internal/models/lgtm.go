package models

import "time"

type LGTM struct {
	ID        string     `json:"id" dynamodbav:"id"`
	Status    LGTMStatus `json:"-"  dynamodbav:"status"`
	CreatedAt time.Time  `json:"-"  dynamodbav:"created_at"`
}

type LGTMs []*LGTM

type LGTMStatus string

const (
	LGTMStatusOK      LGTMStatus = "ok"
	LGTMStatusPending LGTMStatus = "pending"
)

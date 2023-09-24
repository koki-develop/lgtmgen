package models

import "time"

type LGTM struct {
	ID        string    `json:"id" dynamodbav:"id"`
	Status    string    `json:"-"  dynamodbav:"status"`
	CreatedAt time.Time `json:"-"  dynamodbav:"created_at"`
}

type LGTMs []*LGTM

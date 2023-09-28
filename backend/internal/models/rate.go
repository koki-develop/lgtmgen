package models

import "time"

type Rate struct {
	IP      string    `json:"ip"       dynamodbav:"ip"`
	Tier    string    `json:"tier"     dynamodbav:"tier"`
	Count   int       `json:"count"    dynamodbav:"count"`
	ResetAt time.Time `json:"reset_at" dynamodbav:"reset_at"`
}

func (r *Rate) Expired() bool {
	return r.ResetAt.Before(time.Now())
}

func (r *Rate) LimitReached(limit int) bool {
	return r.Count >= limit
}

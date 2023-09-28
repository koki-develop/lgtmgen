package models

import (
	"time"
)

type ReportType string

const (
	ReportTypeIllegal       ReportType = "illegal"
	ReportTypeInappropriate ReportType = "inappropriate"
	ReportTypeOther         ReportType = "other"
)

func (t ReportType) Valid() bool {
	switch t {
	case ReportTypeIllegal, ReportTypeInappropriate, ReportTypeOther:
		return true
	default:
		return false
	}
}

type Report struct {
	ID        string     `json:"id"         dynamodbav:"id"`
	LGTMID    string     `json:"lgtm_id"    dynamodbav:"lgtm_id"`
	Type      ReportType `json:"type"       dynamodbav:"type"`
	Text      string     `json:"text"       dynamodbav:"text"`
	CreatedAt time.Time  `json:"created_at" dynamodbav:"created_at"`
}

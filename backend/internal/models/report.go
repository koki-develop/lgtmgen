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
	ID        string     `json:"id" dynamodbav:"id"`
	LGTMID    string     `json:"-"  dynamodbav:"lgtm_id"`
	Type      ReportType `json:"-"  dynamodbav:"type"`
	Text      string     `json:"-"  dynamodbav:"text"`
	CreatedAt time.Time  `json:"-"  dynamodbav:"created_at"`
}

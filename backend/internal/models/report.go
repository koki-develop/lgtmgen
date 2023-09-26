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
	ID        string     `json:"id" dynamo:"id"`
	LGTMID    string     `json:"-"  dynamo:"lgtm_id"`
	Type      ReportType `json:"-"  dynamo:"type"`
	Text      string     `json:"-"  dynamo:"text"`
	CreatedAt time.Time  `json:"-"  dynamo:"created_at"`
}

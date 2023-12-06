package models

type Tag struct {
	Name  string `json:"name"  dynamodbav:"name"`
	Count int    `json:"count" dynamodbav:"count"`
	Lang  string `json:"-"     dynamodbav:"lang"`
}

type Tags []*Tag

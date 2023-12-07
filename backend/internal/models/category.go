package models

type Category struct {
	Name  string `json:"name"  dynamodbav:"name"  validate:"required"`
	Count int    `json:"count" dynamodbav:"count" validate:"required"`
	Lang  string `json:"-"     dynamodbav:"lang"`
}

type Categories []*Category

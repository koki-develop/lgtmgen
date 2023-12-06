package models

type Tag struct {
	Name  string `dynamodbav:"name"`
	Count int    `dynamodbav:"count"`
	Lang  string `dynamodbav:"lang"`
}

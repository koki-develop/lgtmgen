package models

type News struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewsList []*News

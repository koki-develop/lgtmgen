package models

type Image struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Images []*Image

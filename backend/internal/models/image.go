package models

type Image struct {
	Title string `json:"title" validate:"required"`
	URL   string `json:"url"   validate:"required"`
}

type Images []*Image

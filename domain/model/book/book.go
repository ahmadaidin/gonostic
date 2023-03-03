package book

import "time"

type Filter struct {
	PublishedBefore time.Time `json:"published_before" query:"published_before"`
}

type FindOptions struct {
	Filter Filter   `json:"-"`
	Sort   []string `json:"sort" query:"sort"`
	Search string   `json:"search" query:"search"`
}

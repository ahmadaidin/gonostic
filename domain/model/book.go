package model

import "time"

type BookFilter struct {
	Sort            []string  `json:"sort" query:"sort"`
	Search          string    `json:"search" query:"search"`
	PublishedBefore time.Time `json:"published_before" query:"published_before"`
}

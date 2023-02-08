package model

type Search struct {
	Search string `json:"search" query:"search"`
}

type Sort struct {
	Sort string `json:"sort" query:"sort"`
}

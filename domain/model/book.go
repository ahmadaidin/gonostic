package model

// type BookSort struct {
// 	string `json:"author" query:"author"`
// }

type BookFilter struct {
	Sort Sort `json:"-" `
	Search
}

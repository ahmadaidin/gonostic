package entity

type Book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Author  Person `json:"author"`
}

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

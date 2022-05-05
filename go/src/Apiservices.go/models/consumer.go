package models

type Consumer struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	AuthorId int    `json:"author_id"`
	BookId   int    `json:"book_id"`
}

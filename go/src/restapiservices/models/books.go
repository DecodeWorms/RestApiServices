package models

type Book struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	YearPublished string `json:"year_published"`
}

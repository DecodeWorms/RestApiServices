package models

type Author struct {
	Id     int    `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Gender string `jso:"gender"`
}

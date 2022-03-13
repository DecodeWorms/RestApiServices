package objects

import (
	"context"
)

type AuthorService interface {
	Create(ctx context.Context, cr *Author) error
	Authors(ctx context.Context) ([]*Author, error)
	Author(ctx context.Context, email string) (*Author, error)
	Update(ctx context.Context, cr *Author) error
	Delete(ctx context.Context, email string) error
	ValidateEmail(ctx context.Context, email string) (string, error)
}

type Author struct {
	Id     int    `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Gender string `jso:"gender"`
}

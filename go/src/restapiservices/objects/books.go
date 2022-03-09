package objects

import (
	"context"
	"restapiservices/models"
)

type Book interface {
	Create(ctx context.Context, cr models.Book) error
	Books(ctx context.Context) ([]models.Book, error)
	Book(ctx context.Context, email string) (models.Book, error)
	Update(ctx context.Context, email string) error
	Delete(ctx context.Context, email string) error
}

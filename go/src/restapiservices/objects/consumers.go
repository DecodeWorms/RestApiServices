package objects

import (
	"context"
	"restapiservices/models"
)

type Consumers interface {
	Create(ctx context.Context, cr models.Consumer) error
	Consumers(ctx context.Context) ([]models.Consumer, error)
	Consumer(ctx context.Context, email string) (models.Consumer, error)
	Update(ctx context.Context, email string) error
	Delete(ctx context.Context, email string) error
}

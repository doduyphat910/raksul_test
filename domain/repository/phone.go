package repository

import (
	"awesomeProject1/domain/entity"
	"context"
)

type PhoneRepositoryI interface {
	Create(ctx context.Context, phone entity.Phone) (entity.Phone, error)
}

package repository

import (
	"awesomeProject1/domain/entity"
	"context"
)

type ManufacturerRepositoryI interface {
	GetByID(ctx context.Context, id uint64) (entity.Manufacturer, error)
}

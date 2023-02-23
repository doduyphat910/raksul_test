package mysql

import (
	"awesomeProject1/domain/entity"
	"awesomeProject1/utils"
	"context"
	"gorm.io/gorm"
)

type ManufacturerRepository struct {
	db *gorm.DB
}

func NewManufacturerRepository() *ManufacturerRepository {
	return &ManufacturerRepository{db: utils.GetDB()}
}

func (repo *ManufacturerRepository) GetByID(ctx context.Context, id uint64) (entity.Manufacturer, error) {
	var (
		manuEnt entity.Manufacturer
		err     error
	)
	err = repo.db.Take(&manuEnt).Error
	return manuEnt, err
}

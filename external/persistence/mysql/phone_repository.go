package mysql

import (
	"awesomeProject1/domain/entity"
	"awesomeProject1/utils"
	"context"
	"gorm.io/gorm"
)

type PhoneRepository struct {
	db *gorm.DB
}

func NewPhoneRepository() *PhoneRepository {
	return &PhoneRepository{db: utils.GetDB()}
}

func (repo *PhoneRepository) Create(ctx context.Context, phone entity.Phone) (entity.Phone, error) {
	err := repo.db.Create(&phone).Error
	return phone, err
}

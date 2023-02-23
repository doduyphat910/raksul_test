package usecase

import (
	"awesomeProject1/domain/entity"
	"awesomeProject1/domain/repository"
	"context"
)

type PhoneUsecaseI interface {
	Create(ctx context.Context, phone entity.Phone) (entity.Phone, error)
}

type PhoneUsecase struct {
	manuRepo  repository.ManufacturerRepositoryI
	phoneRepo repository.PhoneRepositoryI
}

func NewPhoneUseCase(phoneRepo repository.PhoneRepositoryI, manuRepo repository.ManufacturerRepositoryI) *PhoneUsecase {
	return &PhoneUsecase{
		manuRepo:  manuRepo,
		phoneRepo: phoneRepo,
	}
}

func (uc PhoneUsecase) Create(ctx context.Context, phone entity.Phone) (entity.Phone, error) {
	_, err := uc.manuRepo.GetByID(ctx, phone.ManufacturerID)
	if err != nil {
		return entity.Phone{}, err
	}
	phoneEnt, err := uc.phoneRepo.Create(ctx, phone)
	return phoneEnt, err
}

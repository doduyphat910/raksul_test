package main

import (
	"awesomeProject1/domain/entity"
	"awesomeProject1/interface/api/presenter"
	"awesomeProject1/registry"
	"github.com/gin-gonic/gin"
)

func phoneCreateHandler(ctx *gin.Context) {
	var (
		req presenter.CreatePhoneRequest
		res presenter.CreatePhoneResponse
		err error
	)
	defer func() {
		ctx.JSON(200, gin.H{
			"data":  res,
			"error": err,
		})
	}()

	if err = ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	phoneEnt := entity.Phone{
		ManufacturerID:  req.ManufacturerID,
		Model:           req.Model,
		Memory:          req.Memory,
		ManufactureYear: req.ManufactureYear,
		OSVersion:       req.OSVersion,
		BodyColor:       req.BodyColor,
		Price:           req.Price,
	}
	phoneCreated, err := registry.InjectedPhoneUseCase().Create(ctx, phoneEnt)
	if err != nil {
		return
	}

	res = presenter.CreatePhoneResponse{
		ID: phoneCreated.ID,
		CreatePhoneRequest: presenter.CreatePhoneRequest{
			ManufacturerID:  phoneCreated.ManufacturerID,
			Model:           phoneCreated.Model,
			Memory:          phoneCreated.Memory,
			ManufactureYear: phoneCreated.ManufactureYear,
			OSVersion:       phoneCreated.OSVersion,
			BodyColor:       phoneCreated.BodyColor,
			Price:           phoneCreated.Price,
		},
		CreatedAt: phoneCreated.CreatedAt,
		UpdatedAt: phoneCreated.UpdatedAt,
	}
}

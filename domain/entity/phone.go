package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Phone struct {
	ID              uint64
	ManufacturerID  uint64
	Manufacturer    Manufacturer
	Model           string
	Memory          string
	ManufactureYear string
	OSVersion       string
	BodyColor       string
	Price           decimal.Decimal
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

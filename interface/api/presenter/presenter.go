package presenter

import (
	"github.com/shopspring/decimal"
	"time"
)

type CreatePhoneRequest struct {
	ManufacturerID  uint64          `json:"manufacturer_id"`
	Model           string          `json:"model"`
	Memory          string          `json:"memory"`
	ManufactureYear string          `json:"manufacture_year"`
	OSVersion       string          `json:"os_version"`
	BodyColor       string          `json:"body_color"`
	Price           decimal.Decimal `json:"price"`
}

type CreatePhoneResponse struct {
	ID uint64 `json:"id"`
	CreatePhoneRequest
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

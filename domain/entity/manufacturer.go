package entity

import "time"

type Manufacturer struct {
	ID        uint64
	Name      string
	Phones    []Phone
	CreatedAt time.Time
	UpdatedAt time.Time
}

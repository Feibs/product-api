package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	Id                int
	Name              string
	Stock             int
	Price             decimal.Decimal
	ProductCategoryId int
	ProductDate       string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductResponse struct {
	Id                int             `json:"id"`
	Name              string          `json:"name"`
	Stock             int             `json:"stock"`
	Price             decimal.Decimal `json:"price"`
	ProductCategoryId int             `json:"category_id"`
	ProductDate       string          `json:"product_date"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}

type ProductRequest struct {
	Name              string          `json:"name"`
	Stock             int             `json:"stock"`
	Price             decimal.Decimal `json:"price"`
	ProductCategoryId int             `json:"category_id"`
	ProductDate       string          `json:"product_date"`
}

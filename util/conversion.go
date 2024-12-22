package util

import (
	"product/dto"
	"product/entity"
)

func ConvertProductToDTO(product *entity.Product) *dto.ProductResponse {
	return &dto.ProductResponse{
		Id:                product.Id,
		Name:              product.Name,
		Stock:             product.Stock,
		Price:             product.Price,
		ProductCategoryId: product.ProductCategoryId,
		ProductDate:       product.ProductDate,
		CreatedAt:         product.CreatedAt,
		UpdatedAt:         product.UpdatedAt,
	}
}

func ConvertDTOToProduct(dto *dto.ProductRequest) *entity.Product {
	return &entity.Product{
		Name:              dto.Name,
		Stock:             dto.Stock,
		Price:             dto.Price,
		ProductCategoryId: dto.ProductCategoryId,
		ProductDate:       dto.ProductDate,
	}
}

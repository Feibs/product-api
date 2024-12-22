package usecase

import (
	"product/dto"
	"product/repo"
	"product/util"
)

type ProductUsecase interface {
	CreateProduct(productRequest *dto.ProductRequest) (*dto.ProductResponse, error)
	ListProducts() ([]*dto.ProductResponse, error)
	GetProductById(id int) (*dto.ProductResponse, error)
	GetProductsByCategoryId(id int) ([]*dto.ProductResponse, error)
	GetProductsByCategoryName(name string) ([]*dto.ProductResponse, error)
	UpdateProductById(id int, params map[string]any) error
}

type productUsecaseImpl struct {
	productRepo repo.ProductRepo
}

func NewProductUsecase(pr repo.ProductRepo) productUsecaseImpl {
	return productUsecaseImpl{
		productRepo: pr,
	}
}

func (uc productUsecaseImpl) CreateProduct(productRequest *dto.ProductRequest) (*dto.ProductResponse, error) {
	createdProduct, err := uc.productRepo.CreateProduct(util.ConvertDTOToProduct(productRequest))
	if err != nil {
		return nil, err
	}
	return util.ConvertProductToDTO(createdProduct), nil
}

func (uc productUsecaseImpl) ListProducts() ([]*dto.ProductResponse, error) {
	products, err := uc.productRepo.ListProducts()
	if err != nil {
		return nil, err
	}
	productsResponse := []*dto.ProductResponse{}
	for _, product := range products {
		productsResponse = append(productsResponse, util.ConvertProductToDTO(&product))
	}
	return productsResponse, nil
}

func (uc productUsecaseImpl) GetProductById(id int) (*dto.ProductResponse, error) {
	product, err := uc.productRepo.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return util.ConvertProductToDTO(product), nil
}

func (uc productUsecaseImpl) GetProductsByCategoryId(id int) ([]*dto.ProductResponse, error) {
	products, err := uc.productRepo.GetProductsByCategoryId(id)
	if err != nil {
		return nil, err
	}
	productsResponse := []*dto.ProductResponse{}
	for _, product := range products {
		productsResponse = append(productsResponse, util.ConvertProductToDTO(&product))
	}
	return productsResponse, nil
}

func (uc productUsecaseImpl) GetProductsByCategoryName(name string) ([]*dto.ProductResponse, error) {
	products, err := uc.productRepo.GetProductsByCategoryName(name)
	if err != nil {
		return nil, err
	}
	productsResponse := []*dto.ProductResponse{}
	for _, product := range products {
		productsResponse = append(productsResponse, util.ConvertProductToDTO(&product))
	}
	return productsResponse, nil
}

func (uc productUsecaseImpl) UpdateProductById(id int, params map[string]any) error {
	err := uc.productRepo.UpdateProductById(id, params)
	if err != nil {
		return err
	}
	return nil
}

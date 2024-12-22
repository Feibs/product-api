package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"product/dto"
	"product/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func NewProductHandler(uc usecase.ProductUsecase) ProductHandler {
	return ProductHandler{
		usecase: uc,
	}
}

func (h ProductHandler) AddProductHandler(ctx *gin.Context) {
	var productRequest dto.ProductRequest
	err := ctx.ShouldBindJSON(&productRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Mismatch data type or malformed request"})
		return
	}

	productResponse, err := h.usecase.CreateProduct(&productRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Mismatch data type or malformed request"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product created", "data": productResponse})
}

func (h ProductHandler) GetProductsHandler(ctx *gin.Context) {
	if rawCategoryID, found := ctx.GetQuery("category_id"); found {
		categoryId, err := strconv.Atoi(rawCategoryID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Id should be integer"})
			return
		}
		productsResponse, err := h.usecase.GetProductsByCategoryId(categoryId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Server error"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "data": productsResponse})
		return
	}

	if categoryName, found := ctx.GetQuery("category_name"); found {
		productsResponse, err := h.usecase.GetProductsByCategoryName(categoryName)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Server error"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "data": productsResponse})
		return
	}

	productsResponse, err := h.usecase.ListProducts()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OK", "data": productsResponse})
}

func (h ProductHandler) GetProductByIdHandler(ctx *gin.Context) {
	rawID := ctx.Param("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Id should be integer"})
		return
	}
	if id < 1 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Id should be positive"})
		return
	}

	productResponse, err := h.usecase.GetProductById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "OK", "data": productResponse})
}

func (h ProductHandler) UpdateProductByIdHandler(ctx *gin.Context) {
	rawID := ctx.Param("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Id should be integer"})
		return
	}
	if id < 1 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Id should be positive"})
		return
	}

	queryParams := map[string]any{"id": id}
	err = ctx.ShouldBindJSON(&queryParams)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Mismatch data type or malformed request"})
		return
	}

	err = h.usecase.UpdateProductById(id, queryParams)
	if errors.Is(err, sql.ErrNoRows) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Mismatch data type or malformed request"})
		return
	}

	productResponse, _ := h.usecase.GetProductById(id)
	ctx.JSON(http.StatusOK, gin.H{"message": "Product updated", "data": productResponse})
}
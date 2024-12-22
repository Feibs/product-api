package main

import (
	"log"
	"net/http"
	"product/database"
	"product/handler"
	"product/repo"
	"product/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	Timeout = 5
	Port    = ":8080"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("unable to connect to the database: %v", err)
	}
	defer db.Close()

	productRepo := repo.NewProductRepo(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	r := gin.Default()
	r.GET("/products", productHandler.GetProductsHandler)
	r.GET("/products/:id", productHandler.GetProductByIdHandler)
	r.POST("/products", productHandler.AddProductHandler)
	r.PATCH("/products/:id", productHandler.UpdateProductByIdHandler)

	s := &http.Server{
		Addr:         Port,
		Handler:      r,
		ReadTimeout:  Timeout * time.Second,
		WriteTimeout: Timeout * time.Second,
	}

	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

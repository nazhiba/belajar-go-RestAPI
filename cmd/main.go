package main

import (
	"go-test/infrastructure/handler"
	"go-test/infrastructure/repository"
	"go-test/application"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := &repository.DBProductRepository{}
	service := &application.ProductService{ProductRepository: repo}
	productHandler := &handler.ProductHandler{Service: service}
	r := gin.Default()

	r.GET("/products", productHandler.GetProducts)
	r.POST("/products", productHandler.CreateProduct)

	r.Run(":8080")

}
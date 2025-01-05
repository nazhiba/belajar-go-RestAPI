package handler

import (
	"go-test/domain"
	"net/http"
	"go-test/application"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *application.ProductService
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.Service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.Service.AddProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)

}
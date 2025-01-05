package application

import (
	"go-test/domain"
)

type ProductService struct {
	ProductRepository ProductRepository
}

func (s *ProductService) GetAllProducts() ([]domain.Product, error) {
	return s.ProductRepository.FindAll()
}

func (s *ProductService) AddProduct(product domain.Product) error {
	return s.ProductRepository.Save(product)
}

type ProductRepository interface {
	FindAll() ([]domain.Product, error)
	Save(product domain.Product) error
}
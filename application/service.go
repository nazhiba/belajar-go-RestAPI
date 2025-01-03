package productservice

import (
	"go-test/domain"
)

type ServiceProduct struct {
	ProductRepository ProductRepository
}

func (s *ServiceProduct) GetAllProducts() ([]domain.Product, error) {
	return s.ProductRepository.FindAll()
}

func (s *ServiceProduct) AddProduct(product domain.Product) error {
	return s.ProductRepository.Save(product)
}

type ProductRepository interface {
	FindAll() ([]domain.Product, error)
	Save(product domain.Product) error
}
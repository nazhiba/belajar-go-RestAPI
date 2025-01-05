package repository

import "go-test/domain"

type DBProductRepository struct {}

func (r *DBProductRepository) FindAll() ([]domain.Product, error) {
	return []domain.Product{
		{ID: 1, Name: "Laptop", Price: 1000},
		{ID: 2, Name: "Mouse", Price: 25},
			}, nil
}

func (r *DBProductRepository) Save(product domain.Product) error {
	return nil
}
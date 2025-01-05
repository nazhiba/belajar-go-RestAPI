package repository
import (
	"go-test/domain"
	"sync"
)

type DBProductRepository struct {
	mu sync.Mutex
	products []domain.Product
}

func NewDBProductRepository() *DBProductRepository {
	return &DBProductRepository{
		products: []domain.Product{
			{ID: 1, Name: "Laptop", Price: 1000},
			{ID: 2, Name: "Mouse", Price: 25},
		},
	}
}

func (r *DBProductRepository) FindAll() ([]domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.products, nil
}

func (r *DBProductRepository) Save(product domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.products = append(r.products, product)
	return nil
}
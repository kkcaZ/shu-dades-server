package domain

import "github.com/kkcaz/shu-dades-server/pkg/models"

type ProductRepository interface {
	Get(id int) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Create(product models.Product) error
	Delete(id int) error
}

type ProductUseCase interface {
	Get(id int) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Search(pageNumber int, pageSize int, sortBy models.SortBy, order models.Order) ([]models.Product, error)
	Create(product models.Product) error
	Update(product *models.Product) error
	Delete(id int) error
}

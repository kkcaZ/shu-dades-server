package domain

import "github.com/kkcaz/shu-dades-server/pkg/models"

type ProductRepository interface {
	Get(id string) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Create(product models.Product) error
	Delete(id string) error
	Subscribe(productId string, subType string, userId string) error
	Unsubscribe(productId string, subType string, userId string) error
	GetSubscriptions(subType string) ([]models.ProductSubscription, error)
	GetSubscriptionsByUser(userId string) ([]models.ProductSubscription, error)
}

type ProductUseCase interface {
	Get(id string) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Search(pageNumber int, pageSize int, sortBy models.SortBy, order models.Order) ([]models.Product, error)
	Create(product models.Product) error
	Update(product *models.Product) error
	Delete(id string) error
	Subscribe(productId string, subType string, userId string) error
	Unsubscribe(productId string, subType string, userId string) error
	SendProductNotifications(subType string) error
	GetProductSubscriptions(userId string) ([]models.ProductSubscription, error)
}

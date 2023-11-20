package product

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"github.com/pkg/errors"
	"log/slog"
	"slices"
	"sort"
)

type productUseCase struct {
	ProductRepository domain.ProductRepository
	Notification      domain.NotificationUseCase
	Logger            slog.Logger
}

func NewProductUseCase(productRepository domain.ProductRepository, notification domain.NotificationUseCase, logger slog.Logger) domain.ProductUseCase {
	return &productUseCase{
		ProductRepository: productRepository,
		Notification:      notification,
		Logger:            logger,
	}
}

func (p productUseCase) Get(id string) (*models.Product, error) {
	product, err := p.ProductRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p productUseCase) GetAll() ([]models.Product, error) {
	products, err := p.ProductRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p productUseCase) Search(pageNumber int, pageSize int, sortBy models.SortBy, order models.Order) ([]models.Product, error) {
	products, err := p.GetAll()
	if err != nil {
		return nil, err
	}

	switch sortBy {
	case models.Name:
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})
	case models.Quantity:
		sort.Slice(products, func(i, j int) bool {
			return products[i].Quantity < products[j].Quantity
		})
	}

	if order == models.Desc {
		slices.Reverse(products)
	}

	start := (pageNumber - 1) * pageSize
	end := pageNumber * pageSize

	if start > len(products) {
		return nil, nil
	}

	if end > len(products) {
		end = len(products)
	}

	return products[start:end], nil
}

func (p productUseCase) Create(product models.Product) error {
	p.Logger.Info("creating product", "product", product)
	product.Id = uuid.New().String()
	err := p.ProductRepository.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (p productUseCase) Update(product *models.Product) error {
	existingProduct, err := p.ProductRepository.Get(product.Id)
	if err != nil {
		return err
	}

	if existingProduct == nil {
		return errors.New("product not found")
	}

	err = p.ProductRepository.Delete(product.Id)
	if err != nil {
		return err
	}

	err = p.ProductRepository.Create(*product)
	if err != nil {
		return err
	}

	return nil
}

func (p productUseCase) Delete(id string) error {
	p.Logger.Info("deleting product", "id", id)
	err := p.ProductRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (p productUseCase) Subscribe(productId string, subType string, userId string) error {
	err := p.ProductRepository.Subscribe(productId, subType, userId)
	if err != nil {
		return err
	}

	return nil
}

func (p productUseCase) Unsubscribe(productId string, subType string, userId string) error {
	err := p.ProductRepository.Unsubscribe(productId, subType, userId)
	if err != nil {
		return err
	}

	return nil
}

func (p productUseCase) SendProductNotifications(subType string) error {
	p.Logger.Info("sending product notifications", "subType", subType)

	subscriptions, err := p.ProductRepository.GetSubscriptions(subType)
	if err != nil {
		return err
	}

	for _, subscription := range subscriptions {
		product, err := p.Get(subscription.ProductId)
		if err != nil {
			p.Logger.Error("failed to get product", "error", err)
			continue
		}

		productUpdate := fmt.Sprintf("Product %s has %v quantity remaining", product.Name, product.Quantity)

		p.Logger.Info("sending notification to users", "users", subscription.Users)
		err = p.Notification.AddForUsers(productUpdate, subscription.Users)
		if err != nil {
			p.Logger.Error("failed to send notification", "error", err)
			continue
		}
	}

	return nil
}

func (p productUseCase) GetProductSubscriptions(userId string) ([]models.ProductSubscription, error) {
	notifications, err := p.ProductRepository.GetSubscriptionsByUser(userId)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

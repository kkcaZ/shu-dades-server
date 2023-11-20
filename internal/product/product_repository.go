package product

import (
	"encoding/json"
	"fmt"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"log/slog"
	"os"
)

type productData struct {
	Products []models.Product `json:"Products"`
}

type productRepository struct {
	Logger               slog.Logger
	Products             []models.Product
	ProductSubscriptions []models.ProductSubscription
}

func NewProductRepository(logger slog.Logger) domain.ProductRepository {
	products, err := readProducts()
	if err != nil {
		panic(err)
	}

	subscriptions := createSubscriptions(products)

	return &productRepository{
		Logger:               logger,
		Products:             products,
		ProductSubscriptions: subscriptions,
	}
}

func readProducts() ([]models.Product, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dat, err := os.ReadFile(fmt.Sprintf("%s/internal/data/product/products.json", currentDir))
	if err != nil {
		return nil, err
	}

	var productData productData
	err = json.Unmarshal(dat, &productData)
	if err != nil {
		return nil, err
	}

	return productData.Products, nil
}

func createSubscriptions(products []models.Product) []models.ProductSubscription {
	var productSubscriptions []models.ProductSubscription
	for _, product := range products {
		productSubscriptions = append(productSubscriptions, models.ProductSubscription{
			ProductId: product.Id,
			SubType:   "hourly",
			Users:     make([]string, 0),
		})
		productSubscriptions = append(productSubscriptions, models.ProductSubscription{
			ProductId: product.Id,
			SubType:   "daily",
			Users:     make([]string, 0),
		})
	}
	return productSubscriptions
}

func (p *productRepository) Get(id string) (*models.Product, error) {
	for _, product := range p.Products {
		if product.Id == id {
			return &product, nil
		}
	}

	return nil, nil
}

func (p *productRepository) GetAll() ([]models.Product, error) {
	return p.Products, nil
}

func (p *productRepository) Create(product models.Product) error {
	p.Logger.Debug("Creating product: {product}", "product", product)
	p.Products = append(p.Products, product)
	return nil
}

func (p *productRepository) Delete(id string) error {
	for i, product := range p.Products {
		if product.Id == id {
			p.Products = append(p.Products[:i], p.Products[i+1:]...)
			return nil
		}
	}

	return nil
}

func (p *productRepository) Subscribe(productId string, subType string, userId string) error {
	p.Logger.Info("subscribing user to product", "productId", productId, "userId", userId)
	var foundSubscription *models.ProductSubscription
	for i, productSubscription := range p.ProductSubscriptions {
		if productSubscription.ProductId == productId && productSubscription.SubType == subType {
			foundSubscription = &p.ProductSubscriptions[i]
		}
	}

	if foundSubscription == nil {
		return fmt.Errorf("product subscription not found")
	}

	users := append([]string{}, foundSubscription.Users...)
	users = append(users, userId)
	foundSubscription.Users = users

	return nil
}

func (p *productRepository) Unsubscribe(productId string, subType string, userId string) error {
	p.Logger.Info("unsubscribing user to product", "productId", productId, "userId", userId)
	var foundSubscription *models.ProductSubscription
	for i, productSubscription := range p.ProductSubscriptions {
		if productSubscription.ProductId == productId && productSubscription.SubType == subType {
			foundSubscription = &p.ProductSubscriptions[i]
		}
	}

	if foundSubscription == nil {
		return fmt.Errorf("product subscription not found")
	}

	var users []string
	for _, user := range foundSubscription.Users {
		if user != userId {
			users = append(users, user)
		}
	}
	foundSubscription.Users = users

	return nil
}

func (p *productRepository) GetSubscriptions(subType string) ([]models.ProductSubscription, error) {
	var subscriptions []models.ProductSubscription
	for _, productSubscription := range p.ProductSubscriptions {
		if productSubscription.SubType == subType {
			subscriptions = append(subscriptions, productSubscription)
		}
	}

	return subscriptions, nil
}

func (p *productRepository) GetSubscriptionsByUser(userId string) ([]models.ProductSubscription, error) {
	var subscriptions []models.ProductSubscription
	for _, productSubscription := range p.ProductSubscriptions {
		for _, user := range productSubscription.Users {
			if user == userId {
				subscriptions = append(subscriptions, productSubscription)
			}
		}
	}

	return subscriptions, nil
}

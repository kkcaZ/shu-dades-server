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
	Products []models.Product `json:"products"`
}

type productRepository struct {
	Logger   slog.Logger
	products []models.Product
}

func NewProductRepository(logger slog.Logger) domain.ProductRepository {
	products, err := readProducts()
	if err != nil {
		panic(err)
	}

	return &productRepository{
		Logger:   logger,
		products: products,
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

func (p *productRepository) Get(id string) (*models.Product, error) {
	for _, product := range p.products {
		if product.Id == id {
			return &product, nil
		}
	}

	return nil, nil
}

func (p *productRepository) GetAll() ([]models.Product, error) {
	return p.products, nil
}

func (p *productRepository) Create(product models.Product) error {
	p.Logger.Debug("Creating product: {product}", "product", product)
	p.products = append(p.products, product)
	return nil
}

func (p *productRepository) Delete(id string) error {
	for i, product := range p.products {
		if product.Id == id {
			p.products = append(p.products[:i], p.products[i+1:]...)
			return nil
		}
	}

	return nil
}

package product

import (
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"github.com/pkg/errors"
	"slices"
	"sort"
)

type productUseCase struct {
	ProductRepository domain.ProductRepository
}

func NewProductUseCase(productRepository domain.ProductRepository) domain.ProductUseCase {
	return &productUseCase{
		ProductRepository: productRepository,
	}
}

func (p productUseCase) Get(id int) (*models.Product, error) {
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

func (p productUseCase) Create(product *models.Product) error {
	products, err := p.GetAll()
	if err != nil {
		return err
	}

	if products == nil {
		product.Id = 1
	} else {
		product.Id = products[len(products)-1].Id + 1
	}

	err = p.ProductRepository.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (p productUseCase) Update(product *models.Product) error {
	product, err := p.ProductRepository.Get(product.Id)
	if err != nil {
		return err
	}

	if product == nil {
		return errors.New("product not found")
	}

	err = p.ProductRepository.Delete(product.Id)
	if err != nil {
		return err
	}

	err = p.ProductRepository.Create(product)
	if err != nil {
		return err
	}

	return nil
}

func (p productUseCase) Delete(id int) error {
	err := p.ProductRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

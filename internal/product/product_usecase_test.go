package product

import (
	"fmt"
	"github.com/kkcaz/shu-dades-server/internal/domain/mocks"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

func TestProductUseCase_Get(t *testing.T) {
	testCases := []struct {
		name      string
		productId string
		product   *models.Product
		err       error
	}{
		{
			name:      "Happy path",
			productId: "1",
			product: &models.Product{
				Id: "1",
			},
			err: nil,
		},
		{
			name:      "Sad path",
			productId: "2",
			product:   nil,
			err:       errors.New("not found"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewProductRepository(t)
			testUc := NewProductUseCase(repo, nil, *logger)

			repo.On("Get", testCase.productId).Return(testCase.product, testCase.err)

			product, err := testUc.Get(testCase.productId)
			assert.Equal(t, testCase.product, product)
			if testCase.err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestProductHandler_GetAll(t *testing.T) {
	testCases := []struct {
		name     string
		products []models.Product
		err      error
	}{
		{
			name: "Happy path",
			products: []models.Product{
				{
					Id: "1",
				},
				{
					Id: "2",
				},
			},
			err: nil,
		},
		{
			name:     "Sad path",
			products: nil,
			err:      errors.New("not found"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewProductRepository(t)
			testUc := NewProductUseCase(repo, nil, *logger)

			repo.On("GetAll").Return(testCase.products, testCase.err)

			products, err := testUc.GetAll()
			assert.Equal(t, testCase.products, products)
			if testCase.err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestProductUseCase_Search(t *testing.T) {
	allProducts := []models.Product{
		{
			Id:       "1",
			Name:     "A",
			Quantity: 1,
		},
		{
			Id:       "2",
			Name:     "B",
			Quantity: 2,
		},
		{
			Id:       "3",
			Name:     "C",
			Quantity: 3,
		},
		{
			Id:       "4",
			Name:     "D",
			Quantity: 4,
		},
		{
			Id:       "5",
			Name:     "E",
			Quantity: 5,
		},
	}

	testCases := []struct {
		name                 string
		pageNumber           int
		pageSize             int
		sortBy               models.SortBy
		order                models.Order
		products             []models.Product
		err                  error
		expectedErr          error
		expectedProductCount int
	}{
		{
			name:                 "Happy path - Name ascending",
			pageNumber:           1,
			pageSize:             len(allProducts),
			sortBy:               models.Name,
			order:                models.Asc,
			products:             allProducts,
			err:                  nil,
			expectedErr:          nil,
			expectedProductCount: len(allProducts),
		},
		{
			name:                 "Happy path - Name descending",
			pageNumber:           1,
			pageSize:             len(allProducts),
			sortBy:               models.Name,
			order:                models.Desc,
			products:             allProducts,
			err:                  nil,
			expectedErr:          nil,
			expectedProductCount: len(allProducts),
		},
		{
			name:                 "Happy path - Quantity ascending",
			pageNumber:           1,
			pageSize:             len(allProducts),
			sortBy:               models.Quantity,
			order:                models.Asc,
			products:             allProducts,
			err:                  nil,
			expectedErr:          nil,
			expectedProductCount: len(allProducts),
		},
		{
			name:                 "Happy path - Quantity descending",
			pageNumber:           1,
			pageSize:             len(allProducts),
			sortBy:               models.Quantity,
			order:                models.Desc,
			products:             allProducts,
			err:                  nil,
			expectedErr:          nil,
			expectedProductCount: len(allProducts),
		},
		{
			name:                 "Happy path - Request more than available",
			pageNumber:           1,
			pageSize:             len(allProducts) + 5,
			sortBy:               models.Name,
			order:                models.Asc,
			products:             allProducts,
			err:                  nil,
			expectedErr:          nil,
			expectedProductCount: len(allProducts),
		},
		{
			name:        "Sad path",
			pageNumber:  1,
			pageSize:    10,
			sortBy:      "name",
			order:       "asc",
			products:    nil,
			err:         errors.New("not found"),
			expectedErr: errors.New("not found"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewProductRepository(t)
			testUc := NewProductUseCase(repo, nil, *logger)

			repo.On("GetAll").Return(testCase.products, testCase.err)

			products, err := testUc.Search(testCase.pageNumber, testCase.pageSize, testCase.sortBy, testCase.order)
			assert.Equal(t, testCase.expectedProductCount, len(products))
			assert.Equal(t, testCase.products, products)
			if testCase.expectedErr != nil {
				assert.Error(t, err)
				return
			}

			if testCase.order == models.Asc && testCase.sortBy == models.Name {
				assert.True(t, products[0].Name < products[1].Name)
			} else if testCase.order == models.Desc && testCase.sortBy == models.Name {
				assert.True(t, products[0].Name > products[1].Name)
			} else if testCase.order == models.Asc && testCase.sortBy == models.Quantity {
				assert.True(t, products[0].Quantity < products[1].Quantity)
			} else if testCase.order == models.Desc && testCase.sortBy == models.Quantity {
				assert.True(t, products[0].Quantity > products[1].Quantity)
			}
		})
	}
}

func TestProductUseCase_Create(t *testing.T) {
	testCases := []struct {
		name    string
		product models.Product
		err     error
	}{
		{
			name: "Happy path",
			product: models.Product{
				Id: "1",
			},
			err: nil,
		},
		{
			name: "Sad path",
			product: models.Product{
				Id: "1",
			},
			err: errors.New("not found"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewProductRepository(t)
			testUc := NewProductUseCase(repo, nil, *logger)

			repo.On("Create", mock.AnythingOfType("models.Product")).Return(testCase.err)

			err := testUc.Create(testCase.product)
			if testCase.err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestProductUseCase_Update(t *testing.T) {
	testCases := []struct {
		name            string
		productId       string
		existingProduct *models.Product
		getErr          error
		product         *models.Product
		deleteErr       error
		createErr       error
	}{
		{
			name:      "Happy path",
			productId: "1",
			existingProduct: &models.Product{
				Id: "1",
			},
			getErr: nil,
			product: &models.Product{
				Id:   "1",
				Name: "A",
			},
			deleteErr: nil,
			createErr: nil,
		},
		{
			name:            "Sad path - get returns error",
			productId:       "1",
			existingProduct: nil,
			getErr:          errors.New("something went wrong"),
			product: &models.Product{
				Id:   "1",
				Name: "A",
			},
		},
		{
			name:            "Sad path - product not found",
			productId:       "1",
			existingProduct: nil,
			getErr:          nil,
			product: &models.Product{
				Id:   "1",
				Name: "A",
			},
		},
		{
			name:      "Sad path - delete returns error",
			productId: "1",
			existingProduct: &models.Product{
				Id: "1",
			},
			getErr: nil,
			product: &models.Product{
				Id:   "1",
				Name: "A",
			},
			deleteErr: errors.New("something went wrong"),
		},
		{
			name:      "Sad path - create returns error",
			productId: "1",
			existingProduct: &models.Product{
				Id: "1",
			},
			getErr: nil,
			product: &models.Product{
				Id:   "1",
				Name: "A",
			},
			deleteErr: nil,
			createErr: errors.New("something went wrong"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewProductRepository(t)
			testUc := NewProductUseCase(repo, nil, *logger)

			repo.On("Get", testCase.productId).Return(testCase.existingProduct, testCase.getErr)
			if testCase.getErr == nil && testCase.existingProduct != nil {
				repo.On("Delete", testCase.productId).Return(testCase.deleteErr)
				if testCase.deleteErr == nil {
					repo.On("Create", mock.AnythingOfType("models.Product")).Return(testCase.createErr)
				}
			}

			err := testUc.Update(testCase.product)
			if testCase.getErr != nil ||
				testCase.deleteErr != nil ||
				testCase.createErr != nil {
				assert.Error(t, err)
			}

			if testCase.getErr == nil && testCase.existingProduct == nil {
				assert.ErrorContains(t, err, "product not found")
			}
		})
	}
}

func TestProductUseCase_Delete(t *testing.T) {
	testCases := []struct {
		name      string
		productId string
		err       error
	}{
		{
			name:      "Happy path",
			productId: "1",
			err:       nil,
		},
		{
			name:      "Sad path",
			productId: "1",
			err:       errors.New("not found"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewProductRepository(t)
			testUc := NewProductUseCase(repo, nil, *logger)

			repo.On("Delete", testCase.productId).Return(testCase.err)

			err := testUc.Delete(testCase.productId)
			if testCase.err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestProductUseCase_Subscribe(t *testing.T) {
	testCases := []struct {
		name      string
		productId string
		subType   string
		userId    string
		err       error
	}{
		{
			name:      "Happy path",
			productId: "1",
			subType:   "daily",
			userId:    "1",
			err:       nil,
		},
		{
			name:      "Sad path",
			productId: "1",
			subType:   "daily",
			userId:    "1",
			err:       errors.New("not found"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewProductRepository(t)
			testUc := NewProductUseCase(repo, nil, *logger)

			repo.On("Subscribe", testCase.productId, testCase.subType, testCase.userId).Return(testCase.err)

			err := testUc.Subscribe(testCase.productId, testCase.subType, testCase.userId)
			if testCase.err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestProductUseCase_Unsubscribe(t *testing.T) {
	testCases := []struct {
		name      string
		productId string
		subType   string
		userId    string
		err       error
	}{
		{
			name:      "Happy path",
			productId: "1",
			subType:   "daily",
			userId:    "1",
			err:       nil,
		},
		{
			name:      "Sad path",
			productId: "1",
			subType:   "daily",
			userId:    "1",
			err:       errors.New("not found"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewProductRepository(t)
			testUc := NewProductUseCase(repo, nil, *logger)

			repo.On("Unsubscribe", testCase.productId, testCase.subType, testCase.userId).Return(testCase.err)

			err := testUc.Unsubscribe(testCase.productId, testCase.subType, testCase.userId)
			if testCase.err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestProductUseCase_SendProductNotifications(t *testing.T) {
	testCases := []struct {
		name                string
		subType             string
		subscriptions       []models.ProductSubscription
		getSubscriptionsErr error
		product             models.Product
		getProductErr       error
		addNotifErr         error
		expectedErr         error
	}{
		{
			name:    "Happy path",
			subType: "daily",
			subscriptions: []models.ProductSubscription{
				{
					ProductId: "1",
					SubType:   "daily",
					Users:     []string{"1"},
				},
			},
			getSubscriptionsErr: nil,
			product: models.Product{
				Id:       "1",
				Name:     "A",
				Quantity: 1,
			},
			getProductErr: nil,
			addNotifErr:   nil,
			expectedErr:   nil,
		},
		{
			name:                "Sad path - get subscriptions error",
			subType:             "daily",
			subscriptions:       nil,
			getSubscriptionsErr: errors.New("something went wrong"),
			expectedErr:         errors.New("something went wrong"),
		},
		{
			name:    "Sad path - get product error",
			subType: "daily",
			subscriptions: []models.ProductSubscription{
				{
					ProductId: "1",
					SubType:   "daily",
					Users:     []string{"1"},
				},
			},
			getSubscriptionsErr: nil,
			product: models.Product{
				Id:       "1",
				Name:     "A",
				Quantity: 1,
			},
			getProductErr: errors.New("something went wrong"),
			expectedErr:   nil,
		},
		{
			name:    "Sad path - add notification error",
			subType: "daily",
			subscriptions: []models.ProductSubscription{
				{
					ProductId: "1",
					SubType:   "daily",
					Users:     []string{"1"},
				},
			},
			getSubscriptionsErr: nil,
			product: models.Product{
				Id:       "1",
				Name:     "A",
				Quantity: 1,
			},
			getProductErr: nil,
			addNotifErr:   errors.New("something went wrong"),
			expectedErr:   nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			logger := slog.Default()
			repo := mocks.NewProductRepository(t)
			notification := mocks.NewNotificationUseCase(t)
			testUc := NewProductUseCase(repo, notification, *logger)

			repo.On("GetSubscriptions", testCase.subType).Return(testCase.subscriptions, testCase.getSubscriptionsErr)
			if testCase.getSubscriptionsErr == nil {
				repo.On("Get", mock.AnythingOfType("string")).Return(&testCase.product, testCase.getProductErr)
				if testCase.getProductErr == nil {
					productUpdate := fmt.Sprintf("Product %s has %v quantity remaining", testCase.product.Name, testCase.product.Quantity)
					notification.On("AddForUsers", productUpdate, mock.AnythingOfType("[]string")).Return(testCase.addNotifErr)
				}
			}

			err := testUc.SendProductNotifications(testCase.subType)
			if testCase.expectedErr != nil {
				assert.Error(t, err)
			}
		})
	}
}

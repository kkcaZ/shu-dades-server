// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	models "github.com/kkcaz/shu-dades-server/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: product
func (_m *ProductRepository) Create(product models.Product) error {
	ret := _m.Called(product)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *ProductRepository) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *ProductRepository) Get(id string) (*models.Product, error) {
	ret := _m.Called(id)

	var r0 *models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Product, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *ProductRepository) GetAll() ([]models.Product, error) {
	ret := _m.Called()

	var r0 []models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Product, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscriptions provides a mock function with given fields: subType
func (_m *ProductRepository) GetSubscriptions(subType string) ([]models.ProductSubscription, error) {
	ret := _m.Called(subType)

	var r0 []models.ProductSubscription
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]models.ProductSubscription, error)); ok {
		return rf(subType)
	}
	if rf, ok := ret.Get(0).(func(string) []models.ProductSubscription); ok {
		r0 = rf(subType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProductSubscription)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(subType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubscriptionsByUser provides a mock function with given fields: userId
func (_m *ProductRepository) GetSubscriptionsByUser(userId string) ([]models.ProductSubscription, error) {
	ret := _m.Called(userId)

	var r0 []models.ProductSubscription
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]models.ProductSubscription, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) []models.ProductSubscription); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ProductSubscription)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: productId, subType, userId
func (_m *ProductRepository) Subscribe(productId string, subType string, userId string) error {
	ret := _m.Called(productId, subType, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(productId, subType, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: productId, subType, userId
func (_m *ProductRepository) Unsubscribe(productId string, subType string, userId string) error {
	ret := _m.Called(productId, subType, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(productId, subType, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

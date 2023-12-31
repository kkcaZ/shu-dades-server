// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BroadcastUseCase is an autogenerated mock type for the BroadcastUseCase type
type BroadcastUseCase struct {
	mock.Mock
}

// AddConnection provides a mock function with given fields: subscribeAddress, publishAddress
func (_m *BroadcastUseCase) AddConnection(subscribeAddress string, publishAddress string) {
	_m.Called(subscribeAddress, publishAddress)
}

// PublishToUsers provides a mock function with given fields: message, eventType, users
func (_m *BroadcastUseCase) PublishToUsers(message string, eventType string, users []string) error {
	ret := _m.Called(message, eventType, users)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string) error); ok {
		r0 = rf(message, eventType, users)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterUser provides a mock function with given fields: addr, userId
func (_m *BroadcastUseCase) RegisterUser(addr string, userId string) {
	_m.Called(addr, userId)
}

// RemoveUser provides a mock function with given fields: addr
func (_m *BroadcastUseCase) RemoveUser(addr string) {
	_m.Called(addr)
}

// NewBroadcastUseCase creates a new instance of BroadcastUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBroadcastUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *BroadcastUseCase {
	mock := &BroadcastUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

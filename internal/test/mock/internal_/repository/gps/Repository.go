// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/kevenmiano/v3/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *Repository) Create(_a0 *domain.GPS) (*domain.GPS, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *domain.GPS
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.GPS) (*domain.GPS, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*domain.GPS) *domain.GPS); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.GPS)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.GPS) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/kevenmiano/v3/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: d
func (_m *UseCase) Execute(d *domain.GPS) (*domain.GPS, error) {
	ret := _m.Called(d)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *domain.GPS
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.GPS) (*domain.GPS, error)); ok {
		return rf(d)
	}
	if rf, ok := ret.Get(0).(func(*domain.GPS) *domain.GPS); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.GPS)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.GPS) error); ok {
		r1 = rf(d)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

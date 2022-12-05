// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AverageUseCase is an autogenerated mock type for the AverageUseCase type
type AverageUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx
func (_m *AverageUseCase) Execute(ctx context.Context) (float64, error) {
	ret := _m.Called(ctx)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context) float64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAverageUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAverageUseCase creates a new instance of AverageUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAverageUseCase(t mockConstructorTestingTNewAverageUseCase) *AverageUseCase {
	mock := &AverageUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
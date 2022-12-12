package mocks

import (
	"PPOB_BACKEND/businesses/payment_method"

	 mock "github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}


func (_m *Repository) CreatePayment(paymentDomain *payment_method.Domain) payment_method.Domain {
	ret := _m.Called(paymentDomain)

	var r0 payment_method.Domain
	if rf, ok := ret.Get(0).(func(*payment_method.Domain) payment_method.Domain); ok {
		r0 = rf(paymentDomain)
	} else {
		r0 = ret.Get(0).(payment_method.Domain)
	}

	return r0
}

func (_m *Repository) DeletePayment(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}


func (_m *Repository) UpdatePaymentByID(id string, paymentDomain *payment_method.Domain) payment_method.Domain {
	ret := _m.Called(id, paymentDomain)

	var r0 payment_method.Domain

	if rf, ok := ret.Get(0).(func(string, *payment_method.Domain) payment_method.Domain); ok {
		r0 = rf(id, paymentDomain)
	} else {
		r0 = ret.Get(0).(payment_method.Domain)
	}

	return r0
}

func (_m *Repository) GetAll() []payment_method.Domain {
	ret := _m.Called()

	var r0 []payment_method.Domain
	if rf, ok := ret.Get(0).(func() []payment_method.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]payment_method.Domain)
		}
	}
	return r0
}

func (_m *Repository) GetSpecificPayment(id string) payment_method.Domain {
	ret := _m.Called(id)

	var r0 payment_method.Domain
	if rf, ok := ret.Get(0).(func(string) payment_method.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(payment_method.Domain)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
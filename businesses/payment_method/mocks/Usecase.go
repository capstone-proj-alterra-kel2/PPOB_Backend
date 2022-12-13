package mocks

import (
	"PPOB_BACKEND/businesses/payment_method"

	mock "github.com/stretchr/testify/mock"
)

type Usecase struct {
	mock.Mock
}

func (_m *Usecase) CreatePayment(paymentDomain *payment_method.Domain) payment_method.Domain {
	ret := _m.Called(paymentDomain)

	var r0 payment_method.Domain
	if rf, ok := ret.Get(0).(func(*payment_method.Domain) payment_method.Domain); ok {
		r0 = rf(paymentDomain)
	} else {
		r0 = ret.Get(0).(payment_method.Domain)
	}

	return r0
}

func (_m *Usecase) DeletePayment(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

func (_m *Usecase) UpdatePaymentByID(id string, paymentDomain *payment_method.Domain) payment_method.Domain {
	ret := _m.Called(id, paymentDomain)

	var r0 payment_method.Domain
	if rf, ok := ret.Get(0).(func(string, *payment_method.Domain) payment_method.Domain); ok {
		r0 = rf(id, paymentDomain)
	} else {
		r0 = ret.Get(0).(payment_method.Domain)
	}

	return r0
}

func (_m *Usecase) GetAll() []payment_method.Domain {
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

func (_m *Usecase) GetSpecificPayment(id string) payment_method.Domain {
	ret := _m.Called(id)

	var r0 payment_method.Domain
	if rf, ok := ret.Get(0).(func(string) payment_method.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(payment_method.Domain)
	}

	return r0
}

type mockConstructorTestingTNewUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecase creates a new instance of Usecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecase(t mockConstructorTestingTNewUsecase) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

package payment_method_test

import (
	"PPOB_BACKEND/businesses/payment_method"
	_paymentMock "PPOB_BACKEND/businesses/payment_method/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	paymentRepository _paymentMock.Repository
	paymentService    payment_method.Usecase

	paymentDomain payment_method.Domain
)

func TestMain(m *testing.M) {
	paymentService = payment_method.NewPaymentMethodUsecase(&paymentRepository)
	paymentDomain = payment_method.Domain{
		Payment_Name: "payment_name",
		Url_Payment:  "url_payment",
		Icon:         "icon",
	}
	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		paymentRepository.On("GetAll").Return([]payment_method.Domain{paymentDomain}).Once()

		result := paymentService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | Invalid", func(t *testing.T) {
		paymentRepository.On("GetAll").Return([]payment_method.Domain{}).Once()

		result := paymentService.GetAll()

		assert.Equal(t, 0, len(result))
	})

}

func TestGetSpecificPayment(t *testing.T) {
	t.Run("Get Specific Payment | Valid", func(t *testing.T) {
		paymentRepository.On("GetSpecificPayment").Return([]payment_method.Domain{paymentDomain}).Once()

		result := paymentService.GetSpecificPayment("1")

		assert.NotNil(t, result)
    
	})
}



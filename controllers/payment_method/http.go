package payment_method

import (
	"PPOB_BACKEND/businesses/payment_method"
	"PPOB_BACKEND/controllers"

	"PPOB_BACKEND/controllers/payment_method/request"
	"PPOB_BACKEND/controllers/payment_method/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	paymentUsecase payment_method.Usecase
}

func NewPaymentController(PaymentUC payment_method.Usecase) *PaymentController {
	return &PaymentController{
		paymentUsecase: PaymentUC,
	}
}


func (ctrl *PaymentController) GetAll(c echo.Context) error {
	PaymentData := ctrl.paymentUsecase.GetAll()
	
	payment := []response.Payment_Method{}

	for _, pm := range PaymentData {
		payment = append(payment, response.FromDomain(pm))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all payments", payment)
}


func (ctrl *PaymentController) GetSpecificPayment(c echo.Context) error {
	var id string = c.Param("id")

	payment := ctrl.paymentUsecase.GetSpecificPayment(id)

	if payment.ID == 0 {
		return controllers.NewResponse(c, http.StatusNotFound,"failed","payment method not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success","payment method found", response.FromDomain(payment))
}


func (ctrl *PaymentController) CreatePayment(c echo.Context) error {
	input := request.PaymentMethod{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	payment := ctrl.paymentUsecase.CreatePayment(input.ToDomain())

	return controllers.NewResponse(c, http.StatusCreated, "success", "note created", response.FromDomain(payment))
}


func (ctrl *PaymentController) UpdatePaymentByID(c echo.Context) error {
	input := request.PaymentMethod{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	var paymentId string = c.Param("id")

	err := input.Validate()	

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	payment := ctrl.paymentUsecase.UpdatePaymentByID(paymentId, input.ToDomain())

	if payment.ID == 0 {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "payment method not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "payment method updated", response.FromDomain(payment))
}


func (ctrl *PaymentController) DeletePayment(c echo.Context) error {
	var paymentId string = c.Param("id")

	isSuccess := ctrl.paymentUsecase.DeletePayment(paymentId)

	if !isSuccess {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "payment method not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "payment method deleted", "")
}
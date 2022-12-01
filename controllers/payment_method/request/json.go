package request

import (
	"PPOB_BACKEND/businesses/payment_method"

	"github.com/go-playground/validator"
)

type PaymentMethod struct {
	Payment_Name string `json:"payment_name" form:"payment_name"`
	Url_Payment string `json:"url_payment" form:"url_payment"`
	Icon string `json:"icon" form:"icon"`
}

func (req *PaymentMethod) ToDomain() *payment_method.Domain {
	return &payment_method.Domain{
		Payment_Name: req.Payment_Name,
		Url_Payment: req.Url_Payment,
		Icon: req.Icon,
	}
}

func (req *PaymentMethod) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}

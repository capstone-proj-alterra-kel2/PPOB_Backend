package request

import (
	"PPOB_BACKEND/businesses/transactions"

	"github.com/go-playground/validator"
)

type Transaction struct {
	ProductID         int    `json:"product_id" validate:"required"`
	TargetPhoneNumber string `json:"target_phone_number"`
}

func (req *Transaction) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		ProductID:         req.ProductID,
		TargetPhoneNumber: req.TargetPhoneNumber,
	}
}

func (req *Transaction) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

type TransactionUpdate struct {
	TargetPhoneNumber string `json:"target_phone_number" validate:"required"`
}

func (req *TransactionUpdate) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		TargetPhoneNumber: req.TargetPhoneNumber,
	}
}

func (req *TransactionUpdate) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

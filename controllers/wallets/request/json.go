package request

import (
	"PPOB_BACKEND/businesses/wallets"

	"github.com/go-playground/validator"
)

type UpdateBalance struct {
	Balance int `json:"balance" validate:"required"`
}

func (req *UpdateBalance) ToDomain() *wallets.UpdateBalanceDomain {
	return &wallets.UpdateBalanceDomain{
		Balance: req.Balance,
	}
}

func (req *UpdateBalance) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}
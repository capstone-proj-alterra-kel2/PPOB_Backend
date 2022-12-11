package request

import (
	"PPOB_BACKEND/businesses/wallet_histories"

	"github.com/go-playground/validator"
)

type WalletHistory struct {
	Income      int    `json:"income"`
	Outcome     int    `json:"outcome"`
	Description string `json:"description"`
}

func (req *WalletHistory) ToDomain() *wallet_histories.Domain {
	return &wallet_histories.Domain{
		Income:      req.Income,
		Outcome:     req.Outcome,
		Description: req.Description,
	}
}

func (req *WalletHistory) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

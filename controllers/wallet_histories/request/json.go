package request

import (
	"PPOB_BACKEND/businesses/wallet_histories"

	"github.com/go-playground/validator"
)

type WalletHistory struct {
	CashIn      int    `json:"cash_in"`
	CashOut     int    `json:"cash_out"`
	Description string `json:"description"`
}

func (req *WalletHistory) ToDomain() *wallet_histories.Domain {
	return &wallet_histories.Domain{
		CashIn: req.CashIn,
		CashOut: req.CashOut,
		Description: req.Description,
	}
}

func (req *WalletHistory) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

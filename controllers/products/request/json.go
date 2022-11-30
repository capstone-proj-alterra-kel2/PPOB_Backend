package request

import (
	"PPOB_BACKEND/businesses/products"

	"github.com/go-playground/validator"
)

type Product struct {
	Name                  string `json:"name" form:"name" validate:"required"`
	Category              string `json:"category" form:"category" validate:"required"`
	Description           string `json:"description" form:"description" validate:"required"`
	Price                 int    `json:"price" form:"price" validate:"required"`
	ProviderID            int    `json:"provider_id" form:"provider_id" validate:"required"`
	StockID               int    `json:"stock_id" form:"stock_id" validate:"required"`
	AdditionalInformation string `json:"additional_information" form:"additional_information"`
}

func (req *Product) ToDomain() *products.Domain {
	return &products.Domain{
		Name:                  req.Name,
		Category:              req.Category,
		Description:           req.Description,
		Price:                 req.Price,
		ProviderID:            req.ProviderID,
		StockID:               req.StockID,
		AdditionalInformation: req.AdditionalInformation,
	}
}

func (req *Product) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

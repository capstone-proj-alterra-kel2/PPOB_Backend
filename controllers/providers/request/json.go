package request

import (
	"PPOB_BACKEND/businesses/providers"

	"github.com/go-playground/validator"
)

type Provider struct {
	Name          string `json:"name" form:"name" validate:"required"`
	ProductTypeID int    `json:"product_type_id"`
}

func (req *Provider) ToDomain() *providers.Domain {
	return &providers.Domain{
		Name:          req.Name,
		ProductTypeID: req.ProductTypeID,
	}
}

func (req *Provider) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

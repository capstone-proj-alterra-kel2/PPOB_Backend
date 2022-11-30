package request

import (
	"PPOB_BACKEND/businesses/stocks"

	"github.com/go-playground/validator"
)

type Stock struct {
	Quantity int `json:"quantity" form:"quantity" validate:"required"`
	// Products []reqproduct.Product `json:"products"`
}

func (req *Stock) ToDomain() *stocks.Domain {
	// var products []products.Domain
	// for _, product := range req.Products {
	// 	products = append(products, *product.ToDomain())
	// }
	// from controllers

	return &stocks.Domain{
		Quantity: req.Quantity,
	}
}

func (req *Stock) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

package request

import (
	"PPOB_BACKEND/businesses/producttypes"

	"github.com/go-playground/validator"
)

type ProductType struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (req *ProductType) ToDomain() *producttypes.Domain {
	return &producttypes.Domain{
		Name: req.Name,
	}
}

func (req *ProductType) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

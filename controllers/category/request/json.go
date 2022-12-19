package request

import (
	"PPOB_BACKEND/businesses/category"

	"github.com/go-playground/validator"
)

type Category struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (req *Category) ToDomain() *category.Domain {
	return &category.Domain{
		Name: req.Name,
	}
}

func (req *Category) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

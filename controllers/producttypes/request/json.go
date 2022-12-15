package request

import (
	"PPOB_BACKEND/businesses/producttypes"

	"github.com/go-playground/validator"
)

type ProductType struct {
	Name       string `json:"name" form:"name" validate:"required"`
	Image      string `json:"image" form :"image" validate:"required"`
	CategoryID int    `json:"category_id" form:"category_id" validate:"required"`
}

func (req *ProductType) ToDomain() *producttypes.Domain {
	return &producttypes.Domain{
		Name:       req.Name,
		Image:      req.Image,
		CategoryID: req.CategoryID,
	}
}

func (req *ProductType) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

type UpdateProductType struct {
	Name       string `json:"name" form:"name" validate:"required"`
	CategoryID int    `json:"category_id" form:"category_id" validate:"required"`
	Image      string `json:"image" form :"image"`
}

func (req *UpdateProductType) ToDomain() *producttypes.Domain {
	return &producttypes.Domain{
		Name:       req.Name,
		Image:      req.Image,
		CategoryID: req.CategoryID,
	}
}

func (req *UpdateProductType) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

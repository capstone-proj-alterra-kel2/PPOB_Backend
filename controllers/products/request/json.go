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
	Stock                 int    `json:"stock" form:"stock" validate:"required"`
	Status                string `json:"status" form:"status" validate:"required"`
	AdditionalInformation string `json:"additional_information" form:"additional_information"`
	IsAvailable           bool   `json:"is_available" form:"is_available" validate:"required"`
	IsPromo               bool   `json:"is_promo" form:"is_promo" validate:"required"`
	IsPromoActive         bool   `json:"is_promo_active" form:"is_promo_active"`
	Discount              int    `json:"discount" form:"discount"`
	PromoStartDate        string `json:"promo_start_date" form:"promo_start_date"`
	PromoEndDate          string `json:"promo_end_date" form:"promo_end_date"`
}

func (req *Product) ToDomain() *products.Domain {
	return &products.Domain{
		Name:                  req.Name,
		Category:              req.Category,
		Description:           req.Description,
		Price:                 req.Price,
		ProviderID:            req.ProviderID,
		Stock:                 req.Stock,
		Status:                req.Status,
		AdditionalInformation: req.AdditionalInformation,
		IsAvailable:           req.IsAvailable,
		IsPromo:               req.IsPromo,
		IsPromoActive:         req.IsPromoActive,
		Discount:              req.Discount,
		PromoStartDate:        req.PromoStartDate,
		PromoEndDate:          req.PromoEndDate,
	}
}

func (req *Product) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}
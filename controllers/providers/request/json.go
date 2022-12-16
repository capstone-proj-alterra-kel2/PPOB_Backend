package request

import (
	"PPOB_BACKEND/businesses/providers"

	"github.com/go-playground/validator"
)

type InputProvider struct {
	ID            int                  `json:"id" form:"id" validate:"required"`
	Name          string               `json:"name" form:"name" validate:"required"`
	Image         string               `json:"image" form:"image" validate:"required"`
	Products      []UpdateCheckProduct `json:"products" form:"products" validate:"required"`
	ProductTypeID int                  `json:"product_type_id"`
}

type UpdateCheckProduct struct {
	ID                    int    `json:"id" form:"id" validate:"required"`
	Name                  string `json:"name" form:"name" validate:"required"`
	Description           string `json:"description" form:"description" validate:"required"`
	Price                 int    `json:"price" form:"price" validate:"required"`
	ProviderID            int    `json:"provider_id" form:"provider_id" validate:"required"`
	Stock                 *int   `json:"stock" form:"stock" validate:"required"`
	Status                string `json:"status" form:"status" validate:"required"`
	AdditionalInformation string `json:"additional_information" form:"additional_information"`
	IsAvailable           *bool  `json:"is_available" form:"is_available" validate:"required"`
	PriceStatus           string `json:"price_status" form:"price_status" validate:"required,oneof=normal promo"`
	IsPromo               *bool  `json:"is_promo" form:"is_promo"`
	IsPromoActive         *bool  `json:"is_promo_active" form:"is_promo_active"`
	Discount              *int   `json:"discount" form:"discount"`
	PromoStartDate        string `json:"promo_start_date" form:"promo_start_date"`
	PromoEndDate          string `json:"promo_end_date" form:"promo_end_date"`
}

func (req *InputProvider) ToDomain() *providers.ProviderDomain {

	var products []providers.UpdateDomain
	for _, product := range req.Products {
		products = append(products, *product.ToDomain())
	}

	return &providers.ProviderDomain{
		ID:            uint(req.ID),
		Name:          req.Name,
		Image:         req.Image,
		Products:      products,
		ProductTypeID: req.ProductTypeID,
	}
}

func (req *UpdateCheckProduct) ToDomain() *providers.UpdateDomain {

	return &providers.UpdateDomain{
		ID:                    uint(req.ID),
		Name:                  req.Name,
		Description:           req.Description,
		Price:                 req.Price,
		ProviderID:            req.ProviderID,
		Stock:                 req.Stock,
		Status:                req.Status,
		AdditionalInformation: req.AdditionalInformation,
		PriceStatus:           req.PriceStatus,
		IsAvailable:           req.IsAvailable,
		IsPromoActive:         req.IsPromoActive,
		Discount:              req.Discount,
		PromoStartDate:        req.PromoStartDate,
		PromoEndDate:          req.PromoEndDate,
	}
}

func (req *UpdateCheckProduct) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

type Provider struct {
	Name          string `json:"name" form:"name" validate:"required"`
	Image         string `json:"image" form:"image" validate:"required"`
	ProductTypeID int    `json:"product_type_id"`
}

func (req *Provider) ToDomain() *providers.Domain {
	return &providers.Domain{
		Name:          req.Name,
		Image:         req.Image,
		ProductTypeID: req.ProductTypeID,
	}
}

func (req *Provider) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

type UpdateData struct {
	Image string `json:"image" form:"image"`
	Name  string `json:"name" form:"name" validate:"required"`
}

func (req *UpdateData) ToDomain() *providers.Domain {
	return &providers.Domain{
		Name:  req.Name,
		Image: req.Image,
	}
}

func (req *UpdateData) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

type InputPhone struct {
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}

func (req *InputPhone) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

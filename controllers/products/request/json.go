package request

import (
	"PPOB_BACKEND/businesses/products"

	"github.com/go-playground/validator"
)

type Product struct {
	Name                  string `json:"name" form:"name" validate:"required"`
	Description           string `json:"description" form:"description" validate:"required"`
	Price                 int    `json:"price" form:"price" validate:"required"`
	ProviderID            int    `json:"provider_id" form:"provider_id" validate:"required"`
	Stock                 *int   `json:"stock" form:"stock" validate:"required"`
	Status                string `json:"status" form:"status" validate:"required"`
	TotalPurchased        int    `json:"total_purchased"`
	AdditionalInformation string `json:"additional_information" form:"additional_information"`
	PriceStatus           string `json:"price_status" form:"price_status" validate:"required,oneof=normal promo"`
	IsAvailable           *bool  `json:"is_available" form:"is_available"`
	IsPromoActive         *bool  `json:"is_promo_active" form:"is_promo_active"`
	Discount              *int   `json:"discount" form:"discount"`
	PromoStartDate        string `json:"promo_start_date" form:"promo_start_date"`
	PromoEndDate          string `json:"promo_end_date" form:"promo_end_date"`
}

func (req *Product) ToDomain() *products.Domain {
	return &products.Domain{
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

func (req *Product) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

type UpdateDataProduct struct {
	Name           string `json:"name" form:"name" validate:"required"`
	Description    string `json:"description" form:"description" validate:"required"`
	Price          int    `json:"price" form:"price" validate:"required"`
	ProviderID     int    `json:"provider_id" form:"provider_id" validate:"required"`
	Stock          *int   `json:"stock" form:"stock" validate:"required"`
	Status         string `json:"status" form:"status" validate:"required"`
	PriceStatus    string `json:"price_status" form:"price_status" validate:"required,oneof=normal promo"`
	IsAvailable    *bool  `json:"is_available" form:"is_available" validate:"required"`
	IsPromoActive  *bool  `json:"is_promo_active" form:"is_promo_active"`
	Discount       *int   `json:"discount" form:"discount"`
	PromoStartDate string `json:"promo_start_date" form:"promo_start_date"`
	PromoEndDate   string `json:"promo_end_date" form:"promo_end_date"`
}

func (req *UpdateDataProduct) ToDomain() *products.UpdateDataDomain {
	return &products.UpdateDataDomain{
		Name:           req.Name,
		Description:    req.Description,
		Price:          req.Price,
		ProviderID:     req.ProviderID,
		Stock:          req.Stock,
		Status:         req.Status,
		IsAvailable:    req.IsAvailable,
		PriceStatus:    req.PriceStatus,
		IsPromoActive:  req.IsPromoActive,
		Discount:       req.Discount,
		PromoStartDate: req.PromoStartDate,
		PromoEndDate:   req.PromoEndDate,
	}
}

func (req *UpdateDataProduct) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

type UpdateStockStatus struct {
	Stock          int    `json:"stock" form:"stock"`
	TotalPurchased int    `json:"total_purchased" form:"total_purchased"`
	Status         string `json:"status" form:"status"`
	IsAvailable    bool   `json:"is_available" form:"is_available"`
}

func (req *UpdateStockStatus) ToDomain() *products.UpdateStockStatusDomain {
	return &products.UpdateStockStatusDomain{
		Stock:          &req.Stock,
		TotalPurchased: req.TotalPurchased,
		Status:         req.Status,
		IsAvailable:    &req.IsAvailable,
	}
}

func (req *UpdateStockStatus) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

type UpdatePromoProduct struct {
	ID             uint   `json:"id"`
	IsAvailable    *bool  `json:"is_available" form:"is_available" validate:"required"`
	PriceStatus    string `json:"price_status" form:"price_status" validate:"required,oneof=normal promo"`
	IsPromoActive  *bool  `json:"is_promo_active" form:"is_promo_active"`
	Discount       *int   `json:"discount" form:"discount"`
	PromoStartDate string `json:"promo_start_date" form:"promo_start_date"`
	PromoEndDate   string `json:"promo_end_date" form:"promo_end_date"`
}

func (req *UpdatePromoProduct) ToDomain() *products.Domain {
	return &products.Domain{
		ID:             req.ID,
		IsAvailable:    req.IsAvailable,
		PriceStatus:    req.PriceStatus,
		IsPromoActive:  req.IsPromoActive,
		Discount:       req.Discount,
		PromoStartDate: req.PromoStartDate,
		PromoEndDate:   req.PromoEndDate,
	}
}

func (req *UpdatePromoProduct) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

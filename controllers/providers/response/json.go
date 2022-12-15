package response

import (
	"PPOB_BACKEND/businesses/products"
	"PPOB_BACKEND/businesses/providers"
	resproduct "PPOB_BACKEND/controllers/products/response"
	"time"

	"gorm.io/gorm"
)

type Provider struct {
	ID        uint                 `json:"id" gorm:"primaryKey"`
	Name      string               `json:"name"`
	Image     string               `json:"image"`
	Products  []resproduct.Product `json:"products"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	DeletedAt gorm.DeletedAt       `json:"deleted_at"`
}

func FromUpdateDomain(domain providers.ProviderDomain) Provider {
	var productsData []resproduct.Product
	productFromDomain := domain.Products

	if len(productFromDomain) != 0 {
		for _, product := range productFromDomain {
			productsData = append(productsData, resproduct.FromDomain(products.Domain{
				ID:                    product.ID,
				Name:                  product.Name,
				Description:           product.Description,
				Price:                 product.Price,
				ProviderID:            product.ProviderID,
				Stock:                 product.Stock,
				Status:                product.Status,
				TotalPurchased:        product.TotalPurchased,
				AdditionalInformation: product.AdditionalInformation,
				PriceStatus:           product.PriceStatus,
				IsAvailable:           product.IsAvailable,
				Discount:              product.Discount,
				PromoStartDate:        product.PromoStartDate,
				PromoEndDate:          product.PromoEndDate,
			}))
		}
	}

	return Provider{
		ID:        domain.ID,
		Name:      domain.Name,
		Image:     domain.Image,
		Products:  productsData,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdateAt,
		DeletedAt: domain.DeletedAt,
	}
}

func FromDomain(domain providers.Domain) Provider {
	var productsData []resproduct.Product
	productFromDomain := domain.Products

	if len(productFromDomain) != 0 {
		for _, product := range productFromDomain {
			productsData = append(productsData, resproduct.FromDomain(products.Domain{
				ID:                    product.ID,
				Name:                  product.Name,
				Description:           product.Description,
				Price:                 product.Price,
				ProviderID:            product.ProviderID,
				Stock:                 product.Stock,
				Status:                product.Status,
				TotalPurchased:        product.TotalPurchased,
				AdditionalInformation: product.AdditionalInformation,
				IsAvailable:           product.IsAvailable,
				PriceStatus:           product.PriceStatus,
				IsPromoActive:         product.IsPromoActive,
				Discount:              product.Discount,
				PromoStartDate:        product.PromoStartDate,
				PromoEndDate:          product.PromoEndDate,
				CreatedAt:             product.CreatedAt,
				UpdatedAt:             product.UpdatedAt,
			}))
		}
	}

	return Provider{
		ID:        domain.ID,
		Name:      domain.Name,
		Image:     domain.Image,
		Products:  productsData,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdateAt,
		DeletedAt: domain.DeletedAt,
	}
}

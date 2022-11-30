package providers

import (
	productdomain "PPOB_BACKEND/businesses/products"
	"PPOB_BACKEND/businesses/providers"
	"PPOB_BACKEND/drivers/postgresql/products"
	"time"

	"gorm.io/gorm"
)

type Provider struct {
	ID            uint               `json:"id" gorm:"size:100;primaryKey"`
	Name          string             `json:"name"`
	ProductTypeID int                `json:"product_type_id"`
	Products      []products.Product `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	DeletedAt     gorm.DeletedAt     `json:"deleted_at"`
}

func FromDomain(domain *providers.Domain) *Provider {
	var productData []products.Product
	productFromDomain := domain.Products

	for _, product := range productFromDomain {
		productData = append(productData, products.Product{
			ID:                    product.ID,
			Name:                  product.Name,
			Category:              product.Category,
			Description:           product.Description,
			Price:                 product.Price,
			ProviderID:            product.ProviderID,
			StockID:               product.StockID,
			TotalPurchased:        product.TotalPurchased,
			AdditionalInformation: product.AdditionalInformation,
			CreatedAt:             product.CreatedAt,
			UpdatedAt:             product.UpdatedAt,
			DeletedAt:             product.DeletedAt,
		})
	}

	return &Provider{
		ID:            domain.ID,
		Name:          domain.Name,
		ProductTypeID: domain.ProductTypeID,
		Products:      productData,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdateAt,
		DeletedAt:     domain.DeletedAt,
	}
}

func (recProvider *Provider) ToDomain() providers.Domain {
	var productFromDomain []productdomain.Domain
	for _, products := range recProvider.Products {
		productFromDomain = append(productFromDomain, products.ToDomain())
	}

	return providers.Domain{
		ID:            recProvider.ID,
		Name:          recProvider.Name,
		ProductTypeID: recProvider.ProductTypeID,
		Products:      productFromDomain,
		CreatedAt:     recProvider.CreatedAt,
		UpdateAt:      recProvider.UpdatedAt,
		DeletedAt:     recProvider.DeletedAt,
	}
}

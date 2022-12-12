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
	Image         string             `json:"image"`
	ProductTypeID int                `json:"product_type_id"`
	Products      []products.Product `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	DeletedAt     gorm.DeletedAt     `json:"deleted_at"`
}

func FromDomainUpdate(domain *providers.ProviderDomain) *Provider {
	var productData []products.Product
	productFromDomain := domain.Products

	for _, product := range productFromDomain {
		productData = append(productData, products.Product{
			ID:                    product.ID,
			Name:                  product.Name,
			Category:              product.Category,
			Description:           product.Description,
			Price:                 product.Price,
			Stock:                 product.Stock,
			Status:                product.Status,
			ProviderID:            product.ProviderID,
			TotalPurchased:        product.TotalPurchased,
			AdditionalInformation: product.AdditionalInformation,
			IsAvailable:           product.IsAvailable,
			IsPromo:               product.IsPromo,
			IsPromoActive:         product.IsPromoActive,
			Discount:              product.Discount,
			PromoStartDate:        product.PromoStartDate,
			PromoEndDate:          product.PromoEndDate,
			CreatedAt:             product.CreatedAt,
			UpdatedAt:             product.UpdatedAt,
			DeletedAt:             product.DeletedAt,
		})
	}

	return &Provider{
		ID:            domain.ID,
		Name:          domain.Name,
		Image:         domain.Image,
		ProductTypeID: domain.ProductTypeID,
		Products:      productData,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdateAt,
		DeletedAt:     domain.DeletedAt,
	}
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
			Stock:                 product.Stock,
			Status:                product.Status,
			ProviderID:            product.ProviderID,
			TotalPurchased:        product.TotalPurchased,
			AdditionalInformation: product.AdditionalInformation,
			IsAvailable:           product.IsAvailable,
			IsPromo:               product.IsPromo,
			IsPromoActive:         product.IsPromoActive,
			Discount:              product.Discount,
			PromoStartDate:        product.PromoStartDate,
			PromoEndDate:          product.PromoEndDate,
			CreatedAt:             product.CreatedAt,
			UpdatedAt:             product.UpdatedAt,
			DeletedAt:             product.DeletedAt,
		})
	}

	return &Provider{
		ID:            domain.ID,
		Name:          domain.Name,
		Image:         domain.Image,
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
		Image:         recProvider.Image,
		ProductTypeID: recProvider.ProductTypeID,
		Products:      productFromDomain,
		CreatedAt:     recProvider.CreatedAt,
		UpdateAt:      recProvider.UpdatedAt,
		DeletedAt:     recProvider.DeletedAt,
	}
}

func (recProvider *Provider) ToDomainUpdate() providers.ProviderDomain {
	var products []providers.UpdateDomain
	for _, product := range recProvider.Products {
		products = append(products, providers.UpdateDomain(product.ToDomain()))
	}

	return providers.ProviderDomain{
		ID:            uint(recProvider.ID),
		Name:          recProvider.Name,
		Image:         recProvider.Image,
		Products:      products,
		ProductTypeID: recProvider.ProductTypeID,
		CreatedAt:     recProvider.CreatedAt,
		UpdateAt:      recProvider.UpdatedAt,
		DeletedAt:     recProvider.DeletedAt,
	}
}

type CheckProductType struct {
	ProductTypeID int
}

func FromChekcProductTypeDomain(domain *providers.CheckProductTypeDomain) *CheckProductType {
	return &CheckProductType{
		ProductTypeID: domain.ProductTypeID,
	}
}

func (recProvider *CheckProductType) ToDomain() providers.CheckProductTypeDomain {
	return providers.CheckProductTypeDomain{
		ProductTypeID: recProvider.ProductTypeID,
	}
}

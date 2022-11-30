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
	Products  []resproduct.Product `json:"products"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	DeletedAt gorm.DeletedAt       `json:"deleted_at"`
}

func FromDomain(domain providers.Domain) Provider {
	var productsData []resproduct.Product
	productFromDomain := domain.Products

	if len(productFromDomain) != 0 {
		for _, product := range productFromDomain {
			productsData = append(productsData, resproduct.FromDomain(products.Domain{
				ID:                    product.ID,
				Name:                  product.Name,
				Category:              product.Category,
				Description:           product.Description,
				Price:                 product.Price,
				ProviderID:            product.ProviderID,
				StockID:               product.StockID,
				TotalPurchased:        product.TotalPurchased,
				AdditionalInformation: product.AdditionalInformation,
			}))
		}
	}

	return Provider{
		ID:        domain.ID,
		Name:      domain.Name,
		Products:  productsData,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdateAt,
		DeletedAt: domain.DeletedAt,
	}
}

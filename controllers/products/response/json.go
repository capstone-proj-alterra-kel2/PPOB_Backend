package response

import (
	"PPOB_BACKEND/businesses/products"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                    uint           `json:"id" gorm:"size:100;primaryKey"`
	Name                  string         `json:"name"`
	Category              string         `json:"category"`
	Description           string         `json:"description"`
	Price                 int            `json:"price"`
	ProviderID            int            `json:"provider_id"`
	StockID               int            `json:"stock_id"`
	TotalPurchased        int            `json:"total_purchased"`
	AdditionalInformation string         `json:"additional_information"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain products.Domain) Product {
	return Product{
		ID:                    domain.ID,
		Name:                  domain.Name,
		Category:              domain.Category,
		Description:           domain.Description,
		Price:                 domain.Price,
		ProviderID:            domain.ProviderID,
		StockID:               domain.StockID,
		TotalPurchased:        domain.TotalPurchased,
		AdditionalInformation: domain.AdditionalInformation,
		CreatedAt:             domain.CreatedAt,
		UpdatedAt:             domain.UpdatedAt,
		DeletedAt:             domain.DeletedAt,
	}
}

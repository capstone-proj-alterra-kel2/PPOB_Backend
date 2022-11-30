package products

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

func FromDomain(domain *products.Domain) *Product {

	return &Product{
		ID:                    domain.ID,
		Name:                  domain.Name,
		Category:              domain.Category,
		Price:                 domain.Price,
		Description:           domain.Description,
		ProviderID:            domain.ProviderID,
		StockID:               domain.StockID,
		TotalPurchased:        domain.TotalPurchased,
		AdditionalInformation: domain.AdditionalInformation,
		CreatedAt:             domain.CreatedAt,
		UpdatedAt:             domain.UpdatedAt,
		DeletedAt:             domain.DeletedAt,
	}
}

func (recProd *Product) ToDomain() products.Domain {
	return products.Domain{
		ID:                    recProd.ID,
		Name:                  recProd.Name,
		Category:              recProd.Category,
		Price:                 recProd.Price,
		Description:           recProd.Description,
		ProviderID:            recProd.ProviderID,
		StockID:               recProd.StockID,
		AdditionalInformation: recProd.AdditionalInformation,
		CreatedAt:             recProd.CreatedAt,
		UpdatedAt:             recProd.UpdatedAt,
		DeletedAt:             recProd.DeletedAt,
	}
}

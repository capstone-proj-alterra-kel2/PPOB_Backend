package products

import (
	"PPOB_BACKEND/businesses/products"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                    uint           `json:"id" gorm:"size:100;primaryKey"`
	Name                  string         `json:"name"`
	Description           string         `json:"description"`
	Price                 int            `json:"price"`
	ProviderID            int            `json:"provider_id"`
	Stock                 *int           `json:"stock"`
	Status                string         `json:"status"`
	TotalPurchased        int            `json:"total_purchased"`
	AdditionalInformation string         `json:"additional_information"`
	IsAvailable           *bool          `json:"is_available"`
	IsPromo               *bool          `json:"is_promo"`
	IsPromoActive         *bool          `json:"is_promo_active" gorm:"default:false"`
	Discount              int            `json:"discount"`
	PromoStartDate        string         `json:"promo_start_date"`
	PromoEndDate          string         `json:"promo_end_date"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *products.Domain) *Product {

	return &Product{
		ID:                    domain.ID,
		Name:                  domain.Name,
		Price:                 domain.Price,
		Description:           domain.Description,
		ProviderID:            domain.ProviderID,
		Stock:                 domain.Stock,
		Status:                domain.Status,
		TotalPurchased:        domain.TotalPurchased,
		AdditionalInformation: domain.AdditionalInformation,
		IsAvailable:           domain.IsAvailable,
		IsPromo:               domain.IsPromo,
		IsPromoActive:         domain.IsPromoActive,
		Discount:              domain.Discount,
		PromoStartDate:        domain.PromoStartDate,
		PromoEndDate:          domain.PromoEndDate,
		CreatedAt:             domain.CreatedAt,
		UpdatedAt:             domain.UpdatedAt,
		DeletedAt:             domain.DeletedAt,
	}
}

func (recProd *Product) ToDomain() products.Domain {
	return products.Domain{
		ID:                    recProd.ID,
		Name:                  recProd.Name,
		Price:                 recProd.Price,
		Description:           recProd.Description,
		ProviderID:            recProd.ProviderID,
		Stock:                 recProd.Stock,
		Status:                recProd.Status,
		TotalPurchased:        recProd.TotalPurchased,
		AdditionalInformation: recProd.AdditionalInformation,
		IsAvailable:           recProd.IsAvailable,
		IsPromo:               recProd.IsPromo,
		IsPromoActive:         recProd.IsPromoActive,
		Discount:              recProd.Discount,
		PromoStartDate:        recProd.PromoStartDate,
		PromoEndDate:          recProd.PromoEndDate,
		CreatedAt:             recProd.CreatedAt,
		UpdatedAt:             recProd.UpdatedAt,
		DeletedAt:             recProd.DeletedAt,
	}
}

func FromUpdatedDomain(domain *products.UpdateDataDomain) *Product {

	return &Product{
		Name:                  domain.Name,
		Price:                 domain.Price,
		Description:           domain.Description,
		ProviderID:            domain.ProviderID,
		Stock:                 domain.Stock,
		Status:                domain.Status,
		AdditionalInformation: domain.AdditionalInformation,
		IsAvailable:           domain.IsAvailable,
		IsPromo:               domain.IsPromo,
		IsPromoActive:         domain.IsPromoActive,
		Discount:              domain.Discount,
		PromoStartDate:        domain.PromoStartDate,
		PromoEndDate:          domain.PromoEndDate,
	}
}

func FromUpdatedStockStatusDomain(domain *products.UpdateStockStatusDomain) *Product {

	return &Product{
		Stock:          domain.Stock,
		Status:         domain.Status,
		TotalPurchased: domain.TotalPurchased,
		IsAvailable:    domain.IsAvailable,
	}
}

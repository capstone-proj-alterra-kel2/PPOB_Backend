package response

import (
	"PPOB_BACKEND/businesses/producttypes"
	"PPOB_BACKEND/businesses/providers"
	resprovider "PPOB_BACKEND/controllers/providers/response"
	"time"

	"gorm.io/gorm"
)

type ProductType struct {
	ID         uint                   `json:"id" gorm:"primaryKey"`
	Name       string                 `json:"name"`
	CategoryID int                    `json:"category_id"`
	Providers  []resprovider.Provider `json:"providers"`
	Image      string                 `json:"image"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
	DeletedAt  gorm.DeletedAt         `json:"deleted_at"`
}

func FromDomain(domain producttypes.Domain) ProductType {
	var providerData []resprovider.Provider
	providerFromDomain := domain.Providers

	for _, provider := range providerFromDomain {
		providerData = append(providerData, resprovider.FromDomain(providers.Domain{
			ID:        provider.ID,
			Name:      provider.Name,
			Image:     provider.Image,
			CreatedAt: provider.CreatedAt,
			UpdateAt:  provider.UpdateAt,
			DeletedAt: provider.DeletedAt,
		}))
	}

	return ProductType{
		ID:         domain.ID,
		Providers:  providerData,
		Name:       domain.Name,
		CategoryID: domain.CategoryID,
		Image:      domain.Image,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdateAt,
		DeletedAt:  domain.DeletedAt,
	}
}

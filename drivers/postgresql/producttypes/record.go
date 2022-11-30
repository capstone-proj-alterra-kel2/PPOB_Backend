package producttypes

import (
	"PPOB_BACKEND/businesses/producttypes"
	providerdomain "PPOB_BACKEND/businesses/providers"
	"PPOB_BACKEND/drivers/postgresql/providers"
	"time"

	"gorm.io/gorm"
)

type ProductType struct {
	ID        uint                 `json:"id" gorm:"size:100;primaryKey"`
	Name      string               `json:"name"`
	Providers []providers.Provider `json:"providers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	DeletedAt gorm.DeletedAt       `json:"deleted_at"`
}

func FromDomain(domain *producttypes.Domain) *ProductType {

	var providersData []providers.Provider
	providerFromDomain := domain.Providers

	for _, provider := range providerFromDomain {
		providersData = append(providersData, providers.Provider{
			ID:        provider.ID,
			Name:      provider.Name,
			CreatedAt: provider.CreatedAt,
			UpdatedAt: provider.UpdateAt,
			DeletedAt: provider.DeletedAt,
		})
	}

	return &ProductType{
		ID:        domain.ID,
		Name:      domain.Name,
		Providers: providersData,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdateAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (recProdType *ProductType) ToDomain() producttypes.Domain {
	var providersFromDomain []providerdomain.Domain
	for _, providerData := range recProdType.Providers {
		providersFromDomain = append(providersFromDomain, providerData.ToDomain())
	}

	return producttypes.Domain{
		ID:        recProdType.ID,
		Name:      recProdType.Name,
		Providers: providersFromDomain,
		CreatedAt: recProdType.CreatedAt,
		UpdateAt:  recProdType.UpdatedAt,
		DeletedAt: recProdType.DeletedAt,
	}
}

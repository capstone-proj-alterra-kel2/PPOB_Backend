package providers

import (
	"PPOB_BACKEND/businesses/products"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID            uint
	Name          string
	Image         string
	ProductTypeID int
	Products      []products.Domain
	CreatedAt     time.Time
	UpdateAt      time.Time
	DeletedAt     gorm.DeletedAt
}

type ProviderDomain struct {
	ID            uint
	Name          string
	Image         string
	ProductTypeID int
	Products      []UpdateDomain
	CreatedAt     time.Time
	UpdateAt      time.Time
	DeletedAt     gorm.DeletedAt
}

type UpdateDomain struct {
	ID                    uint
	Name                  string
	Category              string
	Description           string
	Price                 int
	ProviderID            int
	Stock                 int
	Status                string
	TotalPurchased        int
	AdditionalInformation string
	IsAvailable           bool
	IsPromo               bool
	IsPromoActive         bool
	Discount              int
	PromoStartDate        string
	PromoEndDate          string
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}

type Prefix struct {
	Prefix   string `json:"prefix"`
	Provider string `json:"provider"`
	Type     string `json:"type"`
}

type Usecase interface {
	GetAll() []Domain
	Create(providerDomain *Domain) Domain
	GetOne(provider_id int) Domain
	GetByPhone(phone_number string, product_type_id int) Domain
	Update(providerDomain *Domain, provider_id int) Domain
	UpdateCheck(providerDomain *ProviderDomain, provider_id int) Domain
	Delete(provider_id int) Domain
}

type Repository interface {
	GetAll() []Domain
	Create(providerDomain *Domain) Domain
	GetOne(provider_id int) Domain
	GetByPhone(provider string, product_type_id int) Domain
	Update(providerDomain *Domain, provider_id int) Domain
	UpdateCheck(providerDomain *ProviderDomain, provider_id int) Domain
	Delete(provider_id int) Domain
}

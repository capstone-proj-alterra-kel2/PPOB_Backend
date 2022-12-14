package products

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID                    uint
	Name                  string
	Description           string
	Price                 int
	ProviderID            int
	ProductTypeID         int
	Stock                 *int
	Status                string
	TotalPurchased        int
	AdditionalInformation string
	IsAvailable           *bool
	PriceStatus           string
	IsPromoActive         *bool
	Discount              *int
	PromoStartDate        string
	PromoEndDate          string
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}

type UpdateDataDomain struct {
	Name                  string
	Description           string
	Price                 int
	ProviderID            int
	ProductTypeID         int
	Stock                 *int
	Status                string
	AdditionalInformation string
	IsAvailable           *bool
	PriceStatus           string
	IsPromoActive         *bool
	Discount              *int
	PromoStartDate        string
	PromoEndDate          string
}

type UpdateStockStatusDomain struct {
	Stock          *int
	TotalPurchased int
	Status         string
	IsAvailable    *bool
}

type Usecase interface {
	GetAll(Page int, Size int, Sort string, Search string) (*gorm.DB, []Domain)
	GetAllForUser() []Domain
	Create(productDomain *Domain) (Domain, bool, bool)
	GetOne(product_id int) (Domain, error)
	UpdateData(productDomain *UpdateDataDomain, product_id int) (Domain, error, bool, bool)
	UpdatePromo(productDomain *Domain) Domain
	UpdateStockStatus(productDomain *UpdateStockStatusDomain, product_id int) Domain
	Delete(product_id int) (Domain, error)
}

type Repository interface {
	GetAll(Page int, Size int, Sort string, Search string) (*gorm.DB, []Domain)
	GetAllForUser() []Domain
	Create(productDomain *Domain) (Domain, bool)
	GetOne(product_id int) (Domain, error)
	UpdateData(productDomain *UpdateDataDomain, product_id int) (Domain, error, bool)
	UpdatePromo(productDomain *Domain) Domain
	UpdateStockStatus(productDomain *UpdateStockStatusDomain, product_id int) Domain
	Delete(product_id int) (Domain, error)
}

package products

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID                    uint
	Name                  string
	Category              string
	Description           string
	Price                 int
	ProviderID            int
	Stock                 *int
	Status                string
	TotalPurchased        int
	AdditionalInformation string
	IsAvailable           *bool
	IsPromo               *bool
	IsPromoActive         *bool
	Discount              int
	PromoStartDate        string
	PromoEndDate          string
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}

type UpdateDataDomain struct {
	Name                  string
	Category              string
	Description           string
	Price                 int
	ProviderID            int
	Stock                 *int
	Status                string
	AdditionalInformation string
	IsAvailable           *bool
	IsPromo               *bool
	IsPromoActive         *bool
	Discount              int
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
	Create(productDomain *Domain) Domain
	GetOne(product_id int) Domain
	UpdateData(productDomain *UpdateDataDomain, product_id int) (Domain, error)
	UpdateStockStatus(productDomain *UpdateStockStatusDomain, product_id int) Domain
	Delete(product_id int) Domain
}

type Repository interface {
	GetAll(Page int, Size int, Sort string, Search string) (*gorm.DB, []Domain)
	Create(productDomain *Domain) Domain
	GetOne(product_id int) Domain
	UpdateData(productDomain *UpdateDataDomain, product_id int) (Domain, error)
	UpdateStockStatus(productDomain *UpdateStockStatusDomain, product_id int) Domain
	Delete(product_id int) Domain
}

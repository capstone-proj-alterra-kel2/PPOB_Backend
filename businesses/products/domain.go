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
	StockID               int
	TotalPurchased        int
	AdditionalInformation string
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}

type Usecase interface {
	GetAll() []Domain
	Create(productDomain *Domain) Domain
	GetOne(product_id int) Domain
	Update(productDomain *Domain, product_id int) Domain
	Delete(product_id int) Domain
}

type Repository interface {
	GetAll() []Domain
	Create(productDomain *Domain) Domain
	GetOne(product_id int) Domain
	Update(productDomain *Domain, product_id int) Domain
	Delete(product_id int) Domain
}

package providers

import (
	"PPOB_BACKEND/businesses/products"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID            uint
	Name          string
	ProductTypeID int
	Products      []products.Domain
	CreatedAt     time.Time
	UpdateAt      time.Time
	DeletedAt     gorm.DeletedAt
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
	Delete(provider_id int) Domain
}

type Repository interface {
	GetAll() []Domain
	Create(providerDomain *Domain) Domain
	GetOne(provider_id int) Domain
	GetByPhone(provider string, product_type_id int) Domain
	Update(providerDomain *Domain, provider_id int) Domain
	Delete(provider_id int) Domain
}

package producttypes

import (
	"PPOB_BACKEND/businesses/providers"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	Name      string
	Providers []providers.Domain
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	GetAll() []Domain
	Create(productTypeDomain *Domain) Domain
	GetOne(product_type_id int) Domain
	Update(productTypeDomain *Domain, product_type_id int) Domain
	Delete(product_type_id int) Domain
}

type Repository interface {
	GetAll() []Domain
	Create(productTypeDomain *Domain) Domain
	GetOne(product_type_id int) Domain
	Update(productTypeDomain *Domain, product_type_id int) Domain
	Delete(product_type_id int) Domain
}

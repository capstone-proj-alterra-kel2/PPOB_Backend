package category

import (
	"PPOB_BACKEND/businesses/producttypes"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID          uint
	Name        string
	ProductType []producttypes.Domain
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeeletedAt  gorm.DeletedAt
}

type Usecase interface {
	GetAll(ProductType string) []Domain
	Create(categoryDomain *Domain) Domain
	GetDetail(category_id int) Domain
	Update(categoryDomain *Domain, category_id int) Domain
	Delete(category_id int) (Domain, error)
}

type Repository interface {
	GetAll(ProductType string) []Domain
	Create(categoryDomain *Domain) Domain
	GetDetail(category_id int) Domain
	Update(categoryDomain *Domain, category_id int) Domain
	Delete(category_id int) (Domain, error)
}

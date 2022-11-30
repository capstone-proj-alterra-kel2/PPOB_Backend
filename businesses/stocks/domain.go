package stocks

import (
	"PPOB_BACKEND/businesses/products"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	Quantity  int
	Products  []products.Domain
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	GetAll() []Domain
	Get(stock_id int) Domain
	Create(stockDomain *Domain) Domain
	Update(stockDomain *Domain, stock_id int) Domain
	Delete(stock_id int) Domain
}

type Repository interface {
	GetAll() []Domain
	Get(stock_id int) Domain
	Create(stockDomain *Domain) Domain
	Update(stockDomain *Domain, stock_id int) Domain
	Delete(stock_id int) Domain
}

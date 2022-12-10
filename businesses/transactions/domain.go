package transactions

import (
	"PPOB_BACKEND/businesses/products"
	"PPOB_BACKEND/businesses/users"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID              uint
	ProductID       int
	ProductName     string
	UserID          int
	UserEmail       string
	ProductPrice    int
	ProductDiscount int
	Status          string
	AdminFee        int
	TotalPrice      int
	TransactionDate time.Time
	CreatedAt       time.Time
	UpdateAt        time.Time
	DeletedAt       gorm.DeletedAt
}

type Usecase interface {
	GetAll(Page int, Size int, Sort string, Search string) ([]Domain, *gorm.DB)
	GetDetail(transaction_id int) (Domain, bool)
	GetTransactionHistory(user_id int) []Domain
	Create(productDomain *products.Domain, userDomain *users.Domain, totalAmount int, productDiscount int) Domain
	Delete(transaction_id int) (Domain, bool)
}

type Repository interface {
	GetAll(Page int, Size int, Sort string, Search string) ([]Domain, *gorm.DB)
	GetDetail(transaction_id int) (Domain, bool)
	GetTransactionHistory(user_id int) []Domain
	Create(domain *Domain) Domain
	Delete(transaction_id int) (Domain, bool)
}

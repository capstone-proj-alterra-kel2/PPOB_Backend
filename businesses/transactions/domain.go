package transactions

import (
	"PPOB_BACKEND/businesses/products"
	"PPOB_BACKEND/businesses/producttypes"
	"PPOB_BACKEND/businesses/users"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID                uint
	ProductID         int
	ProductName       string
	ProductTypeID     int
	ProductType       string
	UserID            int
	UserEmail         string
	TargetPhoneNumber string
	ProductPrice      int
	ProductDiscount   int
	Status            string
	AdminFee          int
	TotalPrice        int
	TransactionDate   time.Time
	CreatedAt         time.Time
	UpdateAt          time.Time
	DeletedAt         gorm.DeletedAt
}

type Usecase interface {
	GetAll(Page int, Size int, Sort string, Search string) ([]Domain, *gorm.DB)
	GetDetail(transaction_id int) (Domain, bool)
	GetTransactionHistory(user_id int) []Domain
	Update(transactionDomain *Domain, transaction_id int) (Domain, bool)
	Create(productDomain *products.Domain, userDomain *users.Domain, productTypeDomain *producttypes.Domain, totalAmount int, productDiscount int, targetPhoneNumber string) Domain
	Delete(transaction_id int) (Domain, bool)
}

type Repository interface {
	GetAll(Page int, Size int, Sort string, Search string) ([]Domain, *gorm.DB)
	GetDetail(transaction_id int) (Domain, bool)
	GetTransactionHistory(user_id int) []Domain
	Update(transactionDomain *Domain, transaction_id int) (Domain, bool)
	Create(domain *Domain) Domain
	Delete(transaction_id int) (Domain, bool)
}

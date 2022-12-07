package transactions

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID              uint
	ProductID       int
	UserID          int
	Amount          int
	Status          string
	TransactionDate time.Time
	CreatedAt       time.Time
	UpdateAt        time.Time
	DeletedAt       gorm.DeletedAt
}

type Usecase interface {
	GetAll() ([]Domain, error)
	GetDetail(transaction_id int) Domain
	Create(transactionDomain *Domain) Domain
	Delete(transaction_id int) Domain
}

type Repository interface {
	GetAll(product_type_id int) ([]Domain, error)
	GetDetail(transaction_id int) Domain
	Create(transactionDomain *Domain) Domain
	Delete(transaction_id int) Domain
}

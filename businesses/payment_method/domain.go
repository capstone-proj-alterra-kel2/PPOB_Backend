package payment_method

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           uint
	Payment_Name string
	Url_Payment  string
	Icon         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Usecase interface {
	GetAll() []Domain
	GetSpecificPayment(id string) Domain
	CreatePayment(paymentDomain *Domain) Domain
	UpdatePaymentByID(id string, paymentDomain *Domain) Domain
	DeletePayment(id string) bool
}

type Repository interface {
	GetAll() []Domain
	GetSpecificPayment(id string) Domain
	CreatePayment(paymentDomain *Domain) Domain
	UpdatePaymentByID(id string, paymentDomain *Domain) Domain
	DeletePayment(id string) bool
}

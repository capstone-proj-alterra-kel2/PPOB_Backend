package transactions

import (
	"PPOB_BACKEND/businesses/transactions"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID                uint           `json:"id" gorm:"size:100;primaryKey"`
	ProductID         int            `json:"product_id"`
	ProductTypeID     int            `json:"product_type_id"`
	UserID            int            `json:"user_id"`
	UserEmail         string         `json:"user_email"`
	TargetPhoneNumber string         `json:"target_phone_number"`
	ProductPrice      int            `json:"product_price"`
	ProductDiscount   int            `json:"product_discount"`
	ProductType       string         `json:"product_type"`
	ProductName       string         `json:"product_name"`
	AdminFee          int            `json:"admin_fee"`
	TotalPrice        int            `json:"total_price"`
	TransactionDate   time.Time      `json:"transaction_date"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *transactions.Domain) *Transaction {
	return &Transaction{
		ID:                domain.ID,
		ProductID:         domain.ProductID,
		ProductTypeID:     domain.ProductTypeID,
		UserID:            domain.UserID,
		UserEmail:         domain.UserEmail,
		TargetPhoneNumber: domain.TargetPhoneNumber,
		ProductPrice:      domain.ProductPrice,
		ProductName:       domain.ProductName,
		ProductType:       domain.ProductType,
		ProductDiscount:   domain.ProductDiscount,
		AdminFee:          domain.AdminFee,
		TotalPrice:        domain.TotalPrice,
		TransactionDate:   domain.TransactionDate,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdateAt,
		DeletedAt:         domain.DeletedAt,
	}
}

func (recTrans *Transaction) ToDomain() transactions.Domain {
	return transactions.Domain{
		ID:                recTrans.ID,
		ProductID:         recTrans.ProductID,
		ProductTypeID:     recTrans.ProductTypeID,
		UserID:            recTrans.UserID,
		UserEmail:         recTrans.UserEmail,
		TargetPhoneNumber: recTrans.TargetPhoneNumber,
		ProductPrice:      recTrans.ProductPrice,
		ProductName:       recTrans.ProductName,
		ProductType:       recTrans.ProductType,
		ProductDiscount:   recTrans.ProductDiscount,
		AdminFee:          recTrans.AdminFee,
		TotalPrice:        recTrans.TotalPrice,
		TransactionDate:   recTrans.TransactionDate,
		CreatedAt:         recTrans.CreatedAt,
		UpdateAt:          recTrans.UpdatedAt,
		DeletedAt:         recTrans.DeletedAt,
	}
}

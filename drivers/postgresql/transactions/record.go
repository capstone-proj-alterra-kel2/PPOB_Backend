package transactions

import (
	"PPOB_BACKEND/businesses/transactions"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID              uint           `json:"id" gorm:"size:100;primaryKey"`
	ProductID       int            `json:"product_id"`
	UserID          int            `json:"user_id"`
	Amount          int            `json:"amount"`
	Status          string         `json:"status"`
	TransactionDate time.Time      `json:"transaction_date"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *transactions.Domain) *Transaction {
	return &Transaction{
		ID:              domain.ID,
		ProductID:       domain.ProductID,
		UserID:          domain.UserID,
		Amount:          domain.Amount,
		Status:          domain.Status,
		TransactionDate: domain.TransactionDate,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdateAt,
		DeletedAt:       domain.DeletedAt,
	}
}

func (recTrans *Transaction) ToDomain() transactions.Domain {
	return transactions.Domain{
		ID:              recTrans.ID,
		ProductID:       recTrans.ProductID,
		UserID:          recTrans.UserID,
		Amount:          recTrans.Amount,
		Status:          recTrans.Status,
		TransactionDate: recTrans.TransactionDate,
		CreatedAt:       recTrans.CreatedAt,
		UpdateAt:        recTrans.UpdatedAt,
		DeletedAt:       recTrans.DeletedAt,
	}
}

package response

import (
	"PPOB_BACKEND/businesses/transactions"
	"time"
)

type Transaction struct {
	ID              uint      `json:"id"`
	ProductName     string    `json:"product_name"`
	UserEmail       string    `json:"user_email"`
	ProductPrice    int       `json:"product_price"`
	AdminFee        int       `json:"admin_fee"`
	TotalPrice      int       `json:"total_price"`
	TransactionDate time.Time `json:"transaction_date"`
}

func FromDomain(domain transactions.Domain) Transaction {
	return Transaction{
		ID:              domain.ID,
		ProductName:     domain.ProductName,
		UserEmail:       domain.UserEmail,
		ProductPrice:    domain.ProductPrice,
		AdminFee:        domain.AdminFee,
		TotalPrice:      domain.TotalPrice,
		TransactionDate: domain.TransactionDate,
	}
}

package response

import (
	"PPOB_BACKEND/businesses/wallet_histories"
	"time"

	"gorm.io/gorm"
)

type WalletHistory struct {
	HistoryWalletID uint           `json:"history_wallet_id;primaryKey"`
	NoWallet        string         `json:"no_wallet" gorm:"size:16"`
	CashIn          int            `json:"cash_in"`
	CashOut         int            `json:"cash_out"`
	Description     string         `json:"description"`
	DateWallet      time.Time      `json:"date_wallet"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

type CashInCashOut struct {
	CashIn  int    `json:"cash_in"`
	CashOut int    `json:"cash_out"`
	Month   string `json:"month"`
}

func FromDomainCashInCashOut(domain wallet_histories.CashInCashOutMonthlyDomain) CashInCashOut {
	return CashInCashOut{
		CashIn:  domain.CashIn,
		CashOut: domain.CashOut,
		Month:   domain.Month,
	}
}

func FromDomain(domain wallet_histories.Domain) WalletHistory {
	return WalletHistory{
		HistoryWalletID: domain.HistoryWalletID,
		NoWallet:        domain.NoWallet,
		CashIn:          domain.CashIn,
		CashOut:         domain.CashOut,
		Description:     domain.Description,
		DateWallet:      domain.DateWallet,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
		DeletedAt:       domain.DeletedAt,
	}
}

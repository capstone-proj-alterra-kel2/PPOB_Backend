package response

import (
	"PPOB_BACKEND/businesses/wallet_histories"
	"time"

	"gorm.io/gorm"
)

type WalletHistory struct {
	HistoryWalletID uint           `json:"history_wallet_id;primaryKey"`
	NoWallet        string         `json:"no_wallet" gorm:"size:16"`
	Income          int            `json:"income"`
	Outcome         int            `json:"outcome"`
	Description     string         `json:"description"`
	DateWallet      time.Time      `json:"date_wallet"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

type OutcomeIncome struct {
	Outcome int `json:"income"`
	Income  int `json:"outcome"`
}

func FromDomainOutcomeIncome(domain wallet_histories.OutcomeIncomeMonthlyDomain) OutcomeIncome {
	return OutcomeIncome{
		Outcome: domain.Outcome,
		Income:  domain.Income,
	}
}

func FromDomain(domain wallet_histories.Domain) WalletHistory {
	return WalletHistory{
		HistoryWalletID: domain.HistoryWalletID,
		NoWallet:        domain.NoWallet,
		Income:          domain.Income,
		Outcome:         domain.Outcome,
		Description:     domain.Description,
		DateWallet:      domain.DateWallet,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
		DeletedAt:       domain.DeletedAt,
	}
}

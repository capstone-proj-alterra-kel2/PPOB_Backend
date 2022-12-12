package wallet_histories

import (
	"PPOB_BACKEND/businesses/wallet_histories"
	"time"

	"gorm.io/gorm"
)

type WalletHistory struct {
	HistoryWalletID uint           `json:"history_wallet_id" gorm:"primaryKey"`
	NoWallet        string         `json:"no_wallet" gorm:"size:16;"`
	Income          int            `json:"income"`
	Outcome         int            `json:"outcome"`
	Description     string         `json:"description"`
	DateWallet      time.Time      `json:"date_wallet"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *wallet_histories.Domain) *WalletHistory {
	return &WalletHistory{
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

func (rec *WalletHistory) ToDomain() wallet_histories.Domain {
	return wallet_histories.Domain{
		HistoryWalletID: rec.HistoryWalletID,
		NoWallet:        rec.NoWallet,
		Income:          rec.Income,
		Outcome:         rec.Outcome,
		Description:     rec.Description,
		DateWallet:      rec.DateWallet,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
		DeletedAt:       rec.DeletedAt,
	}
}

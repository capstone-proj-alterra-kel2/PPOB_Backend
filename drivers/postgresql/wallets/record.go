package wallets

import (
	busWalletHistory "PPOB_BACKEND/businesses/wallet_histories"

	sqlWalletHistory "PPOB_BACKEND/drivers/postgresql/wallet_histories"

	"PPOB_BACKEND/businesses/wallets"

	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	NoWallet        string                           `json:"no_wallet" gorm:"size:16;primaryKey" `
	UserID          uint                             `json:"user_id"`
	Balance         int                              `json:"balance"`
	HistoriesWallet []sqlWalletHistory.WalletHistory `json:"histories_wallet" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:no_wallet"`
	CreatedAt       time.Time                        `json:"created_at"`
	UpdatedAt       time.Time                        `json:"updated_at"`
	DeletedAt       gorm.DeletedAt                   `json:"deleted_at"`
}

func FromDomain(domain *wallets.Domain) *Wallet {
	var historiesData []sqlWalletHistory.WalletHistory
	historiesFromDomain := domain.HistoriesWallet
	for _, history := range historiesFromDomain {
		historiesData = append(historiesData, sqlWalletHistory.WalletHistory{
			HistoryWalletID: history.HistoryWalletID,
			NoWallet:        history.NoWallet,
			CashIn:          history.CashIn,
			CashOut:         history.CashOut,
			DateWallet:      history.DateWallet,
			CreatedAt:       history.UpdatedAt,
			UpdatedAt:       history.UpdatedAt,
			DeletedAt:       history.DeletedAt,
		})
	}

	return &Wallet{
		NoWallet:        domain.NoWallet,
		UserID:          domain.UserID,
		Balance:         domain.Balance,
		HistoriesWallet: historiesData,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
		DeletedAt:       domain.DeletedAt,
	}
}

func (rec *Wallet) ToDomain() wallets.Domain {

	var historiesFromDomain []busWalletHistory.Domain
	for _, history := range rec.HistoriesWallet {
		historiesFromDomain = append(historiesFromDomain, history.ToDomain())
	}

	return wallets.Domain{
		NoWallet:        rec.NoWallet,
		UserID:          rec.UserID,
		Balance:         rec.Balance,
		HistoriesWallet: historiesFromDomain,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
		DeletedAt:       rec.DeletedAt,
	}
}

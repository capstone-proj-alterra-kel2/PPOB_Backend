package response

import (
	resWalletHistory "PPOB_BACKEND/controllers/wallet_histories/response"

	"PPOB_BACKEND/businesses/wallet_histories"
	"PPOB_BACKEND/businesses/wallets"
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	NoWallet        string                           `json:"no_wallet" gorm:"size:16;primaryKey;unique"`
	UserID          uint                             `json:"user_id"`
	Balance         int                              `json:"balance"`
	HistoriesWallet []resWalletHistory.WalletHistory `json:"histories_wallet" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:no_wallet"`
	CreatedAt       time.Time                        `json:"created_at"`
	UpdatedAt       time.Time                        `json:"updated_at"`
	DeletedAt       gorm.DeletedAt                   `json:"deleted_at"`
}

func FromDomain(domain wallets.Domain) Wallet {
	var walletHistoryData []resWalletHistory.WalletHistory
	walletHistoryToDomain := domain.HistoriesWallet
	for _, history := range walletHistoryToDomain {
		walletHistoryData = append(walletHistoryData, resWalletHistory.FromDomain(wallet_histories.Domain{
			HistoryWalletID: history.HistoryWalletID,
			NoWallet:        history.NoWallet,
			CashIn:          history.CashIn,
			CashOut:         history.CashOut,
			DateWallet:      history.DateWallet,
			CreatedAt:       history.CreatedAt,
			UpdatedAt:       history.UpdatedAt,
			DeletedAt:       history.DeletedAt,
		}))
	}
	return Wallet{
		NoWallet:        domain.NoWallet,
		UserID:          domain.UserID,
		Balance:         domain.Balance,
		HistoriesWallet: walletHistoryData,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
		DeletedAt:       domain.DeletedAt,
	}
}

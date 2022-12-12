package users

import (
	"PPOB_BACKEND/businesses/users"
	"PPOB_BACKEND/drivers/postgresql/roles"

	recWalletHistory "PPOB_BACKEND/drivers/postgresql/wallet_histories"
	"PPOB_BACKEND/drivers/postgresql/wallets"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" gorm:"size:100;primaryKey"`
	Name        string         `json:"name" `
	PhoneNumber string         `json:"phone_number" gorm:"unique"`
	Email       string         `json:"email" gorm:"unique" `
	Password    string         `json:"password" `
	RoleID      uint           `json:"role_id"`
	Role        roles.Role     `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Image       string         `json:"image"`
	Wallet      wallets.Wallet `json:"wallet" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *users.Domain) *User {

	var historiesData []recWalletHistory.WalletHistory
	historiesFromDomain := domain.Wallet.HistoriesWallet
	for _, history := range historiesFromDomain {
		historiesData = append(historiesData, recWalletHistory.WalletHistory{
			HistoryWalletID: history.HistoryWalletID,
			NoWallet:        history.NoWallet,
			Income:          history.Income,
			Outcome:         history.Outcome,
			DateWallet:      history.DateWallet,
			CreatedAt:       history.CreatedAt,
			UpdatedAt:       history.UpdatedAt,
			DeletedAt:       history.DeletedAt,
		})
	}

	return &User{
		ID:          domain.ID,
		Name:        domain.Name,
		PhoneNumber: domain.PhoneNumber,
		Email:       domain.Email,
		Password:    domain.Password,
		RoleID:      domain.RoleID,
		Image: domain.Image,
		Wallet: wallets.Wallet{
			NoWallet:        domain.Wallet.NoWallet,
			UserID:          domain.Wallet.UserID,
			Balance:         domain.Wallet.Balance,
			HistoriesWallet: historiesData,
			CreatedAt:       domain.Wallet.CreatedAt,
			UpdatedAt:       domain.Wallet.UpdatedAt,
			DeletedAt:       domain.Wallet.DeletedAt,
		},
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdateAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *User) ToDomain() users.Domain {

	return users.Domain{
		ID:          rec.ID,
		Name:        rec.Name,
		PhoneNumber: rec.PhoneNumber,
		Email:       rec.Email,
		Password:    rec.Password,
		RoleID:      rec.RoleID,
		RoleName:    rec.Role.RoleName,
		Image:       rec.Image,
		Wallet:      rec.Wallet.ToDomain(),
		CreatedAt:   rec.CreatedAt,
		UpdateAt:    rec.UpdatedAt,
		DeletedAt:   rec.DeletedAt,
	}
}

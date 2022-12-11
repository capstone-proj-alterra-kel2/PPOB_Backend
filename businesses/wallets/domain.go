package wallets

import (
	busWalletHistory "PPOB_BACKEND/businesses/wallet_histories"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	NoWallet        string
	UserID          uint
	Balance         int
	HistoriesWallet []busWalletHistory.Domain
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type UpdateBalanceDomain struct {
	Balance int
}
type Usecase interface {
	GetWalletUser(idUser string) Domain
	GetAllWallet(Page int, Size int, Sort string) (*gorm.DB, []Domain)
	GetDetailWallet(noWallet string) Domain
	UpdateBalance(idUser string, balanceDomain *UpdateBalanceDomain) (Domain, error)
	IsiSaldo(idUser string, balanceDomain *UpdateBalanceDomain) (Domain, error)
}

type Repository interface {
	GetWalletUser(idUser string) Domain
	GetAllWallet(Page int, Size int, Sort string) (*gorm.DB, []Domain)
	GetDetailWallet(noWallet string) Domain
	UpdateBalance(idUser string, balanceDomain *UpdateBalanceDomain) (Domain, error)
	IsiSaldo(idUser string, balanceDomain *UpdateBalanceDomain) (Domain, error)
}

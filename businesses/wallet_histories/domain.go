package wallet_histories

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	HistoryWalletID uint
	NoWallet        string
	CashIn          int
	CashOut         int
	Description     string
	DateWallet      time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type CashInCashOutMonthlyDomain struct {
	CashIn  int
	CashOut int
	Month   string
}

type Usecase interface {
	GetWalletHistories(NoWallet string) []Domain
	GetWalletHistoriesMonthly(NoWallet string) []Domain
	GetCashInCashOutMonthly(NoWallet string) CashInCashOutMonthlyDomain
	GetDetailWalletHistories(idHistory string) Domain
	UpdateWalletHistories(idHistory string, domainHistory *Domain) Domain
	DeleteWalletHistories(idHistory string) bool
	CreateManual(NoWallet string, domainHistory *Domain) Domain
	CreateWalletHistory(NoWallet string, cashIn int, cashOut int, Description string) Domain
}

type Repository interface {
	GetWalletHistories(NoWallet string) []Domain
	GetWalletHistoriesMonthly(NoWallet string) []Domain
	GetCashInCashOutMonthly(NoWallet string) CashInCashOutMonthlyDomain
	GetDetailWalletHistories(idHistory string) Domain
	UpdateWalletHistories(idHistory string, domainHistory *Domain) Domain
	DeleteWalletHistories(idHistory string) bool
	CreateManual(NoWallet string, domainHistory *Domain) Domain
	CreateWalletHistory(NoWallet string, cashIn int, cashOut int, Description string) Domain
}

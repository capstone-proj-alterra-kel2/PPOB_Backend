package wallet_histories

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	HistoryWalletID uint
	NoWallet        string
	Income          int
	Outcome         int
	Description     string
	DateWallet      time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type OutcomeIncomeMonthlyDomain struct {
	Outcome int
	Income  int
	Month   string
}

type Usecase interface {
	GetWalletHistories(NoWallet string) []Domain
	GetWalletHistoriesMonthly(NoWallet string) []Domain
	GetOutcomeIncomeMonthly(NoWallet string) OutcomeIncomeMonthlyDomain
	GetDetailWalletHistories(idHistory string) Domain
	UpdateWalletHistories(idHistory string, domainHistory *Domain) Domain
	DeleteWalletHistories(idHistory string) bool
	CreateManual(NoWallet string, domainHistory *Domain) Domain
	CreateWalletHistory(NoWallet string, income int, outcome int, Description string) Domain
}

type Repository interface {
	GetWalletHistories(NoWallet string) []Domain
	GetWalletHistoriesMonthly(NoWallet string) []Domain
	GetOutcomeIncomeMonthly(NoWallet string) OutcomeIncomeMonthlyDomain
	GetDetailWalletHistories(idHistory string) Domain
	UpdateWalletHistories(idHistory string, domainHistory *Domain) Domain
	DeleteWalletHistories(idHistory string) bool
	CreateManual(NoWallet string, domainHistory *Domain) Domain
	CreateWalletHistory(NoWallet string, income int, outcome int, Description string) Domain
}

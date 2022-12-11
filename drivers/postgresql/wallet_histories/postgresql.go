package wallet_histories

import (
	"PPOB_BACKEND/businesses/wallet_histories"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type walletHistoryRepository struct {
	conn *gorm.DB
}

func NewPostgreSQLRepository(conn *gorm.DB) wallet_histories.Repository {
	return &walletHistoryRepository{
		conn: conn,
	}
}

func (whr *walletHistoryRepository) GetWalletHistories(NoWallet string) []wallet_histories.Domain {
	var rec []WalletHistory

	whr.conn.Find(&rec, "no_wallet = ?", NoWallet)

	historiesDomain := []wallet_histories.Domain{}

	for _, history := range rec {
		historiesDomain = append(historiesDomain, history.ToDomain())
	}

	return historiesDomain
}

func (whr *walletHistoryRepository) CreateWalletHistory(NoWallet string, Income int, Outcome int, Description string) wallet_histories.Domain {
	rec := WalletHistory{}
	rec.NoWallet = NoWallet
	rec.Income = Income
	rec.Outcome = Outcome
	rec.Description = Description
	result := whr.conn.Create(&rec)
	result.Last(&rec)
	return rec.ToDomain()
}

func (whr *walletHistoryRepository) GetWalletHistoriesMonthly(NoWallet string) []wallet_histories.Domain {
	var rec []WalletHistory
	currentTime := time.Now()
	month := currentTime.Local().Month()
	monthN := int(month)
	year := currentTime.Local().Year()
	layout := "2006-01-02T15:04:05.000000+00:00"
	startdate, _ := time.Parse(layout, fmt.Sprintf("%d-%d-01T00:00:00.000000+00:00", year, monthN))
	enddate, _ := time.Parse(layout, fmt.Sprintf("%d-%d-31T23:59:59.000000+00:00", year, monthN))

	whr.conn.Where("created_at >= ? AND created_at <= ?", startdate, enddate).Find(&rec, "no_wallet = ?", NoWallet)

	historiesDomain := []wallet_histories.Domain{}

	for _, history := range rec {
		historiesDomain = append(historiesDomain, history.ToDomain())
	}

	return historiesDomain
}

func (whr *walletHistoryRepository) GetOutcomeIncomeMonthly(NoWallet string) wallet_histories.OutcomeIncomeMonthlyDomain {

	currentTime := time.Now()
	month := currentTime.Local().Month()
	monthN := int(month)
	year := currentTime.Local().Year()
	layout := "2006-01-02T15:04:05.000000+00:00"
	startdate, _ := time.Parse(layout, fmt.Sprintf("%d-%d-01T00:00:00.000000+00:00", year, monthN))
	enddate, _ := time.Parse(layout, fmt.Sprintf("%d-%d-31T23:59:59.000000+00:00", year, monthN))
	table := whr.conn.Table("wallet_histories").Where("no_wallet = ?", NoWallet).Where("created_at >= ? AND created_at <= ?", startdate, enddate)
	var outcome int
	var income int
	outcomeTotal := table.Select("sum(outcome)").Row()
	incomeTotal := table.Select("sum(income)").Row()
	outcomeTotal.Scan(&outcome)
	incomeTotal.Scan(&income)

	outcomeincomeDomain := wallet_histories.OutcomeIncomeMonthlyDomain{
		Income:  income,
		Outcome: outcome,
	}
	return outcomeincomeDomain
}

func (whr *walletHistoryRepository) CreateManual(NoWallet string, historyDomain *wallet_histories.Domain) wallet_histories.Domain {
	rec := FromDomain(historyDomain)
	rec.NoWallet = NoWallet

	result := whr.conn.Create(&rec)
	result.Last(&rec)
	return rec.ToDomain()
}

func (whr *walletHistoryRepository) UpdateWalletHistories(idHistory string, historyDomain *wallet_histories.Domain) wallet_histories.Domain {
	var walletHistory wallet_histories.Domain = whr.GetDetailWalletHistories(idHistory)
	updatedHistory := FromDomain(&walletHistory)
	updatedHistory.Income = historyDomain.Income
	updatedHistory.Outcome = historyDomain.Outcome
	updatedHistory.Description = historyDomain.Description
	whr.conn.Save(&updatedHistory)
	return updatedHistory.ToDomain()
}

func (whr *walletHistoryRepository) DeleteWalletHistories(idHistory string) bool {
	var historyDomain wallet_histories.Domain = whr.GetDetailWalletHistories(idHistory)
	deletedHistory := FromDomain(&historyDomain)
	if result := whr.conn.Unscoped().Delete(&deletedHistory); result.RowsAffected == 0 {
		return false
	}
	return true
}

func (whr *walletHistoryRepository) GetDetailWalletHistories(idHistory string) wallet_histories.Domain {
	walletHistory := WalletHistory{}
	whr.conn.First(&walletHistory, "history_wallet_id = ?")
	return walletHistory.ToDomain()
}

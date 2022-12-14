package wallets

import (
	"PPOB_BACKEND/businesses/wallets"
	"strings"

	"gorm.io/gorm"
)

type walletRepository struct {
	conn *gorm.DB
}

func NewPostgreSQLRepository(conn *gorm.DB) wallets.Repository {
	return &walletRepository{
		conn: conn,
	}
}

func (wr *walletRepository) GetWalletUser(idUser string) wallets.Domain {
	var wallet Wallet

	wr.conn.
		Preload("HistoriesWallet").
		First(&wallet, "user_id=?", idUser)

	return wallet.ToDomain()
}

func (wr *walletRepository) TopUpBalance(idUser string, balanceDomain *wallets.UpdateBalanceDomain) (wallets.Domain, error) {

	var wallet wallets.Domain = wr.GetWalletUser(idUser)
	updatedData := FromDomain(&wallet)

	updatedData.Balance = updatedData.Balance + balanceDomain.Balance

	err := wr.conn.Save(&updatedData).Error
	if err != nil {
		return updatedData.ToDomain(), err
	}

	return updatedData.ToDomain(), nil
}

func (wr *walletRepository) UpdateBalance(idUser string, balanceDomain *wallets.UpdateBalanceDomain) (wallets.Domain, error) {
	var wallet wallets.Domain = wr.GetWalletUser(idUser)
	updatedData := FromDomain(&wallet)

	updatedData.Balance = balanceDomain.Balance

	err := wr.conn.Save(&updatedData).Error
	if err != nil {
		return updatedData.ToDomain(), err
	}

	return updatedData.ToDomain(), nil
}

func (wr *walletRepository) GetAllWallet(Page int, Size int, Sort string) (*gorm.DB, []wallets.Domain) {
	var rec []Wallet
	var sort string
	var model *gorm.DB
	if strings.HasPrefix(Sort, "-") {
		sort = Sort[1:] + " DESC"
	} else {
		sort = Sort[0:] + " ASC"
	}
	model = wr.conn.Order(sort).Model(&rec)

	wr.conn.Preload("wallet").Offset(Page).Limit(Size).Order(sort).Find(&rec)

	walletDomain := []wallets.Domain{}
	for _, wallet := range rec {
		walletDomain = append(walletDomain, wallet.ToDomain())
	}

	return model, walletDomain
}

func (wr *walletRepository) GetDetailWallet(noWallet string) wallets.Domain {
	var wallet Wallet

	wr.conn.Preload("Wallet").First(&wallet, "no_wallet=?", noWallet)

	return wallet.ToDomain()
}

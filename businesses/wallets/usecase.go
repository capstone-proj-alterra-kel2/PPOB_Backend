package wallets

import "gorm.io/gorm"

type walletUsecase struct {
	walletRepository Repository
}

func NewWalletUseCase(wr Repository) Usecase {
	return &walletUsecase{
		walletRepository: wr,
	}
}

func (wu *walletUsecase) GetWalletUser(idUser string) Domain {
	return wu.walletRepository.GetWalletUser(idUser)
}

func (wu *walletUsecase) TopUpBalance(idUser string, balanceDomain *UpdateBalanceDomain) (Domain, error) {
	return wu.walletRepository.TopUpBalance(idUser, balanceDomain)
}

func (wu *walletUsecase) UpdateBalance(idUser string, balanceDomain *UpdateBalanceDomain) (Domain, error) {
	return wu.walletRepository.UpdateBalance(idUser, balanceDomain)
}

func (wu *walletUsecase) UpdateAdminBalance(idUser string, balanceDomain *UpdateBalanceDomain) (Domain, error) {
	return wu.walletRepository.UpdateBalance(idUser, balanceDomain)
}

func (wu *walletUsecase) GetAllWallet(Page int, Size int, Sort string) (*gorm.DB, []Domain) {
	return wu.walletRepository.GetAllWallet(Page, Size, Sort)
}

func (wu *walletUsecase) GetDetailWallet(noWallet string) Domain {
	return wu.walletRepository.GetDetailWallet(noWallet)
}
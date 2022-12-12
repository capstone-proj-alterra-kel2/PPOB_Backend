package wallet_histories

type walletHistoryUsecase struct {
	walletHistoryRepository Repository
}

func NewWalletHistoryUseCase(whr Repository) Usecase {
	return &walletHistoryUsecase{
		walletHistoryRepository: whr,
	}
}

func (whu *walletHistoryUsecase) GetWalletHistories(NoWallet string) []Domain {
	return whu.walletHistoryRepository.GetWalletHistories(NoWallet)
}

func (whu *walletHistoryUsecase) GetCashInCashOutMonthly(NoWallet string) CashInCashOutMonthlyDomain {
	return whu.walletHistoryRepository.GetCashInCashOutMonthly(NoWallet)
}

func (whu *walletHistoryUsecase) GetWalletHistoriesMonthly(NoWallet string) []Domain {
	return whu.walletHistoryRepository.GetWalletHistoriesMonthly(NoWallet)
}

func (whu *walletHistoryUsecase) CreateWalletHistory(NoWallet string, cashIn int, cashOut int, Description string) Domain {
	return whu.walletHistoryRepository.CreateWalletHistory(NoWallet, cashIn, cashOut, Description)
}

func (whu *walletHistoryUsecase) GetDetailWalletHistories(idHistory string) Domain {
	return whu.walletHistoryRepository.GetDetailWalletHistories(idHistory)
}

func (whu *walletHistoryUsecase) UpdateWalletHistories(idHistory string, domainHistory *Domain) Domain {
	return whu.walletHistoryRepository.UpdateWalletHistories(idHistory, domainHistory)
}

func (whu *walletHistoryUsecase) DeleteWalletHistories(idHistory string) bool {
	return whu.walletHistoryRepository.DeleteWalletHistories(idHistory)
}

func (whu *walletHistoryUsecase) CreateManual(NoWallet string, domainHistory *Domain) Domain {
	return whu.walletHistoryRepository.CreateManual(NoWallet, domainHistory)
}

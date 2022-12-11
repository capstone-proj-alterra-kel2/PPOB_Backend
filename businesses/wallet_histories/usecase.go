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

func (whu *walletHistoryUsecase) GetOutcomeIncomeMonthly(NoWallet string) OutcomeIncomeMonthlyDomain {
	return whu.walletHistoryRepository.GetOutcomeIncomeMonthly(NoWallet)
}

func (whu *walletHistoryUsecase) GetWalletHistoriesMonthly(NoWallet string) []Domain {
	return whu.walletHistoryRepository.GetWalletHistoriesMonthly(NoWallet)
}

func (whu *walletHistoryUsecase) CreateWalletHistory(NoWallet string, Income int, Outcome int, Description string) Domain {
	return whu.walletHistoryRepository.CreateWalletHistory(NoWallet, Income, Outcome, Description)
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

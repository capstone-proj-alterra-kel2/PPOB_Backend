package drivers

import (
	productDomain "PPOB_BACKEND/businesses/products"
	productTypeDomain "PPOB_BACKEND/businesses/producttypes"
	providerDomain "PPOB_BACKEND/businesses/providers"
	trnsactionDomain "PPOB_BACKEND/businesses/transactions"
	userDomain "PPOB_BACKEND/businesses/users"
	productDB "PPOB_BACKEND/drivers/postgresql/products"
	productTypeDB "PPOB_BACKEND/drivers/postgresql/producttypes"
	providerDB "PPOB_BACKEND/drivers/postgresql/providers"
	transactionDB "PPOB_BACKEND/drivers/postgresql/transactions"
	userDB "PPOB_BACKEND/drivers/postgresql/users"

	walletDomain "PPOB_BACKEND/businesses/wallets"
	walletDB "PPOB_BACKEND/drivers/postgresql/wallets"

	walletHistoryDomain "PPOB_BACKEND/businesses/wallet_histories"
	walletHistoryDB "PPOB_BACKEND/drivers/postgresql/wallet_histories"

	categoryDomain "PPOB_BACKEND/businesses/category"
	categoryDB "PPOB_BACKEND/drivers/postgresql/category"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewPostgreSQLRepository(conn)
}

func NewProductRepository(conn *gorm.DB) productDomain.Repository {
	return productDB.NewPostgreSQLRepository(conn)
}

func NewProductTypeRepository(conn *gorm.DB) productTypeDomain.Repository {
	return productTypeDB.NewPostgreSQLRepository(conn)
}

func NewProviderRepository(conn *gorm.DB) providerDomain.Repository {
	return providerDB.NewPostgreSQLRepository(conn)
}

func NewWalletHistoryRepository(conn *gorm.DB) walletHistoryDomain.Repository {
	return walletHistoryDB.NewPostgreSQLRepository(conn)
}

func NewWalletRepository(conn *gorm.DB) walletDomain.Repository {
	return walletDB.NewPostgreSQLRepository(conn)
}

func NewTransactionRepository(conn *gorm.DB) trnsactionDomain.Repository {
	return transactionDB.NewTransactionRepository(conn)
}

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewCategoryRepository(conn)
}

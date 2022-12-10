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

func NewTransactionRepository(conn *gorm.DB) trnsactionDomain.Repository {
	return transactionDB.NewTransactionRepository(conn)
}

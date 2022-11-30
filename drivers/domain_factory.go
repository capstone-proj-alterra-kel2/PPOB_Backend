package drivers

import (
	productDomain "PPOB_BACKEND/businesses/products"
	productTypeDomain "PPOB_BACKEND/businesses/producttypes"
	providerDomain "PPOB_BACKEND/businesses/providers"
	stockDomain "PPOB_BACKEND/businesses/stocks"
	userDomain "PPOB_BACKEND/businesses/users"
	productDB "PPOB_BACKEND/drivers/postgresql/products"
	productTypeDB "PPOB_BACKEND/drivers/postgresql/producttypes"
	providerDB "PPOB_BACKEND/drivers/postgresql/providers"
	stockDB "PPOB_BACKEND/drivers/postgresql/stocks"
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

func NewStockRepository(conn *gorm.DB) stockDomain.Repository {
	return stockDB.NewPostgreSQLRepository(conn)
}

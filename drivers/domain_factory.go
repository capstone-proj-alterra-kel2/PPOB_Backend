package drivers

import (
	userDomain "PPOB_BACKEND/businesses/users"
	paymentmethodDomain "PPOB_BACKEND/businesses/payment_method"
	userDB "PPOB_BACKEND/drivers/postgresql/users"
	paymentmethodDB "PPOB_BACKEND/drivers/postgresql/payment_method"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewPostgreSQLRepository(conn)
}

func NewPaymentMethodRepository(conn *gorm.DB) paymentmethodDomain.Repository {
	return paymentmethodDB.NewPostgreSQLRepository(conn)
}
package drivers

import (
	userDomain "PPOB_BACKEND/businesses/users"
	userDB "PPOB_BACKEND/drivers/postgresql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewPostgreSQLRepository(conn)
}
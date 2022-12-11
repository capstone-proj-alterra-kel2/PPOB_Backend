package users

import (
	"PPOB_BACKEND/businesses/wallets"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID          uint
	Name        string
	PhoneNumber string
	Email       string
	Password    string
	RoleID      uint
	RoleName    string
	Wallet      wallets.Domain
	Image       string
	CreatedAt   time.Time
	UpdateAt    time.Time
	DeletedAt   gorm.DeletedAt
}

type LoginDomain struct {
	Email    string
	Password string
}
type UpdatePasswordDomain struct {
	OldPassword string
	NewPassword string
}

type UpdateDataDomain struct {
	Image       string
	Name        string
	PhoneNumber string
	Email       string
}

type UpdateImageDomain struct {
	Image string
}

type Usecase interface {
	GetAll(Page int, Size int, Sort string, Search string) (*gorm.DB, []Domain)
	GetAllAdmin(Page int, Size int, Sort string, Search string) (*gorm.DB, []Domain)
	Register(userDomain *Domain) (Domain, error)
	CreateAdmin(userDomain *Domain) (Domain, error)
	DeleteUser(idUser string) bool
	Login(userDomain *LoginDomain) string
	Profile(idUser string) Domain
	UpdatePassword(idUser string, passwordDomain *UpdatePasswordDomain) bool
	UpdateData(idUser string, dataDomain *UpdateDataDomain) (Domain, error)
	UpdateImage(idUser string, imageDomain *UpdateImageDomain) (Domain, error)
	CheckDuplicateUser(Email string, PhoneNumber string) (bool, bool)
}

type Repository interface {
	GetAll(Page int, Size int, Sort string, Search string) (*gorm.DB, []Domain)
	GetAllAdmin(Page int, Size int, Sort string, Search string) (*gorm.DB, []Domain)
	Register(userDomain *Domain) (Domain, error)
	CreateAdmin(userDomain *Domain) (Domain, error)
	DeleteUser(idUser string) bool
	Login(userDomain *LoginDomain) Domain
	Profile(idUser string) Domain
	UpdatePassword(idUser string, passwordDomain *UpdatePasswordDomain) bool
	UpdateData(idUser string, dataDomain *UpdateDataDomain) (Domain, error)
	UpdateImage(idUser string, imageDomain *UpdateImageDomain) (Domain, error)
	CheckDuplicateUser(Email string, PhoneNumber string) (bool, bool)
}

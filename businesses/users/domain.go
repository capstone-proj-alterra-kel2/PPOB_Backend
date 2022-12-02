package users

import (
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
	GetAll() []Domain
	Register(userDomain *Domain) (Domain, error)
	CreateAdmin(userDomain *Domain) (Domain, error)
	Login(userDomain *LoginDomain) string
	Profile(idUser string) Domain
	UpdatePassword(idUser string, passwordDomain *UpdatePasswordDomain) bool
	UpdateData(idUser string, dataDomain *UpdateDataDomain) (Domain, error)
	UpdateImage(idUser string, imageDomain *UpdateImageDomain) (Domain, error)
}

type Repository interface {
	GetAll() []Domain
	Register(userDomain *Domain) (Domain, error)
	CreateAdmin(userDomain *Domain) (Domain, error)
	Login(userDomain *LoginDomain) Domain
	Profile(idUser string) Domain
	UpdatePassword(idUser string, passwordDomain *UpdatePasswordDomain) bool
	UpdateData(idUser string, dataDomain *UpdateDataDomain) (Domain, error)
	UpdateImage(idUser string, imageDomain *UpdateImageDomain) (Domain, error)
}

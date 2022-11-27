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

type UpdatePasswordDomain struct {
	OldPassword string
	NewPassword string
}

type Usecase interface {
	GetAll() []Domain
	Register(userDomain *Domain) Domain
	Login(userDomain *Domain) string
	Profile(idUser string) Domain
	UpdatePassword(idUser string, passwordDomain *UpdatePasswordDomain) bool
}

type Repository interface {
	GetAll() []Domain
	Register(userDomain *Domain) Domain
	Login(userDomain *Domain) Domain
	Profile(idUser string) Domain
	UpdatePassword(idUser string, passwordDomain *UpdatePasswordDomain) bool
}

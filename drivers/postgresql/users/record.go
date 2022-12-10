package users

import (
	"PPOB_BACKEND/businesses/users"
	"PPOB_BACKEND/drivers/postgresql/roles"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" gorm:"size:100;primaryKey"`
	Name        string         `json:"name" `
	PhoneNumber string         `json:"phone_number" gorm:"unique"`
	Email       string         `json:"email" gorm:"unique" `
	Password    string         `json:"password" `
	RoleID      uint           `json:"role_id"`
	Role        roles.Role     `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Balance     int            `json:"balance"`
	Image       string         `json:"image"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *users.Domain) *User {

	return &User{
		ID:          domain.ID,
		Name:        domain.Name,
		PhoneNumber: domain.PhoneNumber,
		Email:       domain.Email,
		Password:    domain.Password,
		RoleID:      domain.RoleID,
		Balance:     domain.Balance,
		Image:       domain.Image,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdateAt,
		DeletedAt:   domain.DeletedAt,
	}
}

func (rec *User) ToDomain() users.Domain {

	return users.Domain{
		ID:          rec.ID,
		Name:        rec.Name,
		PhoneNumber: rec.PhoneNumber,
		Email:       rec.Email,
		Password:    rec.Password,
		RoleID:      rec.RoleID,
		RoleName:    rec.Role.RoleName,
		Balance:     rec.Balance,
		Image:       rec.Image,
		CreatedAt:   rec.CreatedAt,
		UpdateAt:    rec.UpdatedAt,
		DeletedAt:   rec.DeletedAt,
	}
}

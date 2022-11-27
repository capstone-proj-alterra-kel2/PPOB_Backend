package response

import (
	"PPOB_BACKEND/businesses/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	PhoneNumber string         `json:"phone_number"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Image       string         `json:"image"`
	RoleID      uint           `json:"role_id"`
	RoleName    string         `json:"role_name"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:          domain.ID,
		Name:        domain.Name,
		PhoneNumber: domain.PhoneNumber,
		Email:       domain.Email,
		Password:    domain.Password,
		RoleID:      domain.RoleID,
		RoleName:    domain.RoleName,
		Image:       domain.Image,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdateAt,
		DeletedAt:   domain.DeletedAt,
	}
}

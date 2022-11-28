package request

import (
	"PPOB_BACKEND/businesses/users"

	"github.com/go-playground/validator"
)

type User struct {
	Name        string `json:"name" form:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	Password    string `json:"password" form:"password" validate:"required"`
	Image       string `json:"image" form:"image"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdatePassword struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type UpdateData struct {
	Name        string `json:"name" form:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email"`
}

func (req *User) ToDomain() *users.Domain {
	return &users.Domain{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
		Image:       req.Image,
	}
}

func (req *User) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

func (req *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *UserLogin) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

func (req *UpdatePassword) ToDomain() *users.UpdatePasswordDomain {
	return &users.UpdatePasswordDomain{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
}

func (req *UpdatePassword) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

func (req *UpdateData) ToDomain() *users.UpdateDataDomain {
	return &users.UpdateDataDomain{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}
}

func (req *UpdateData) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}

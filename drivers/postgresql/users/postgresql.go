package users

import (
	"PPOB_BACKEND/businesses/users"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewPostgreSQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) GetAll() []users.Domain {
	var rec []User

	ur.conn.Find(&rec)

	userDomain := []users.Domain{}

	for _, user := range rec {
		userDomain = append(userDomain, user.ToDomain())
	}
	return userDomain
}

func (ur *userRepository) Register(userDomain *users.Domain) users.Domain {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)
	rec := FromDomain(userDomain)
	rec.Password = string(password)
	rec.RoleID = 1
	result := ur.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (ur *userRepository) Login(userDomain *users.Domain) users.Domain {
	var user User
	ur.conn.First(&user, "email=?", userDomain.Email)

	if user.ID == 0 {
		fmt.Println("user not found")
		return users.Domain{}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password)); err != nil {
		fmt.Println("wrong password")
		return users.Domain{}
	}
	return user.ToDomain()
}

func (ur *userRepository) Profile(idUser string) users.Domain {
	var user User

	ur.conn.First(&user, "id=?", idUser)

	return user.ToDomain()
}

func (ur *userRepository) UpdatePassword(idUser string, passDomain *users.UpdatePasswordDomain) bool {
	var user users.Domain = ur.Profile(idUser)
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passDomain.OldPassword)); err != nil {
		return false
	}
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(passDomain.NewPassword), bcrypt.DefaultCost)
	updated := FromDomain(&user)
	updated.Password = string(newPassword)

	ur.conn.Save(&updated)
	return true
}

func (ur *userRepository) UpdateData(idUser string, dataDomain *users.UpdateDataDomain) (users.Domain, error) {

	var user users.Domain = ur.Profile(idUser)
	updatedData := FromDomain(&user)

	updatedData.Name = dataDomain.Name
	updatedData.PhoneNumber = dataDomain.PhoneNumber
	updatedData.Email = dataDomain.Email
	
	err := ur.conn.Save(&updatedData).Error
	if err != nil {
		return updatedData.ToDomain(), err
	}
	fmt.Println(err)

	return updatedData.ToDomain(), nil
}
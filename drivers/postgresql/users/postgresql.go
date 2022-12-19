package users

import (
	"PPOB_BACKEND/businesses/users"
	"fmt"
	"strings"

	"github.com/nsuprun/ccgen"
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

func (ur *userRepository) GetAll(Page int, Size int, Sort string, Search string) (*gorm.DB, []users.Domain) {
	var rec []User
	var sort string
	var search string
	var model *gorm.DB
	if strings.HasPrefix(Sort, "-") {
		sort = Sort[1:] + " DESC"
	} else {
		sort = Sort[0:] + " ASC"
	}
	model = ur.conn.Order(sort).Model(&rec).Where("users.role_id = ?", 1).Preload("Role").Preload("Wallet").Preload("Wallet.HistoriesWallet")
	if Search != "" {
		search = "%" + Search + "%"
		model = ur.conn.Order(sort).Model(&rec).Where("users.name LIKE ? AND users.role_id = ?", search, 1).Preload("Role").Preload("Wallet").Preload("Wallet.HistoriesWallet")
	}

	ur.conn.Preload("Role").Preload("Wallet").Preload("Wallet.HistoriesWallet").Offset(Page).Limit(Size).Order(sort).Where("users.name LIKE ? AND users.role_id = ?", search, 1).Find(&rec)

	userDomain := []users.Domain{}
	for _, user := range rec {
		userDomain = append(userDomain, user.ToDomain())
	}

	return model, userDomain
}

func (ur *userRepository) GetAllAdmin(Page int, Size int, Sort string, Search string) (*gorm.DB, []users.Domain) {
	var rec []User
	var sort string
	var search string
	var model *gorm.DB
	if strings.HasPrefix(Sort, "-") {
		sort = Sort[1:] + " DESC"
	} else {
		sort = Sort[0:] + " ASC"
	}
	model = ur.conn.Order(sort).Model(&rec).Where("users.role_id = ?", 2).Preload("Role").Preload("Wallet").Preload("Wallet.HistoriesWallet")
	if Search != "" {
		search = "%" + Search + "%"
		model = ur.conn.Order(sort).Model(&rec).Where("users.name LIKE ? AND users.role_id = ?", search, 2).Preload("Role").Preload("Wallet").Preload("Wallet.HistoriesWallet")
	}

	ur.conn.Preload("Role").Preload("Wallet").Preload("Wallet.HistoriesWallet").Offset(Page).Limit(Size).Order(sort).Where("users.name LIKE ? AND users.role_id = ?", search, 2).Find(&rec)
	userDomain := []users.Domain{}
	for _, user := range rec {
		userDomain = append(userDomain, user.ToDomain())
	}

	return model, userDomain
}

func (ur *userRepository) Register(userDomain *users.Domain) (users.Domain, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)
	rec := FromDomain(userDomain)
	rec.Password = string(password)
	rec.RoleID = 1
	rec.Wallet.NoWallet = ccgen.Solo.GenerateOfLength(16)
	rec.Wallet.Balance = 0
	if err := ur.conn.Preload("Role").Create(&rec).Error; err != nil {
		return rec.ToDomain(), err
	}
	return rec.ToDomain(), nil
}

func (ur *userRepository) CreateAdmin(userDomain *users.Domain) (users.Domain, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)
	rec := FromDomain(userDomain)
	rec.Password = string(password)
	rec.RoleID = 2
	rec.Wallet.NoWallet = ccgen.Solo.GenerateOfLength(16)
	rec.Wallet.Balance = 0
	if err := ur.conn.Create(&rec).Error; err != nil {
		return rec.ToDomain(), err
	}
	return rec.ToDomain(), nil
}

func (ur *userRepository) DeleteUser(idUser string) bool {
	var user users.Domain = ur.Profile(idUser)

	deletedUser := FromDomain(&user)

	if result := ur.conn.Delete(&deletedUser); result.RowsAffected == 0 {
		return false
	}
	return true
}

func (ur *userRepository) Login(loginDomain *users.LoginDomain) users.Domain {
	var user User
	ur.conn.First(&user, "email=?", loginDomain.Email)

	if user.ID == 0 {
		fmt.Println("user not found")
		return users.Domain{}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDomain.Password)); err != nil {
		fmt.Println("wrong password")
		return users.Domain{}
	}
	return user.ToDomain()
}

func (ur *userRepository) Profile(idUser string) users.Domain {
	var user User

	ur.conn.Preload("Role").Preload("Wallet").Preload("Wallet.HistoriesWallet").First(&user, "id=?", idUser)

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

	if dataDomain.Image != "" {
		updatedData.Image = dataDomain.Image
	}
	if dataDomain.Name != "" {
		updatedData.Name = dataDomain.Name
	}
	if dataDomain.PhoneNumber != "" {
		updatedData.PhoneNumber = dataDomain.PhoneNumber
	}
	if dataDomain.Email != "" {
		updatedData.Email = dataDomain.Email
	}

	err := ur.conn.Save(&updatedData).Error
	if err != nil {
		return updatedData.ToDomain(), err
	}

	return updatedData.ToDomain(), nil
}

func (ur *userRepository) UpdateImage(idUser string, imageDomain *users.UpdateImageDomain) (users.Domain, error) {
	var user users.Domain = ur.Profile(idUser)
	updatedData := FromDomain(&user)

	updatedData.Image = imageDomain.Image

	err := ur.conn.Save(&updatedData).Error
	if err != nil {
		return updatedData.ToDomain(), err
	}

	return updatedData.ToDomain(), nil
}

func (ur *userRepository) CheckDuplicateUser(Email string, PhoneNumber string) (bool, bool) {
	var rec []User
	var DuplicateEmail bool
	var DuplicatePhoneNumber bool
	if isEmailDuplicate := ur.conn.First(&rec, "users.email = ?", Email).Error; isEmailDuplicate != gorm.ErrRecordNotFound {
		DuplicateEmail = true
	}

	if isPhoneNumberDuplicate := ur.conn.First(&rec, "users.phone_number = ?", PhoneNumber).Error; isPhoneNumberDuplicate != gorm.ErrRecordNotFound {
		DuplicatePhoneNumber = true
	}
	if DuplicateEmail && DuplicatePhoneNumber {
		return true, true
	}
	if DuplicateEmail {
		return true, false
	}
	if DuplicatePhoneNumber {
		return false, true
	}
	return false, false
}

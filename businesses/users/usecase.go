package users

import (
	"PPOB_BACKEND/app/middlewares"
	"strconv"

	"gorm.io/gorm"
)

type userUsecase struct {
	userRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewUserUseCase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &userUsecase{
		userRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *userUsecase) GetAll(Page int, Size int, Sort string, Search string) (*gorm.DB, []Domain) {
	return uu.userRepository.GetAll(Page, Size, Sort, Search)
}

func (uu *userUsecase) GetAllAdmin(Page int, Size int, Sort string, Search string) (*gorm.DB, []Domain) {
	return uu.userRepository.GetAllAdmin(Page, Size, Sort, Search)
}

func (uu *userUsecase) Register(userDomain *Domain) (Domain, error) {
	return uu.userRepository.Register(userDomain)
}

func (uu *userUsecase) CreateAdmin(userDomain *Domain) (Domain, error) {
	return uu.userRepository.CreateAdmin(userDomain)
}

func (uu *userUsecase) DeleteUser(idUser string) bool {
	return uu.userRepository.DeleteUser(idUser)
}

func (uu *userUsecase) Login(loginDomain *LoginDomain) string {
	user := uu.userRepository.Login(loginDomain)
	if user.ID == 0 {
		return ""
	}
	idUser := strconv.FormatUint(uint64(user.ID), 10)
	profile := uu.userRepository.Profile(idUser)
	token := uu.jwtAuth.GenerateToken(user.ID, profile.RoleName)
	return token
}

func (uu *userUsecase) Profile(idUser string) Domain {
	return uu.userRepository.Profile(idUser)
}

func (uu *userUsecase) UpdatePassword(idUser string, passwordDomain *UpdatePasswordDomain) bool {
	return uu.userRepository.UpdatePassword(idUser, passwordDomain)
}

func (uu *userUsecase) UpdateData(idUser string, dataDomain *UpdateDataDomain) (Domain, error) {
	return uu.userRepository.UpdateData(idUser, dataDomain)
}

func (uu *userUsecase) UpdateImage(idUser string, imageDomain *UpdateImageDomain) (Domain, error) {
	return uu.userRepository.UpdateImage(idUser, imageDomain)
}

func (uu *userUsecase) CheckDuplicateUser(Email string, PhoneNumber string) (bool, bool) {
	return uu.userRepository.CheckDuplicateUser(Email, PhoneNumber)
}

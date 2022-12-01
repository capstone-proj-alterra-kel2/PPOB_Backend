package users

import (
	"PPOB_BACKEND/app/middlewares"
	"strconv"
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

func (uu *userUsecase) GetAll() []Domain {
	return uu.userRepository.GetAll()
}

func (uu *userUsecase) Register(userDomain *Domain) (Domain, error) {
	return uu.userRepository.Register(userDomain)
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

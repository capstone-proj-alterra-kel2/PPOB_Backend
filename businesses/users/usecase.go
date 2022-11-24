package users

import (
	"PPOB_BACKEND/app/middlewares"
)

type userUsecase struct {
	userRepository Repository
	jwtAuth *middlewares.ConfigJWT
}

func NewUserUseCase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &userUsecase {
		userRepository: ur,
		jwtAuth: jwtAuth,
	}
}


func (uu *userUsecase) GetAll() []Domain {
	return uu.userRepository.GetAll()
}

func (uu *userUsecase) Register(userDomain *Domain) Domain {
	return uu.userRepository.Register(userDomain)
}

func (uu *userUsecase) Login(userDomain *Domain) string {
	user := uu.userRepository.Login(userDomain)
	if user.ID == 0 {
		return ""
	}
	token := uu.jwtAuth.GenerateToken(user.ID, user.RoleID)
	return token
}
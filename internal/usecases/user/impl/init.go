package impl

import (
	"base/internal/repository/db/user"
	ujwt "base/internal/usecases/jwt"
	userUsecase "base/internal/usecases/user"
)

type usecase struct {
	User user.Repository
	Ujwt ujwt.Usecase
}

func New(userRepository user.Repository, ujwt ujwt.Usecase) userUsecase.Usecase {
	return &usecase{
		User: userRepository,
		Ujwt: ujwt,
	}
}

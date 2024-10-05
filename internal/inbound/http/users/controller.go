package users

import (
	jwtUsecase "base/internal/usecases/jwt"
	"base/internal/usecases/user"

	"go.uber.org/dig"
)

type Controller struct {
	dig.In
	Users user.Usecase
	Jwt   jwtUsecase.Usecase
}

var _ StrictServerInterface = (*Controller)(nil)

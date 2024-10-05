package labels

import (
	jwtUsecase "base/internal/usecases/jwt"
	"base/internal/usecases/label"

	"go.uber.org/dig"
)

type Controller struct {
	dig.In
	Jwt    jwtUsecase.Usecase
	Labels label.Usecase
}

var _ StrictServerInterface = (*Controller)(nil)

package tasks

import (
	jwtUsecase "base/internal/usecases/jwt"
	"base/internal/usecases/task"

	"go.uber.org/dig"
)

type Controller struct {
	dig.In
	Jwt   jwtUsecase.Usecase
	Tasks task.Usecase
}

var _ StrictServerInterface = (*Controller)(nil)

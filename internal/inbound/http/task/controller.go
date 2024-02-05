package task

import (
	taskUsecase "base/internal/usecases/task"

	"go.uber.org/dig"
)

type Controller struct {
	dig.In
	Task taskUsecase.Usecase
}

var _ StrictServerInterface = (*Controller)(nil)

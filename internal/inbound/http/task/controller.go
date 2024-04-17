package task

import (
	taskUsecase "base/internal/usecase/task"

	"go.uber.org/dig"
)

type Controller struct {
	dig.In
	TaskUsecase taskUsecase.Usecase
}

var _ StrictServerInterface = (*Controller)(nil)

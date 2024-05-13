package impl

import (
	"base/internal/outbound/db/task"
	taskUsecase "base/internal/usecase/task"

	"go.uber.org/dig"
)

type usecase struct {
	dig.In
	TaskRepository task.Repository
}

func New(taskRepository task.Repository) taskUsecase.Usecase {
	return &usecase{
		TaskRepository: taskRepository,
	}
}

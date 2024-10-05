package impl

import (
	"base/internal/repository/db/task"
	taskUsecase "base/internal/usecases/task"
)

type usecase struct {
	Task task.Repository
}

func New(taskRepository task.Repository) taskUsecase.Usecase {
	return &usecase{
		Task: taskRepository,
	}
}

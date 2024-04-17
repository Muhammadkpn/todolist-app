package impl

import (
	// 	"base/internal/common/sdksql"
	// 	"base/internal/model"
	// 	"base/internal/repository/db/task"
	"base/internal/outbound/db/task"
	taskUsecase "base/internal/usecases/task"
)

type usecase struct {
	TaskRepository task.Repository
	// TaskGenericRepository sdksql.Repository[model.Task]
}

func New(taskRepository task.Repository) taskUsecase.Usecase {
	return &usecase{
		TaskRepository: taskRepository,
	}
}

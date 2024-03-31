package impl

import (
	// 	"base/internal/common/sdksql"
	// 	"base/internal/model"
	// 	"base/internal/repository/db/task"
	"base/internal/repository/db/task"
	taskUsecase "base/internal/usecase/task"
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

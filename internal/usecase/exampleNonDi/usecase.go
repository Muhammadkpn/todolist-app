package task

import (
	"base/internal/usecase/model"
	"context"
)

type Usecase interface {
	GetAllTasks(ctx context.Context) (tasks []model.Task, err error)
}

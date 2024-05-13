package task

import (
	"base/internal/outbound/model"
	"context"
)

type Repository interface {
	GetAllTasks(ctx context.Context) (tasks []model.Task, err error)
}

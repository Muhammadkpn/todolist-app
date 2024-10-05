package task

import (
	"base/internal/repository/db/model"
	"context"
)

type Repository interface {
	GetTask(ctx context.Context, filter map[string]interface{}, sortBy string, order string) (task []model.Task, err error)
	GetTaskId(ctx context.Context, id int) (consent model.Task, err error)
	CreateTask(ctx context.Context, request model.Task) (err error)
	UpdateTask(ctx context.Context, request model.Task) (err error)
	DeleteTask(ctx context.Context, id int) (err error)
}

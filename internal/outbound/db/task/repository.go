package task

import (
	"base/internal/outbound/model"
	"context"
)

type Repository interface {
	CreateTask(ctx context.Context, request model.Task) (response model.Task, err error)
	GetAllTasks(ctx context.Context) (tasks []model.Task, err error)
	GetTaskByID(ctx context.Context, id int64) (task model.Task, err error)
	UpdateTask(ctx context.Context, id int64, title string) (task model.Task, err error)
	DeleteTask(ctx context.Context, id int64) (err error)
}

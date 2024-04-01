package task

import (
	"base/internal/usecase/model"
	"context"
)

type Usecase interface {
	CreateTask(ctx context.Context, title string) (task model.Task, err error)
	GetAllTasks(ctx context.Context) (tasks []model.Task, err error)
	GetTaskByID(ctx context.Context, id int64) (task model.Task, err error)
	UpdateTask(ctx context.Context, id int64, title string) (task model.Task, err error)
	DeleteTask(ctx context.Context, id int64) (err error)
}

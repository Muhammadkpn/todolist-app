package task

import (
	"base/internal/usecases/model"
	"context"
)

type Usecase interface {
	CreateTask(ctx context.Context, title string) (task model.Task, err error)
	GetAllTasks(ctx context.Context) (tasks []model.Task, err error)
	GetTaskByID(ctx context.Context, id int) (task model.Task, err error)
	UpdateTask(ctx context.Context, id int, title string) (task model.Task, err error)
	DeleteTask(ctx context.Context, id int) (err error)
}

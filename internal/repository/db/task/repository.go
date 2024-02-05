package task

import (
	"base/internal/repository/db/model"
	"context"
)

type Repository interface {
	CreateTask(title string) (task model.Task, err error)
	GetAllTasks(ctx context.Context) (tasks []model.Task, err error)
	GetTaskByID(id int) (task model.Task, err error)
	UpdateTask(id int, title string) (task model.Task, err error)
	DeleteTask(id int) (err error)
}

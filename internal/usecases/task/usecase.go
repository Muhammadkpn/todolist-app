package task

import (
	"base/internal/inbound/model"
	"context"
	"time"
)

type Usecase interface {
	GetTaskById(ctx context.Context, taskId int) (tasks model.Tasks, err error)
	GetTask(ctx context.Context, filter map[string]interface{}, sortBy string, order string) (tasks []model.Tasks, err error)
	CreateTask(ctx context.Context, user_id int, username string, title string, description string, status string, due_date time.Time, labels []int) (err error)
	EditTask(ctx context.Context, user_id int, username string, task_id int, title string, description string, status string, due_date time.Time, labels []int) (err error)
	DeleteTask(ctx context.Context, task_id int) (err error)
}

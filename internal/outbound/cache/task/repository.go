package task

import (
	"base/internal/outbound/db/model"
	"context"
)

type Repository interface {
	GetTaskByID(ctx context.Context, id int64) (task model.Task, err error)
	SetTaskByID(ctx context.Context, task model.Task) (err error)
}

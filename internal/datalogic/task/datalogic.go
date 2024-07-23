package task

import (
	"base/internal/datalogic/model"
	"context"
)

type Datalogic interface {
	GetTaskByID(ctx context.Context, id int64) (task model.Task, err error)
}

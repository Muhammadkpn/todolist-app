package impl

import (
	"base/internal/repository/db/model"
	"base/pkg/constant"
	"context"
)

func (r *repository) CreateTask(title string) (task model.Task, err error) {
	_, err = r.Db.ExecContext(context.Background(), constant.CreateTaskQuery, title)

	return
}

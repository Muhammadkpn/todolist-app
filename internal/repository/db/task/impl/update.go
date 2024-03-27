package impl

import (
	"base/internal/repository/db/model"
	"context"
)

func (r *repository) UpdateTask(ctx context.Context, id int64, title string) (task model.Task, err error) {
	// _, err = r.DbGorm.ExecContext(context.Background(), constant.UpdateTaskQuery, id, title)

	return
}

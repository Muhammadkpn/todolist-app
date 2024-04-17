package impl

import (
	"base/internal/outbound/db/model"
)

func (r *repository) UpdateTask(id int64, title string) (task model.Task, err error) {
	// _, err = r.DbGorm.ExecContext(context.Background(), constant.UpdateTaskQuery, id, title)

	return
}

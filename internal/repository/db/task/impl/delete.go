package impl

import (
	"base/pkg/constant"
	"context"
)

func (r *repository) DeleteTask(id int64) (err error) {
	_, err = r.Db.ExecContext(context.Background(), constant.DeleteTaskQuery, id)

	return
}

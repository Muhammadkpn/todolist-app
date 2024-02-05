package impl

import (
	"base/pkg/constant"
	"context"
)

func (r *repository) DeleteTask(id int) (err error) {
	_, err = r.Db.ExecContext(context.Background(), constant.DeleteTaskQuery, id)

	return
}

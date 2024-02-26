package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"
)

func (r *repository) DeleteTask(ctx context.Context, id int64) (err error) {
	// _, err = r.Db.ExecContext(context.Background(), constant.DeleteTaskQuery, id)
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	task := model.Task{
		ID: id,
	}

	db := r.DbGorm.WithContext(ctx)
	db.Model(task).Update("status", 0)

	return
}

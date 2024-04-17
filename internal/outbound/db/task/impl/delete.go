package impl

import (
	"base/internal/outbound/db/model"
	pkgHelper "base/pkg/helper"
	"context"
)

func (r *repository) DeleteTask(ctx context.Context, id int64) (err error) {
	// _, err = r.Db.ExecContext(context.Background(), constant.DeleteTaskQuery, id)
	span, ctx := pkgHelper.UpdateCtxSpanRepository(ctx)
	defer span.End()

	task := model.Task{
		ID: id,
	}

	db := r.DbGorm.WithContext(ctx)
	result := db.Model(task).Update("status", 0)
	if result.Error != nil {
		err = result.Error

		return
	}

	return
}

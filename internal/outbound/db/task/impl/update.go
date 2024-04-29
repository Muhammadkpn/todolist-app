package impl

import (
	"base/internal/outbound/model"
	"base/pkg/constant"
	pkgHelper "base/pkg/helper"
	"context"
)

func (r *repository) UpdateTask(ctx context.Context, id int64, title string) (task model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanRepository(ctx)
	defer span.End()

	db := r.Db.WithContext(ctx)
	result := db.Exec(constant.UpdateTaskQuery, id, title)
	if result.Error != nil {
		err = result.Error

		return
	}

	return
}

package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"base/pkg/constant"
	"context"
)

func (r *repository) UpdateTask(ctx context.Context, id int64, title string) (task model.Task, err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	db := r.DbGorm.WithContext(ctx)
	result := db.Exec(constant.UpdateTaskQuery, id, title)
	if result.Error != nil {
		err = result.Error

		return
	}

	return
}

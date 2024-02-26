package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"
)

func (r *repository) CreateTask(ctx context.Context, request model.Task) (response model.Task, err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	db := r.DbGorm.WithContext(ctx)
	result := db.Table("task").Create(&request)
	if result.Error != nil {
		err = result.Error

		return
	}

	return
}

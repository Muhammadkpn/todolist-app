package impl

import (
	"base/internal/outbound/db/model"
	pkgHelper "base/pkg/helper"
	"context"
)

func (r *repository) CreateTask(ctx context.Context, request model.Task) (response model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanRepository(ctx)
	defer span.End()

	db := r.Db.WithContext(ctx)
	result := db.Table("task").Create(&request)
	if result.Error != nil {
		err = result.Error

		return
	}

	return
}

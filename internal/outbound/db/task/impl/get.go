package impl

import (
	"base/internal/outbound/model"
	pkgHelper "base/pkg/helper"
	"context"
)

func (r *repository) GetAllTasks(ctx context.Context) (tasks []model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanRepository(ctx)
	defer span.End()

	db := r.Db.WithContext(ctx)
	result := db.Table("task").Where("status = ?", 1).Find(&tasks)
	if result.Error != nil {
		err = result.Error

		return
	}

	return
}

func (r *repository) GetTaskByID(ctx context.Context, id int64) (task model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanRepository(ctx)
	defer span.End()

	db := r.Db.WithContext(ctx)
	result := db.Table("task").Where("id = ?", id).First(&task)
	if result.Error != nil {
		err = result.Error

		return
	}

	return
}

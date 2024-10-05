package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"
)

// CreateTask creates a new task in the database
func (r *repository) CreateTask(ctx context.Context, request model.Task) (err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	db := r.DbGorm.WithContext(ctx)

	result := db.Model(&model.Task{}).Omit("Label.*").Create(&request)

	err = util.HandleError(result)

	return
}

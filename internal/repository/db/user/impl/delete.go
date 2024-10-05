package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"

	"gorm.io/gorm"
)

func (r *repository) DeleteUser(ctx context.Context, id int) (err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	// Directly delete the user by ID
	result := r.DbGorm.WithContext(ctx).Delete(&model.User{}, id)
	if err = result.Error; err != nil {
		return err
	}

	err = util.HandleError(result)

	// If no rows were affected, it means the task was not found
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

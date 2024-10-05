package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"
)

// GetLabelUserId retrieves a label by user ID
func (r *repository) GetLabelUserId(ctx context.Context, userID int) (label []model.Label, err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	result := r.DbGorm.WithContext(ctx).Where("user_id = ?", userID).Find(&label)

	err = util.HandleError(result)

	return
}

// GetAll label
func (r *repository) GetAllLabel(ctx context.Context) (label []model.Label, err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	result := r.DbGorm.WithContext(ctx).Find(&label)

	err = util.HandleError(result)

	return
}

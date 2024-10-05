package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"
)

// CreateLabel creates a new label in the database
func (r *repository) CreateLabel(ctx context.Context, request model.Label) (err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	result := r.DbGorm.WithContext(ctx).Create(&request)
	err = util.HandleError(result)

	return
}

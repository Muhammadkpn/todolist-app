package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"
	"time"
)

// UpdateLabel updates an existing label by ID
func (r *repository) UpdateLabel(ctx context.Context, id int, request model.Label) (err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	db := r.DbGorm.WithContext(ctx)

	req := model.Label{
		Name:        request.Name,
		UpdatedBy:   request.UpdatedBy,
		UpdatedDate: time.Now(),
	}

	result := db.Model(&model.Label{}).Where("\"id\" = ?", id).Updates(&req)
	err = util.HandleErrorID(result)

	return
}

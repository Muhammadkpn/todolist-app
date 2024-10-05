package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"
	"time"
)

func (r *repository) UpdateUser(ctx context.Context, id int, request model.User) (user model.User, err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	db := r.DbGorm.WithContext(ctx)

	req := model.User{
		Username:    request.Username,
		Email:       request.Email,
		Password:    request.Password,
		UpdatedBy:   request.UpdatedBy,
		UpdatedDate: time.Now(),
	}

	result := db.Model(&user).Where("\"id\" = ?", id).Updates(&req)
	err = util.HandleError(result)

	return
}

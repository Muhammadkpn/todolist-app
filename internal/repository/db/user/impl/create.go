package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"
)

// CreateUser creates a new user in the database
func (r *repository) CreateUser(ctx context.Context, request model.User) (user model.User, err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	result := r.DbGorm.WithContext(ctx).Model(&user).Create(&request)
	err = util.HandleError(result)
	user = request

	return
}

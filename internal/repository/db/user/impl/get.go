package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"

	"gorm.io/gorm"
)

func (r *repository) GetUserByUsername(ctx context.Context, username string) (user model.User, err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	result := r.DbGorm.WithContext(ctx).Where("username = ?", username).Find(&user)

	err = util.HandleError(result)

	return
}

func (r *repository) GetUserByID(ctx context.Context, id int) (user model.User, err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	result := r.DbGorm.WithContext(ctx).Where("id = ?", id).First(&user)
	// result := r.DbGorm.Scopes(setSearchPath("postgres")).Find(&user)
	err = util.HandleErrorID(result)

	return
}

func setSearchPath(schema string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.Exec("SET search_path TO ?", schema)
		return db
	}
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	result := r.DbGorm.WithContext(ctx).Where("email = ?", email).Find(&user)

	err = util.HandleError(result)

	return
}

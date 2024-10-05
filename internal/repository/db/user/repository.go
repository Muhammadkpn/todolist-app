package user

import (
	"base/internal/repository/db/model"
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByID(ctx context.Context, id int) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	UpdateUser(ctx context.Context, id int, user model.User) (model.User, error)
	DeleteUser(ctx context.Context, id int) error
}

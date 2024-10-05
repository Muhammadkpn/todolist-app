package user

import (
	"base/internal/inbound/model"
	"context"
)

type Usecase interface {
	Register(ctx context.Context, username string, email string, password string) (user model.Users, err error)
	Login(ctx context.Context, usernameOrEmail string, password string) (user model.Users, err error)
	GetUserByID(ctx context.Context, id int) (user model.Users, err error)
}

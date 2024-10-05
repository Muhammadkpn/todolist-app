package label

import (
	"base/internal/repository/db/model"
	"context"
)

type Repository interface {
	GetAllLabel(ctx context.Context) (label []model.Label, err error)
	GetLabelUserId(ctx context.Context, user_id int) (label []model.Label, err error)
	CreateLabel(ctx context.Context, request model.Label) (err error)
	UpdateLabel(ctx context.Context, id int, request model.Label) (err error)
	DeleteLabel(ctx context.Context, id int) (err error)
}

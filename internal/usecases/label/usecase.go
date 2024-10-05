package label

import (
	"base/internal/inbound/model"
	"context"
)

type Usecase interface {
	GetLabel(ctx context.Context, user_id int) (label []model.Labels, err error)
	CreateLabel(ctx context.Context, label_name string, username string, user_id int) (err error)
	EditLabel(ctx context.Context, label_name string, username string, label_id int) (err error)
	DeleteLabel(ctx context.Context, label_id int) (err error)
}

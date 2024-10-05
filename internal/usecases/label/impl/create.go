package impl

import (
	dbModel "base/internal/repository/db/model"
	"base/internal/util"
	"context"
	"time"
)

func (u *usecase) CreateLabel(ctx context.Context, label_name string, username string, user_id int) (err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	label := dbModel.Label{
		CreatedBy:   username,
		CreatedDate: time.Now(),
		Name:        label_name,
		UpdatedBy:   username,
		UpdatedDate: time.Now(),
		UserID:      user_id,
	}

	err = u.Label.CreateLabel(ctx, label)
	if err != nil {
		return err
	}
	return
}

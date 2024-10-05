package impl

import (
	dbModel "base/internal/repository/db/model"
	"base/internal/util"
	"context"
	"errors"
	"time"
)

func (u *usecase) EditLabel(ctx context.Context, label_name string, username string, label_id int) (err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	label := dbModel.Label{
		Name:        label_name,
		UpdatedBy:   username,
		UpdatedDate: time.Now(),
	}

	err = u.Label.UpdateLabel(ctx, label_id, label)
	if err != nil {
		if err.Error() == "data not found" {
			return errors.New("label not found")
		}
		return err
	}
	return
}

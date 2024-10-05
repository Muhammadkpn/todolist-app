package impl

import (
	"base/internal/util"
	"context"
	"errors"
)

func (u *usecase) DeleteLabel(ctx context.Context, label_id int) (err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	err = u.Label.DeleteLabel(ctx, label_id)
	if err != nil {
		if err.Error() == "data not found" {
			return errors.New("label not found")
		}
		return err
	}
	return
}

package impl

import (
	"base/internal/util"
	"context"
	"errors"
)

func (u *usecase) DeleteTask(ctx context.Context, task_id int) (err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	err = u.Task.DeleteTask(ctx, task_id)
	if err != nil {
		if err.Error() == "data not found" {
			return errors.New("task not found")
		}
		return err
	}
	return
}

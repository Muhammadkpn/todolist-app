package impl

import (
	"base/internal/util"
	"context"
)

func (u *usecase) DeleteTask(ctx context.Context, id int64) (err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	err = u.TaskRepository.DeleteTask(ctx, id)

	return
}

package impl

import (
	"base/internal/usecases/model"
	"base/internal/util"
	"context"
)

func (u *usecase) UpdateTask(ctx context.Context, id int64, title string) (task model.Task, err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	_, err = u.TaskRepository.UpdateTask(ctx, id, title)
	if err != nil {
		return model.Task{}, err
	}

	return
}

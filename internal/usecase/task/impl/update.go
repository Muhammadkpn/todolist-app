package impl

import (
	"base/internal/usecase/model"
	pkgHelper "base/pkg/helper"
	"context"
)

func (u *usecase) UpdateTask(ctx context.Context, id int64, title string) (task model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	_, err = u.TaskDb.UpdateTask(ctx, id, title)
	if err != nil {
		return model.Task{}, err
	}

	// invalidate cache

	// set cache pake datalogic

	return
}

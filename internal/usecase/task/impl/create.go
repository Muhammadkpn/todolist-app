package impl

import (
	dbModel "base/internal/outbound/db/model"
	"base/internal/usecase/model"
	pkgHelper "base/pkg/helper"
	"context"
)

func (u *usecase) CreateTask(ctx context.Context, title string) (task model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	taskDb := dbModel.Task{
		Title: title,
	}

	res, err := u.TaskDb.CreateTask(ctx, taskDb)
	if err != nil {
		return
	}

	task = model.Task{
		Title: res.Title,
	}

	return
}

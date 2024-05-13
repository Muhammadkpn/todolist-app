package impl

import (
	"base/internal/usecase/model"
	pkgHelper "base/pkg/helper"
	"context"
)

func (u *usecase) GetAllTasks(ctx context.Context) (tasks []model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	res, err := u.exampleNonDiRepository.GetAllTasks(ctx)
	if err != nil {
		return
	}

	for _, data := range res {
		tasks = append(tasks, model.Task{
			ID:     data.ID,
			Title:  data.Title,
			Status: data.Status,
		})
	}

	return
}

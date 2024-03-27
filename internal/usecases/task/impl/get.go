package impl

import (
	"base/internal/usecases/model"
	"base/internal/util"
	"context"
)

func (u *usecase) GetAllTasks(ctx context.Context) (tasks []model.Task, err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	res, err := u.TaskRepository.GetAllTasks(ctx)
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

func (u *usecase) GetTaskByID(ctx context.Context, id int64) (task model.Task, err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	_, err = u.TaskRepository.GetTaskByID(ctx, id)
	if err != nil {
		return model.Task{}, err
	}

	return model.Task{}, nil
}

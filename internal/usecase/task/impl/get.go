package impl

import (
	outboundModel "base/internal/outbound/model"
	"base/internal/usecase/model"
	pkgHelper "base/pkg/helper"
	"context"
)

func (u *usecase) GetAllTasks(ctx context.Context) (tasks []model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	res, err := u.TaskRepository.GetAllTasks(ctx)
	if err != nil {
		return
	}

	// sample call auth
	u.AuthRepository["active_directory"].Login(ctx, outboundModel.AuthRequest{
		Username: "025632",
		Password: "BSM123bsm",
	})

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
	span, ctx := pkgHelper.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	_, err = u.TaskRepository.GetTaskByID(ctx, id)
	if err != nil {
		return model.Task{}, err
	}

	return model.Task{}, nil
}

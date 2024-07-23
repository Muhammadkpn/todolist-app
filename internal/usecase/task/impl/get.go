package impl

import (
	outboundModel "base/internal/outbound/authentication/model"
	"base/internal/usecase/model"
	pkgHelper "base/pkg/helper"
	"context"
	"fmt"
)

func (u *usecase) GetAllTasks(ctx context.Context) (tasks []model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	res, err := u.TaskDb.GetAllTasks(ctx)
	if err != nil {
		return
	}

	// sample call auth
	resAd, errAd := u.AuthRepository["active_directory"].Login(ctx, outboundModel.AuthRequest{
		Username: "025632",
		Password: "BSM123bsm",
		AppName:  "Virtual Account",
	})
	// temporary print
	fmt.Println(resAd)
	fmt.Println(errAd)

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

	data, err := u.TaskDatalogic.GetTaskByID(ctx, id)
	if err != nil {
		return
	}

	task = model.Task{
		ID:     data.ID,
		Title:  data.Title,
		Status: data.Status,
	}

	return task, nil
}

package impl

import (
	dbModel "base/internal/outbound/model"
	"base/internal/usecase/model"
	pkgHelper "base/pkg/helper"
	"context"
)

// import (
// 	"base/internal/model"
// 	"context"

// 	"go.elastic.co/apm"
// 	"gorm.io/gorm"
// )

func (u *usecase) CreateTask(ctx context.Context, title string) (task model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	taskDb := dbModel.Task{
		Title: title,
	}

	res, err := u.TaskRepository.CreateTask(ctx, taskDb)
	if err != nil {
		return
	}

	task = model.Task{
		Title: res.Title,
	}

	return
}

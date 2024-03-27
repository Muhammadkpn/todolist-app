package impl

import (
	dbModel "base/internal/repository/db/model"
	"base/internal/usecases/model"
	"base/internal/util"
	"context"
)

// import (
// 	"base/internal/model"
// 	"context"

// 	"go.elastic.co/apm"
// 	"gorm.io/gorm"
// )

func (u *usecase) CreateTask(ctx context.Context, title string) (task model.Task, err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
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

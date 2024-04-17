package impl

import (
	obModel "base/internal/outbound/db/model"
	"base/internal/usecases/model"
	"context"

	"go.elastic.co/apm"
)

// import (
// 	"base/internal/model"
// 	"context"

// 	"go.elastic.co/apm"
// 	"gorm.io/gorm"
// )

func (u *usecase) CreateTask(ctx context.Context, title string) (task model.Task, err error) {
	// task, err = u.TaskRepository.CreateTask(title)
	span, _ := apm.StartSpan(ctx, "usecase", "CreateTask")
	defer span.End()

	taskDb := obModel.Task{
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

	// err = u.TaskGenericRepository.GetDb().Transaction(func(tx *gorm.DB) error {
	// 	errTx := u.TaskGenericRepository.Create(context.Background(), task, tx) // Pass tx to Insert
	// 	if errTx != nil {
	// 		return errTx
	// 	}

	// 	task.Title = "qq"
	// 	errTx = u.TaskGenericRepository.Create(context.Background(), task, tx) // Pass tx to Insert
	// 	if err != nil {
	// 		//
	// 		return err
	// 	}

	// 	return errTx // nil = commmit, err = rollback
	// })

	// u.TaskGenericRepository.Create(context.Background(), task, nil)

	return
}

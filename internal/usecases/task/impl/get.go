package impl

import (
	"base/internal/usecases/model"
	pkgHelper "base/pkg/helper"
	"context"

	"go.elastic.co/apm"
)

// import (
// 	"base/internal/common/util"
// 	"base/internal/model"
// 	"context"
// 	"time"

// 	"go.elastic.co/apm"
// )

func (u *usecase) GetAllTasks(ctx context.Context) (tasks []model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	// fmt.Println("masuk siniasdnaisd")

	// return

	// time.Sleep(2 * time.Second)

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

// func (u *usecase) GetAllTasks(ctx context.Context) (tasks []model.Task, err error) {
// span, ctx := util.UpdateCtxSpanUsecase(ctx)
// defer span.End()

// 	time.Sleep(1 * time.Second)

// tasks, err = u.TaskRepository.GetAllTasks(ctx)
// if err != nil {
// 	return
// }

// 	// v2 inject generic repo to usecase
// 	// filterWhere2 := sdksql.FilterWhere(
// 	// 	[]sdksql.Condition{
// 	// 		{
// 	// 			Field:    "id",
// 	// 			Operator: sdksql.Or,
// 	// 			Value: []sdksql.Condition{
// 	// 				{
// 	// 					Field:    "id",
// 	// 					Operator: sdksql.Equal,
// 	// 					Value:    "4",
// 	// 				},
// 	// 				{
// 	// 					Field:    "id",
// 	// 					Operator: sdksql.Equal,
// 	// 					Value:    5,
// 	// 				},
// 	// 			},
// 	// 		},
// 	// 		{
// 	// 			Field:    "status",
// 	// 			Operator: sdksql.Equal,
// 	// 			Value:    0,
// 	// 		},
// 	// 	},
// 	// )
// 	// filterWhere2 := sdksql.FilterWhere(
// 	// 	[]sdksql.Condition{
// 	// 		{
// 	// 			Field:    "status",
// 	// 			Operator: sdksql.Equal,
// 	// 			Value:    0,
// 	// 		},
// 	// 	},
// 	// )
// 	// tasksV2, errV2 := u.TaskGenericRepository.FindAll(ctx, filterWhere2)
// 	// if err != nil {
// 	// 	return tasksV2, errV2
// 	// }

// 	// tasks = tasksV2
// 	// err = errV2

// 	// return

// 	// query := "select * from task where status = @status and title like @title"
// 	// mapArgs := map[string]interface{}{
// 	// 	"status": 0,
// 	// 	"title":  "ma%",
// 	// }
// 	// tasksV3, errV3 := u.TaskGenericRepository.FindRaw(ctx, query, mapArgs)
// 	// if err != nil {
// 	// 	fmt.Println("masuk err?", errV3)
// 	// 	return tasksV3, errV3
// 	// }

// 	// fmt.Print("masuk ga sih?")
// 	// fmt.Println(tasksV3)

// 	// tasks = tasksV3
// 	// err = errV3

// 	// fmt.Println("task", tasksV3)
// 	// fmt.Println("err", errV3)

// 	return
// }

func (u *usecase) GetTaskByID(ctx context.Context, id int64) (task model.Task, err error) {
	span, _ := apm.StartSpan(ctx, "usecase", "GetTaskByID")
	defer span.End()

	_, err = u.TaskRepository.GetTaskByID(id)
	if err != nil {
		return model.Task{}, err
	}

	return model.Task{}, nil
}

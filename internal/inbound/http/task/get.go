package task

import (
	"base/internal/inbound/model"
	pkgHelper "base/pkg/helper"
	"context"

	sdklog "gitlab.banksinarmas.com/go/sdkv2/log"
)

func (c *Controller) GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error) {
	span, ctx := pkgHelper.UpdateCtxSpanController(ctx)
	defer span.End()
	// time.Sleep(1 * time.Second)

	// var resq GetTasks200JSONResponse

	// return resq, nil

	tasks, err := c.Task.GetAllTasks(ctx)
	if err != nil {
		statusCode := pkgHelper.FromErrorMap(err, model.ErrorMap)
		switch statusCode {
		default:
			return GetTasks500JSONResponse{Message: err.Error()}, nil
		}
	}

	var res GetTasks200JSONResponse

	for _, task := range tasks {
		*res.Data = append(*res.Data, model.Task{
			Id:     &task.ID,
			Status: &task.Status,
			Title:  &task.Title,
		})
	}

	sdklog.Debug(ctx, "gogogogo")

	return res, nil
}

func (c *Controller) GetTasksTaskId(ctx context.Context, request GetTasksTaskIdRequestObject) (GetTasksTaskIdResponseObject, error) {
	span, ctx := pkgHelper.UpdateCtxSpanController(ctx)
	defer span.End()

	task, err := c.Task.GetTaskByID(ctx, request.TaskId)
	if err != nil {
		statusCode := pkgHelper.FromErrorMap(err, model.ErrorMap)
		switch statusCode {
		default:
			return GetTasksTaskId500JSONResponse{Message: err.Error()}, nil
		}
	}

	res := GetTasksTaskId200JSONResponse{
		Data: &model.Task{
			Id:     &task.ID,
			Title:  &task.Title,
			Status: &task.Status,
		},
	}

	return res, nil
}

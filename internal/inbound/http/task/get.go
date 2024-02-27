package task

import (
	"base/internal/inbound/model"
	"base/internal/util"
	"context"

	sdklog "gitlab.banksinarmas.com/go/sdkv2/log"
)

func (c *Controller) GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()
	// time.Sleep(1 * time.Second)

	// var resq GetTasks200JSONResponse

	// return resq, nil

	tasks, err := c.Task.GetAllTasks(ctx)
	if err != nil {
		errCode := "asd"
		errMsg := err.Error()

		return GetTasks500JSONResponse{
			ErrorCode:    &errCode,
			ErrorMessage: &errMsg,
		}, nil
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
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	task, err := c.Task.GetTaskByID(ctx, request.TaskId)
	if err != nil {
		errCode := ""
		errMsg := err.Error()

		return GetTasksTaskId500JSONResponse{
			ErrorCode:    &errCode,
			ErrorMessage: &errMsg,
		}, nil
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

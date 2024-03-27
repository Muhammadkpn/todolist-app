package task

import (
	"base/internal/inbound/model"
	"base/internal/util"
	"context"
)

func (c *Controller) GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error) {
	// Update context and start span
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	// Get all tasks
	tasks, err := c.Task.GetAllTasks(ctx)
	if err != nil {
		// If there's an error, return 500 response
		return GetTasks500JSONResponse{
			ErrorCode:    util.StringPointer(""),
			ErrorMessage: util.StringPointer(err.Error()),
		}, nil
	}

	// Prepare response
	var responseData []model.Task
	for _, task := range tasks {
		responseData = append(responseData, model.Task{
			Id:     &task.ID,
			Status: &task.Status,
			Title:  &task.Title,
		})
	}

	return GetTasks200JSONResponse{
		Data: &responseData,
	}, nil
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

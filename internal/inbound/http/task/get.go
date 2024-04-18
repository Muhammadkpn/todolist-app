package task

import (
	"base/internal/inbound/model"
	pkgHelper "base/pkg/helper"
	"context"
)

func (c *Controller) GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error) {
	span, ctx := pkgHelper.UpdateCtxSpanController(ctx)
	defer span.End()

	// Get all tasks
	tasks, err := c.TaskUsecase.GetAllTasks(ctx)
	if err != nil {
		statusCode := pkgHelper.FromErrorMap(err.Error(), model.ErrorMap)
		switch statusCode {
		default:
			return GetTasks500JSONResponse{Message: err.Error()}, nil
		}
	}

	// Prepare response
	var responseData []model.Task
	for _, task := range tasks {
		responseData = append(responseData, model.Task{
			Id:     task.ID,
			Status: task.Status,
			Title:  task.Title,
		})
	}

	return GetTasks200JSONResponse{
		Data: &responseData,
	}, nil
}

func (c *Controller) GetTasksTaskId(ctx context.Context, request GetTasksTaskIdRequestObject) (GetTasksTaskIdResponseObject, error) {
	span, ctx := pkgHelper.UpdateCtxSpanController(ctx)
	defer span.End()

	task, err := c.TaskUsecase.GetTaskByID(ctx, request.TaskId)
	if err != nil {
		statusCode := pkgHelper.FromErrorMap(err.Error(), model.ErrorMap)
		switch statusCode {
		default:
			return GetTasksTaskId500JSONResponse{Message: err.Error()}, nil
		}
	}

	res := GetTasksTaskId200JSONResponse{
		Data: &model.Task{
			Id:     task.ID,
			Title:  task.Title,
			Status: task.Status,
		},
	}

	return res, nil
}

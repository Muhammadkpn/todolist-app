package task

import (
	"base/internal/util"
	"context"

	sdklog "github.com/fericosali/sdk/log"
)

func (c *Controller) GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	tasks, err := c.Task.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	var res GetTasks200JSONResponse

	for _, task := range tasks {
		ID := int(task.ID)

		res = append(res, struct {
			Id     *int    "json:\"id,omitempty\""
			Status *int    "json:\"status,omitempty\""
			Title  *string "json:\"title,omitempty\""
		}{
			Id:     &ID,
			Title:  &task.Title,
			Status: &task.Status,
		})
	}

	sdklog.Debug(context.Background(), "gogogogo")

	return res, nil
}

func (c *Controller) GetTasksTaskId(ctx context.Context, request GetTasksTaskIdRequestObject) (GetTasksTaskIdResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	_, err := c.Task.GetTaskByID(ctx, request.TaskId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

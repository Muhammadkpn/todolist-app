package task

import (
	"base/internal/util"
	"context"
)

func (c *Controller) PostTasks(ctx context.Context, request PostTasksRequestObject) (PostTasksResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	_, err := c.TaskUsecase.CreateTask(ctx, *request.Body.Title)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

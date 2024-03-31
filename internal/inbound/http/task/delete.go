package task

import (
	"base/internal/util"
	"context"
)

func (c *Controller) DeleteTasksTaskId(ctx context.Context, request DeleteTasksTaskIdRequestObject) (DeleteTasksTaskIdResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	err := c.TaskUsecase.DeleteTask(ctx, request.TaskId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

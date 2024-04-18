package task

import (
	pkgHelper "base/pkg/helper"
	"context"
)

func (c *Controller) DeleteTasksTaskId(ctx context.Context, request DeleteTasksTaskIdRequestObject) (DeleteTasksTaskIdResponseObject, error) {
	span, ctx := pkgHelper.UpdateCtxSpanController(ctx)
	defer span.End()

	err := c.TaskUsecase.DeleteTask(ctx, request.TaskId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

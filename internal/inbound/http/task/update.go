package task

import (
	pkgHelper "base/pkg/helper"
	"context"
)

func (c *Controller) PutTasksTaskId(ctx context.Context, request PutTasksTaskIdRequestObject) (PutTasksTaskIdResponseObject, error) {
	span, ctx := pkgHelper.UpdateCtxSpanController(ctx)
	defer span.End()

	// _, err := c.Task.UpdateTask(ctx, request.TaskId, *request.Body.Title)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

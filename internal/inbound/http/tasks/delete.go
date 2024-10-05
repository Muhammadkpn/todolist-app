package tasks

import (
	"base/internal/util"
	"base/pkg/constant"
	"context"
	"net/http"
	"strconv"
)

// Delete Task
// (DELETE /tasks/{id})
func (c *Controller) DeleteTasksById(ctx context.Context, request DeleteTasksByIdRequestObject) (DeleteTasksByIdResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	_, _, err := c.Jwt.CheckToken(ctx, request.Params.Authorization)
	if err != nil {
		return DeleteTasksById401JSONResponse{
			ErrorCode:    http.StatusUnauthorized,
			ErrorMessage: constant.MSG_UNAUTHORIZED,
		}, nil
	}

	taskId, _ := strconv.Atoi(request.Id)
	err = c.Tasks.DeleteTask(ctx, taskId)

	if err != nil {
		if err.Error() == "label not found" {
			return DeleteTasksById401JSONResponse{
				ErrorCode:    http.StatusNotFound,
				ErrorMessage: err.Error(),
			}, nil
		}
		return DeleteTasksById400JSONResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}, nil
	}

	return DeleteTasksById200JSONResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_GENERAL_SUCCESS,
	}, nil
}

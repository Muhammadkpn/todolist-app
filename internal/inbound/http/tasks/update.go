package tasks

import (
	"base/internal/util"
	"base/pkg/constant"
	"context"
	"net/http"
	"strconv"
)

// Edit Task
// (PUT /tasks/{id})
func (c *Controller) EditTasksById(ctx context.Context, request EditTasksByIdRequestObject) (EditTasksByIdResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	claims, err := c.Jwt.GetClaims(ctx, request.Params.Authorization)
	if err != nil {
		return EditTasksById401JSONResponse{
			ErrorCode:    http.StatusUnauthorized,
			ErrorMessage: constant.MSG_UNAUTHORIZED,
		}, nil
	}

	// CreateTask(ctx context.Context, user_id int, username string, title string, description string, status string, due_date time.Time, labels []int) (err error)
	userId, _ := strconv.Atoi(claims.Id)
	taskId, _ := strconv.Atoi(request.Id)
	err = c.Tasks.EditTask(ctx, userId, claims.Username, taskId, *request.Body.Title, *request.Body.Description, *request.Body.Status, request.Body.DueDate.Time, *request.Body.Labels)

	if err != nil {
		return EditTasksById400JSONResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}, nil
	}

	return EditTasksById200JSONResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_GENERAL_SUCCESS,
	}, nil
}

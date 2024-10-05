package tasks

import (
	"base/internal/util"
	"base/pkg/constant"
	"context"
	"net/http"
	"strconv"
)

// Create New Task
// (POST /tasks)
func (c *Controller) CreateTask(ctx context.Context, request CreateTaskRequestObject) (CreateTaskResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	claims, err := c.Jwt.GetClaims(ctx, request.Params.Authorization)
	if err != nil {
		return CreateTask401JSONResponse{
			ErrorCode:    http.StatusUnauthorized,
			ErrorMessage: constant.MSG_UNAUTHORIZED,
		}, nil
	}

	// CreateTask(ctx context.Context, user_id int, username string, title string, description string, status string, due_date time.Time, labels []int) (err error)
	userId, _ := strconv.Atoi(claims.Id)
	err = c.Tasks.CreateTask(ctx, userId, claims.Username, request.Body.Title, request.Body.Description, request.Body.Status, request.Body.DueDate.Time, *request.Body.Labels)

	if err != nil {
		return CreateTask400JSONResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}, nil
	}

	return CreateTask200JSONResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_GENERAL_SUCCESS,
	}, nil
}

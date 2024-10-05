package tasks

import (
	"base/internal/util"
	"base/pkg/constant"
	"context"
	"net/http"
	"strconv"
)

// Get Tasks
// (GET /tasks)
func (c *Controller) GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	_, _, err := c.Jwt.CheckToken(ctx, request.Params.Authorization)
	if err != nil {
		return GetTasks401JSONResponse{
			ErrorCode:    http.StatusUnauthorized,
			ErrorMessage: constant.MSG_UNAUTHORIZED,
		}, nil
	}

	filter := map[string]interface{}{}
	if request.Params.Status != nil {
		filter["status"] = *request.Params.Status
	}

	if request.Params.Label != nil {
		filter["labelId"] = *request.Params.Label
	}

	if request.Params.UserId != nil {
		filter["userId"] = *request.Params.UserId
	}

	sortBy := "id" // Default sort field
	if request.Params.SortBy != nil {
		sortBy = *request.Params.SortBy
	}

	orderBy := "ASC" // Default sort order
	if request.Params.OrderBy != nil {
		orderBy = *request.Params.OrderBy
	}

	res, err := c.Tasks.GetTask(ctx, filter, sortBy, orderBy)
	if err != nil {
		return GetTasks400JSONResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: constant.MSG_GENERAL_ERROR,
		}, nil
	}

	return GetTasks200JSONResponse{
		Data:         &res,
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_GENERAL_SUCCESS,
	}, nil
}

// Get Tasks by ID
// (GET /tasks/{id})
func (c *Controller) GetTasksById(ctx context.Context, request GetTasksByIdRequestObject) (GetTasksByIdResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	_, _, err := c.Jwt.CheckToken(ctx, request.Params.Authorization)
	if err != nil {
		return GetTasksById401JSONResponse{
			ErrorCode:    http.StatusUnauthorized,
			ErrorMessage: constant.MSG_UNAUTHORIZED,
		}, nil
	}

	taskId, _ := strconv.Atoi(request.Id)
	res, err := c.Tasks.GetTaskById(ctx, taskId)

	if err != nil {
		return GetTasksById400JSONResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: constant.MSG_GENERAL_ERROR,
		}, nil
	}

	return GetTasksById200JSONResponse{
		Data:         &res,
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_GENERAL_SUCCESS,
	}, nil
}

package users

import (
	"base/internal/util"
	"base/pkg/constant"
	"context"
	"net/http"
	"strconv"
)

// Get User Info by User ID
// (GET /users/{id})
func (c *Controller) GetUsersByUserId(ctx context.Context, request GetUsersByUserIdRequestObject) (GetUsersByUserIdResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	_, _, err := c.Jwt.CheckToken(ctx, request.Params.Authorization)
	if err != nil {
		return GetUsersByUserId401JSONResponse{
			ErrorCode:    http.StatusUnauthorized,
			ErrorMessage: constant.MSG_UNAUTHORIZED,
		}, nil
	}

	paramId, _ := strconv.Atoi(request.Id)
	req, err := c.Users.GetUserByID(ctx, paramId)
	if err != nil {
		if err.Error() == "data not found" {
			return GetUsersByUserId401JSONResponse{
				ErrorCode:    http.StatusNotFound,
				ErrorMessage: err.Error(),
			}, nil
		}
		return GetUsersByUserId400JSONResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}, nil
	}

	return GetUsersByUserId200JSONResponse{
		Data:         &req,
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_GENERAL_SUCCESS,
	}, nil
}

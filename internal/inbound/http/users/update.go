package users

import (
	"base/pkg/constant"
	"context"
	"net/http"
)

// Update user
// (PUT /users/{id})
func (c *Controller) UpdateUser(ctx context.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {
	return UpdateUser200JSONResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_GENERAL_SUCCESS,
	}, nil
}

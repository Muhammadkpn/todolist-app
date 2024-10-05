package labels

import (
	"base/internal/util"
	"base/pkg/constant"
	"context"
	"net/http"
)

// Get Labels
// (GET /labels)
func (c *Controller) GetLabels(ctx context.Context, request GetLabelsRequestObject) (GetLabelsResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	_, _, err := c.Jwt.CheckToken(ctx, request.Params.Authorization)
	if err != nil {
		return GetLabels401JSONResponse{
			ErrorCode:    http.StatusUnauthorized,
			ErrorMessage: constant.MSG_UNAUTHORIZED,
		}, nil
	}

	user_id := 0
	if request.Params.UserId != nil {
		user_id = *request.Params.UserId
	}
	res, err := c.Labels.GetLabel(ctx, user_id)
	if err != nil {
		return GetLabels400JSONResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: constant.MSG_GENERAL_ERROR,
		}, nil
	}

	return GetLabels200JSONResponse{
		Data:         &res,
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_GENERAL_SUCCESS,
	}, nil
}

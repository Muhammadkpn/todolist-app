package labels

import (
	"base/internal/util"
	"base/pkg/constant"
	"context"
	"net/http"
	"strconv"
)

// Get Labels
// (GET /labels)
func (c *Controller) CreateNewLabel(ctx context.Context, request CreateNewLabelRequestObject) (CreateNewLabelResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	claims, err := c.Jwt.GetClaims(ctx, request.Params.Authorization)
	if err != nil {
		return CreateNewLabel401JSONResponse{
			ErrorCode:    http.StatusUnauthorized,
			ErrorMessage: constant.MSG_UNAUTHORIZED,
		}, nil
	}

	userId, _ := strconv.Atoi(claims.Id)
	err = c.Labels.CreateLabel(ctx, request.Body.Name, claims.Username, userId)

	if err != nil {
		return CreateNewLabel400JSONResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}, nil
	}
	return CreateNewLabel200JSONResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_GENERAL_SUCCESS,
	}, nil
}

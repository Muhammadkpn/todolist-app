package labels

import (
	"base/internal/util"
	"base/pkg/constant"
	"context"
	"net/http"
	"strconv"
)

// Edit label
// (PUT /labels/{id})
func (c *Controller) EditLabelByID(ctx context.Context, request EditLabelByIDRequestObject) (EditLabelByIDResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	claims, err := c.Jwt.GetClaims(ctx, request.Params.Authorization)
	if err != nil {
		return EditLabelByID401JSONResponse{
			ErrorCode:    http.StatusUnauthorized,
			ErrorMessage: constant.MSG_UNAUTHORIZED,
		}, nil
	}

	labelId, _ := strconv.Atoi(request.Id)
	err = c.Labels.EditLabel(ctx, request.Body.Name, claims.Username, labelId)

	if err != nil {
		if err.Error() == "label not found" {
			return EditLabelByID401JSONResponse{
				ErrorCode:    http.StatusNotFound,
				ErrorMessage: err.Error(),
			}, nil
		}
		return EditLabelByID400JSONResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}, nil
	}

	return EditLabelByID200JSONResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_GENERAL_SUCCESS,
	}, nil
}

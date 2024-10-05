package labels

import (
	"base/internal/util"
	"base/pkg/constant"
	"context"
	"net/http"
	"strconv"
)

// Delete Label
// (DELETE /labels/{id})
func (c *Controller) DeleteLabelByID(ctx context.Context, request DeleteLabelByIDRequestObject) (DeleteLabelByIDResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	_, _, err := c.Jwt.CheckToken(ctx, request.Params.Authorization)
	if err != nil {
		return DeleteLabelByID401JSONResponse{
			ErrorCode:    http.StatusUnauthorized,
			ErrorMessage: constant.MSG_UNAUTHORIZED,
		}, nil
	}

	labelId, _ := strconv.Atoi(request.Id)
	err = c.Labels.DeleteLabel(ctx, labelId)

	if err != nil {
		if err.Error() == "label not found" {
			return DeleteLabelByID401JSONResponse{
				ErrorCode:    http.StatusNotFound,
				ErrorMessage: err.Error(),
			}, nil
		}
		return DeleteLabelByID400JSONResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}, nil
	}

	return DeleteLabelByID200JSONResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_GENERAL_SUCCESS,
	}, nil
}

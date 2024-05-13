package user

import (
	"net/http"

	pkgHelper "base/pkg/helper"

	"github.com/labstack/echo/v4"
	sdkResponse "gitlab.banksinarmas.com/go/sdkv2/common/response"
)

func (c *Controller) GetTask(eCtx echo.Context) error {
	span, ctx := pkgHelper.UpdateCtxSpanController(eCtx.Request().Context())
	defer span.End()

	task, err := c.ExampleNonDiUsecase.GetAllTasks(ctx)
	if err != nil {
		return sdkResponse.NewJSONResponse(eCtx, http.StatusInternalServerError, nil)
	}

	return sdkResponse.NewJSONResponse(eCtx, http.StatusOK, task)
}

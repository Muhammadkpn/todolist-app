package healthCheck

import (
	pkgHelper "base/pkg/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

// endpoint to check health check of the system depedencies, should be 1 to 1 to repository
func (c *Controller) HealthCheck(eCtx echo.Context) error {
	span, ctx := pkgHelper.UpdateCtxSpanController(eCtx.Request().Context())
	defer span.End()

	resUsecase, err := c.HealthCheckUsecase.HealthCheck(ctx)
	if err != nil {
		return eCtx.String(http.StatusInternalServerError, err.Error())
	}

	httpStatus := http.StatusOK
	for _, eachRes := range resUsecase {
		if eachRes.Error != nil {
			httpStatus = http.StatusInternalServerError

			break
		}
	}

	return eCtx.JSON(httpStatus, resUsecase)
}

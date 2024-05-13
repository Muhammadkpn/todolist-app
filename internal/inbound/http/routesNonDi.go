package http

import (
	"base/pkg/resource/model"

	"github.com/labstack/echo/v4"
)

func Routes(ec *echo.Echo) {
	ec.GET("/tasks/non_di", model.ControllerList.ControllerA.GetTask)
}

package http

import (
	"base/internal/inbound/http/petStore"
	"base/internal/inbound/http/task"
	pkgResource "base/pkg/resource"

	sdkHttpMiddleware "github.com/fericosali/sdk/appRunner/middleware/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Resource pkgResource.Resource

	PetStore petStore.Controller
	Task     task.Controller
}

func (i Inbound) Routes(ec *echo.Echo) {
	ec.Use(sdkHttpMiddleware.AppContext())
	base := ec.Group(i.Resource.Config.AppConfig.BasePath)

	v1 := base.Group("/v1")

	petStore.RegisterHandlers(v1, petStore.NewStrictHandler(&i.PetStore, nil))

	taskGroup := base.Group("")
	task.RegisterHandlers(taskGroup, task.NewStrictHandler(&i.Task, nil))
}

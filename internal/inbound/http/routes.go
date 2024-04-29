package http

import (
	"base/internal/inbound/http/healthCheck"
	"base/internal/inbound/http/petStore"
	"base/internal/inbound/http/task"
	"base/internal/inbound/http/user"
	pkgResource "base/pkg/resource"

	sdkHttpMiddleware "gitlab.banksinarmas.com/go/sdkv2/appRunner/middleware/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Resource pkgResource.Resource

	PetStore    petStore.Controller
	Task        task.Controller
	User        user.Controller
	HealthCheck healthCheck.Controller
}

func (i Inbound) Routes(ec *echo.Echo) {
	ec.Use(sdkHttpMiddleware.AppContext())
	ec.Validator = i.Resource.Validator
	base := ec.Group(i.Resource.Config.AppConfig.BasePath)

	// route for service health_check, mandatory to add if add new depedencies currently only postgres
	ec.GET("/health_check", i.HealthCheck.HealthCheck)

	v1 := base.Group("/v1")

	petStore.RegisterHandlers(v1, petStore.NewStrictHandler(&i.PetStore, nil))

	taskGroup := base.Group("")
	task.RegisterHandlers(taskGroup, task.NewStrictHandler(&i.Task, nil))

	userGroup := v1.Group("/user")
	userGroup.POST("", i.User.AddUser)
}

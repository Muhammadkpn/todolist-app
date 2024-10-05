package usecase

import (
	jwtUsecase "base/internal/usecases/jwt"
	jwtUsecaseImpl "base/internal/usecases/jwt/impl"
	labelUsecase "base/internal/usecases/label"
	labelUsecaseImpl "base/internal/usecases/label/impl"
	"base/internal/usecases/shared"
	sharedV1 "base/internal/usecases/shared/impl/v1"
	taskUsecase "base/internal/usecases/task"
	taskUsecaseImpl "base/internal/usecases/task/impl"
	userUsecase "base/internal/usecases/user"
	userUsecaseImpl "base/internal/usecases/user/impl"
	"fmt"

	"go.uber.org/dig"
)

// (optional)can inject usecase in global variable like row 19 (see internal/inbound/http/user/controller.go user for example)
type Usecase struct {
	dig.In

	Task   taskUsecase.Usecase
	User   userUsecase.Usecase
	Label  labelUsecase.Usecase
	Shared shared.UseCase
	Jwt    jwtUsecase.Usecase
}

// (required)depedency injection for usecase, for inject the depedency we need to provide usecase`s init func in the register function
func Register(container *dig.Container) error {
	if err := container.Provide(sharedV1.New); err != nil {
		return fmt.Errorf("[DI] error provide sharedV1 usecase: %+v", err)
	}

	if err := container.Provide(taskUsecaseImpl.New); err != nil {
		return fmt.Errorf("[DI] error provide Task usecase: %+v", err)
	}

	if err := container.Provide(userUsecaseImpl.New); err != nil {
		return fmt.Errorf("[DI] error provide User usecase: %+v", err)
	}

	if err := container.Provide(labelUsecaseImpl.New); err != nil {
		return fmt.Errorf("[DI] error provide Label usecase: %+v", err)
	}

	if err := container.Provide(jwtUsecaseImpl.New); err != nil {
		return fmt.Errorf("[DI] error provide JWT usecase: %+v", err)
	}

	return nil
}

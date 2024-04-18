package usecase

import (
	"base/internal/usecase/petStore"
	petStoreV1 "base/internal/usecase/petStore/impl/v1"
	"base/internal/usecase/shared"
	sharedV1 "base/internal/usecase/shared/impl/v1"
	taskUsecaseImpl "base/internal/usecase/task/impl"
	"fmt"

	"go.uber.org/dig"
)

// (optional)can inject usecase in global variable like row 19 (see internal/inbound/http/user/controller.go user for example)
type Usecase struct {
	dig.In

	PetStoreV1 petStore.UseCase `name:"petStoreV1"`
	Shared     shared.UseCase
}

// (required)depedency injection for usecase, for inject the depedency we need to provide usecase`s init func in the register function
func Register(container *dig.Container) error {
	err := container.Provide(taskUsecaseImpl.New)
	if err != nil {
		return err
	}

	if err := container.Provide(petStoreV1.New, dig.Name("petStoreV1")); err != nil {
		return fmt.Errorf("[DI] error provide petStoreV1 usecase: %+v", err)
	}

	if err := container.Provide(sharedV1.New); err != nil {
		return fmt.Errorf("[DI] error provide sharedV1 usecase: %+v", err)
	}

	return nil
}

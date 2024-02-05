package usecase

import (
	"base/internal/usecases/petStore"
	petStoreV1 "base/internal/usecases/petStore/impl/v1"
	petStoreV2 "base/internal/usecases/petStore/impl/v2"
	"base/internal/usecases/shared"
	sharedV1 "base/internal/usecases/shared/impl/v1"
	taskUsecaseImpl "base/internal/usecases/task/impl"
	"fmt"

	"go.uber.org/dig"
)

type Usecase struct {
	dig.In

	// TaskUsecase taskUsecase.Usecase
	PetStoreV1 petStore.UseCase `name:"petStoreV1"`
	PetStoreV2 petStore.UseCase `name:"petStoreV2"`
	Shared     shared.UseCase
}

func Register(container *dig.Container) error {
	err := container.Provide(taskUsecaseImpl.New)
	if err != nil {
		return err
	}

	if err := container.Provide(petStoreV1.New, dig.Name("petStoreV1")); err != nil {
		return fmt.Errorf("[DI] error provide petStoreV1 usecase: %+v", err)
	}

	if err := container.Provide(petStoreV2.New, dig.Name("petStoreV2")); err != nil {
		return fmt.Errorf("[DI] error provide petStoreV2 usecase: %+v", err)
	}

	if err := container.Provide(sharedV1.New); err != nil {
		return fmt.Errorf("[DI] error provide sharedV1 usecase: %+v", err)
	}

	return nil
}

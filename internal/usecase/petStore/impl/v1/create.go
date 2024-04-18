package petStoreV1

import (
	obModel "base/internal/outbound/model"
	"base/internal/usecase/model"
	pkgError "base/pkg/constant/error"
	"context"

	sdkError "gitlab.banksinarmas.com/go/sdkv2/common/error"
)

// Create implements petStore.UseCase.
func (u *usecase) Create(ctx context.Context, request model.CreatePetRequest) (model.Pet, error) {
	if err := u.resource.Validator.Validate(request); err != nil {
		return model.Pet{}, sdkError.New(pkgError.ERR_FAILED_ON_VALIDATOR, err)
	}

	data, err := u.outbound.Repository.PetStore.Create(ctx, u.resource.DatabaseSQL, obModel.Pet{
		Name:    request.Name,
		PetType: request.PetType,
	})
	if err != nil {
		return model.Pet{}, sdkError.Wrap(err.(sdkError.Err))
	}

	return model.Pet{}.FromDBModel(data), nil
}

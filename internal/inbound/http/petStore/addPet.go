package petStore

import (
	"base/internal/inbound/model"
	ucModel "base/internal/usecase/model"
	pkgHelper "base/pkg/helper"
	"context"

	sdkLogger "gitlab.banksinarmas.com/go/sdkv2/log/logger"
	sdkTime "gitlab.banksinarmas.com/go/sdkv2/time"
)

func (c *Controller) AddPet(ctx context.Context, request AddPetRequestObject) (AddPetResponseObject, error) {
	data, err := c.UseCase.PetStoreV1.Create(ctx, CreatePetRequestToUcModel(request))
	if err != nil {
		c.Resource.Logger.Error(ctx, "Error Add Pet", sdkLogger.WithError(err))

		statusCode := pkgHelper.FromErrorMap(err, model.ErrorMap)
		return AddPetdefaultJSONResponse{
			StatusCode: statusCode,
			Body: model.BaseResponse{
				Message:   err.Error(),
				Timestamp: sdkTime.Now().Unix(),
			},
		}, nil
	}

	return AddPet200JSONResponse{
		Data:      PetFromUcModel(data),
		Timestamp: sdkTime.Now().Unix(),
	}, nil
}

func PetFromUcModel(data ucModel.Pet) model.Pet {
	return model.Pet{
		Id:      int64(data.ID),
		Name:    data.Name,
		PetType: data.PetType,
	}
}

func CreatePetRequestToUcModel(data AddPetRequestObject) ucModel.CreatePetRequest {
	return ucModel.CreatePetRequest{
		Name:    data.Body.Name,
		PetType: data.Body.PetType,
	}
}

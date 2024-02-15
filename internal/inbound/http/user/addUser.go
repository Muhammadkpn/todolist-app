package user

import (
	"base/internal/inbound/model"
	"net/http"

	"github.com/labstack/echo/v4"
	sdkRequest "gitlab.banksinarmas.com/go/sdkv2/common/request"
	sdkResponse "gitlab.banksinarmas.com/go/sdkv2/common/response"
)

func (c *Controller) AddUser(eCtx echo.Context) error {
	var req model.CreateUserRequest
	if err := sdkRequest.BindAndValidate(eCtx, &req); err != nil {
		return sdkResponse.NewJSONResponse(eCtx, http.StatusBadRequest, err)
	}

	return sdkResponse.NewJSONResponse(eCtx, http.StatusOK, model.GetUserDetailResponse{
		ID:       1,
		FullName: req.FullName,
		Address:  req.Address,
	})
}

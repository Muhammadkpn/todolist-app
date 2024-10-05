package users

import (
	uModel "base/internal/usecases/model"
	"base/internal/util"
	"base/pkg/constant"
	"context"
	"net/http"
	"strconv"
)

// Login User
// (POST /login)
func (c *Controller) LoginUser(ctx context.Context, request LoginUserRequestObject) (LoginUserResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	reqLogin, err := c.Users.Login(ctx, request.Body.UsernameOrEmail, request.Body.Password)
	if err != nil {
		if err.Error() == "invalid credentials" {
			return LoginUser401JSONResponse{
				ErrorCode:    http.StatusUnauthorized,
				ErrorMessage: err.Error(),
			}, nil
		}
		return LoginUser400JSONResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}, nil
	}

	tokenStr, err := c.Jwt.GenerateToken(ctx, &uModel.DataJwt{
		Id:       strconv.Itoa(reqLogin.UserId),
		Username: reqLogin.Username,
		Email:    reqLogin.Email,
	})
	reqLogin.Jwt = &tokenStr

	return LoginUser200JSONResponse{
		Data:         &reqLogin,
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_SUCCESS_LOGIN,
	}, nil
}

// Register User
// (POST /register)
func (c *Controller) RegisterUser(ctx context.Context, request RegisterUserRequestObject) (RegisterUserResponseObject, error) {
	span, ctx := util.UpdateCtxSpanController(ctx)
	defer span.End()

	reqRegister, err := c.Users.Register(ctx, request.Body.Username, *request.Body.Email, request.Body.Password)
	if err != nil {
		return RegisterUser400JSONResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}, nil
	}

	return RegisterUser200JSONResponse{
		Data:         &reqRegister,
		ErrorCode:    http.StatusOK,
		ErrorMessage: constant.MSG_GENERAL_SUCCESS,
	}, nil
}

package impl

import (
	"base/internal/usecases/model"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (u *usecase) GenerateToken(ctx context.Context, data *model.DataJwt) (tokenStr string, err error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	now := time.Now()
	payload := &model.JwtMapClaims{
		DataJwt: model.DataJwt{
			Id:       data.Id,
			Username: data.Username,
			Email:    data.Email,
		},
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime), IssuedAt: jwt.NewNumericDate(now)},
	}
	token := func() *jwt.Token {
		var claims jwt.Claims = payload
		return &jwt.Token{Header: map[string]interface{}{"typ": "JWT", "alg": jwt.SigningMethod(jwt.SigningMethodHS256).Alg()}, Claims: claims, Method: jwt.SigningMethod(jwt.SigningMethodHS256)}
	}()
	secretKey := u.cfg.TokenSecret
	tokenStr, err = token.SignedString([]byte(secretKey))
	if err != nil {
		return tokenStr, errors.New("invalid token signed, " + err.Error())
	}

	return tokenStr, nil
}

func (u *usecase) GenerateTokenReset(ctx context.Context, data *model.DataJwtReset) (tokenStr string, err error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	now := time.Now()
	payload := &model.JwtMapClaimsReset{
		DataJwtReset: model.DataJwtReset{
			ID:        data.ID,
			ResetCode: data.ResetCode,
		},
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime), IssuedAt: jwt.NewNumericDate(now)},
	}
	token := func() *jwt.Token {
		var claims jwt.Claims = payload
		return &jwt.Token{Header: map[string]interface{}{"typ": "JWT", "alg": jwt.SigningMethod(jwt.SigningMethodHS256).Alg()}, Claims: claims, Method: jwt.SigningMethod(jwt.SigningMethodHS256)}
	}()
	secretKey := u.cfg.TokenSecret
	tokenStr, err = token.SignedString([]byte(secretKey))
	if err != nil {
		return tokenStr, errors.New("invalid token signed, " + err.Error())
	}

	return tokenStr, nil
}

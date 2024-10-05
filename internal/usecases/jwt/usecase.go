package jwt

import (
	"base/internal/usecases/model"
	"context"

	"github.com/golang-jwt/jwt/v4"
)

type Usecase interface {
	CheckToken(ctx context.Context, reqToken string) (claim jwt.MapClaims, isBool bool, err error)
	GetClaims(ctx context.Context, reqToken string) (claim model.DataJwt, err error)
	GenerateToken(ctx context.Context, data *model.DataJwt) (tokenStr string, err error)
	GenerateTokenReset(ctx context.Context, data *model.DataJwtReset) (tokenStr string, err error)
	CheckTokenReset(ctx context.Context, reqToken string) (claim jwt.MapClaims, isBool bool, err error)
	LogoutToken(ctx context.Context, token string) (err error)
	IsIdemBlacklisted(ctx context.Context, key string) bool
}

package interfaces

import (
	"base/internal/outbound/authentication/model"
	"context"
)

type Authentication interface {
	Login(ctx context.Context, req model.AuthRequest) (res model.AuthResponse, err error)
}

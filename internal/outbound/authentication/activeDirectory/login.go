package activedirectory

import (
	"base/internal/outbound/authentication/model"
	pkgHelper "base/pkg/helper"
	"context"

	sdkAuth "gitlab.banksinarmas.com/go/sdkv2/authentication/auth"
	aDirectoryAuth "gitlab.banksinarmas.com/go/sdkv2/authentication/integrations/activeDirectory"
)

func (r *repository) Login(ctx context.Context, req model.AuthRequest) (res model.AuthResponse, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanRepository(ctx)
	defer span.End()

	var address []aDirectoryAuth.Address

	for _, addr := range r.cfg.Auth.ActiveDirectory.Address {
		address = append(address, aDirectoryAuth.Address{
			Host: addr.Host,
			Port: addr.Port,
		})
	}

	authAd, err := aDirectoryAuth.New(aDirectoryAuth.Cfg{
		Network:             r.cfg.Auth.ActiveDirectory.Network,
		Username:            r.cfg.Auth.ActiveDirectory.Username,
		Password:            r.cfg.Auth.ActiveDirectory.Password,
		BaseDistinguishName: r.cfg.Auth.ActiveDirectory.BaseDistinguishName,
		Address:             address,
	})

	resAd, err := authAd.Login(ctx, sdkAuth.AuthRequest{
		Username: req.Username,
		Password: req.Password,
		AppName:  req.AppName,
	})

	res = model.AuthResponse{
		RoleList: resAd.RoleList,
		FullName: res.FullName,
	}

	return
}

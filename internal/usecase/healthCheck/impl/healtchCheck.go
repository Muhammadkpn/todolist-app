package impl

import (
	pkgHelper "base/pkg/helper"
	"base/pkg/resource/model"
	"context"
)

func (u *usecase) HealthCheck(ctx context.Context) (res []model.HealthCheckResponse, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	// database
	res = append(res, u.Db.Ping(ctx, u.cfg)...)

	// redis
	res = append(res, u.Redis.Ping(ctx, u.cfg)...)

	// add more depedencies if exist

	return
}

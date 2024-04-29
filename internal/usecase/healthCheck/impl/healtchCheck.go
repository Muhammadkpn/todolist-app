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
	hc := u.Db.Ping(ctx, u.cfg)
	res = append(res, hc...)

	// add more depedencies if exist

	return
}

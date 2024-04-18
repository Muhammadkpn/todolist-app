package petStore

import (
	"base/internal/outbound/model"
	pkgError "base/pkg/constant/error"
	"context"
	"fmt"

	sdkError "gitlab.banksinarmas.com/go/sdkv2/common/error"
	sdkLog "gitlab.banksinarmas.com/go/sdkv2/log/logger"

	"gorm.io/gorm"
)

// Create implements Repository.
func (r *repository) Create(ctx context.Context, db *gorm.DB, data model.Pet) (model.Pet, error) {
	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		r.resource.Logger.Error(ctx, fmt.Sprintf("%s - Error save", tag),
			sdkLog.WithAttr("data", data),
			sdkLog.WithError(err))

		if err == gorm.ErrRecordNotFound {
			return model.Pet{}, sdkError.New(pkgError.ERR_INVALID_ID, err)
		}
		return model.Pet{}, sdkError.New(err.Error(), err)
	}

	return data, nil
}

package util

import (
	"errors"

	"gorm.io/gorm"
)

func HandleErrorID(result *gorm.DB) (err error) {
	if result.RowsAffected == 0 || result.Error == gorm.ErrRecordNotFound {
		err = errors.New("data not found")
		return
	}

	err = HandleError(result)
	return
}

func HandleError(result *gorm.DB) (err error) {
	if result.Error != nil {
		err = errors.New("internal error")
		return
	}

	return
}

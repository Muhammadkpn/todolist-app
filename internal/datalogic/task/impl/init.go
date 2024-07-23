package impl

import (
	taskDatalogic "base/internal/datalogic/task"
	taskCache "base/internal/outbound/cache/task"
	taskDb "base/internal/outbound/db/task"
	pkgConfig "base/pkg/config"

	sdkLogger "gitlab.banksinarmas.com/go/sdkv2/log/logger"
)

type datalogic struct {
	TaskDb    taskDb.Repository
	TaskCache taskCache.Repository
}

func New(taskDb taskDb.Repository, log sdkLogger.Logger, cfg pkgConfig.Config, taskCache taskCache.Repository) taskDatalogic.Datalogic {
	return &datalogic{
		TaskDb:    taskDb,
		TaskCache: taskCache,
	}
}

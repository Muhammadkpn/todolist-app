package impl

import (
	taskDatalogic "base/internal/datalogic/task"
	"base/internal/outbound/authentication"
	taskDb "base/internal/outbound/db/task"
	interfaces "base/internal/outbound/interface"
	taskUsecase "base/internal/usecase/task"
	pkgConfig "base/pkg/config"

	sdkLogger "gitlab.banksinarmas.com/go/sdkv2/log/logger"
)

type usecase struct {
	TaskDb         taskDb.Repository
	AuthRepository map[string]interfaces.Authentication
	TaskDatalogic  taskDatalogic.Datalogic
}

func New(taskDb taskDb.Repository, log sdkLogger.Logger, cfg pkgConfig.Config, taskDatalogic taskDatalogic.Datalogic) taskUsecase.Usecase {
	return &usecase{
		TaskDb:         taskDb,
		AuthRepository: authentication.InitRepositories(cfg),
		TaskDatalogic:  taskDatalogic,
	}
}

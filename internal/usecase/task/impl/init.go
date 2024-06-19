package impl

import (
	"base/internal/outbound/authentication"
	"base/internal/outbound/db/task"
	interfaces "base/internal/outbound/interface"
	taskUsecase "base/internal/usecase/task"
	pkgConfig "base/pkg/config"

	sdkLogger "gitlab.banksinarmas.com/go/sdkv2/log/logger"
)

type usecase struct {
	TaskRepository task.Repository
	AuthRepository map[string]interfaces.Authentication
}

func New(taskRepository task.Repository, log sdkLogger.Logger, cfg pkgConfig.Config) taskUsecase.Usecase {
	return &usecase{
		TaskRepository: taskRepository,
		AuthRepository: authentication.InitRepositories(cfg),
	}
}

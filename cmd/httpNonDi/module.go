package main

import (
	exampleNonDiController "base/internal/inbound/http/exampleNonDi"
	exampleNonDiRepository "base/internal/outbound/db/exampleNonDi"
	exampleNonDiRepositoryImpl "base/internal/outbound/db/exampleNonDi/impl"
	exampleNonDiUsecase "base/internal/usecase/exampleNonDi"
	exampleNonDiUsecaseImpl "base/internal/usecase/exampleNonDi/impl"
	pkgConfig "base/pkg/config"
	"base/pkg/resource/injection"
	"base/pkg/resource/model"
)

var (
	// ** Resource **
	Config      pkgConfig.Config
	databaseSQL model.Database

	// ** Repository **
	repoA exampleNonDiRepository.Repository

	// ** Usecase **
	usecaseA exampleNonDiUsecase.Usecase
)

func initializeResource() {
	cfg, err := pkgConfig.NewConfig()
	if err != nil {
		panic("Error on init config")
	}
	Config = cfg

	db, err := injection.NewDatabase(cfg)
	if err != nil {
		panic("Error on init database")
	}
	databaseSQL = db
}

func initializeRepository() {
	repoA = exampleNonDiRepositoryImpl.New(databaseSQL, Config)
}

func initializeUsecase() {
	usecaseA = exampleNonDiUsecaseImpl.New(repoA)
}

func initializeController() {
	controllerA := exampleNonDiController.Controller{
		ExampleNonDiUsecase: usecaseA,
	}

	model.ControllerList = model.Controller{
		ControllerA: &controllerA,
	}
}

func InitModule() {
	initializeResource()
	initializeRepository()
	initializeUsecase()
	initializeController()
}

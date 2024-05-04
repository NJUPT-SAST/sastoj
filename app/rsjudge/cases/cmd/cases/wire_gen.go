// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/rsjudge/cases/internal/biz"
	"sastoj/app/rsjudge/cases/internal/conf"
	"sastoj/app/rsjudge/cases/internal/data"
	"sastoj/app/rsjudge/cases/internal/server"
	"sastoj/app/rsjudge/cases/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	casesRepo := data.NewJudgeRepo(dataData, logger)
	casesUsecase := biz.NewCasesUsecase(casesRepo, logger)
	casesService := service.NewCasesService(casesUsecase)
	grpcServer := server.NewGRPCServer(confServer, casesService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
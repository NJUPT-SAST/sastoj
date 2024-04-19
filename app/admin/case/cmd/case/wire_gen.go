// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"sastoj/app/admin/case/internal/biz"
	"sastoj/app/admin/case/internal/conf"
	"sastoj/app/admin/case/internal/data"
	"sastoj/app/admin/case/internal/server"
	"sastoj/app/admin/case/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	caseRepo := data.NewProblemCaseRepo(dataData, logger)
	caseUsecase := biz.NewCaseUsecase(caseRepo, logger)
	caseService := service.NewCaseService(caseUsecase)
	grpcServer := server.NewGRPCServer(confServer, caseService, logger)
	httpServer := server.NewHTTPServer(confServer, caseService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}

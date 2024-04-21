// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"sastoj/app/admin/contest/internal/biz"
	"sastoj/app/admin/contest/internal/conf"
	"sastoj/app/admin/contest/internal/data"
	"sastoj/app/admin/contest/internal/server"
	"sastoj/app/admin/contest/internal/service"

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
	contestRepo := data.NewContestRepo(dataData, logger)
	contestUsecase := biz.NewContestUsecase(contestRepo, logger)
	contestService := service.NewContestService(contestUsecase)
	grpcServer := server.NewGRPCServer(confServer, contestService, logger)
	httpServer := server.NewHTTPServer(confServer, contestService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}

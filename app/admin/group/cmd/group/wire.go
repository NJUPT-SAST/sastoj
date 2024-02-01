//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"sastoj/app/admin/group/internal/biz"
	"sastoj/app/admin/group/internal/conf"
	"sastoj/app/admin/group/internal/data"
	"sastoj/app/admin/group/internal/server"
	"sastoj/app/admin/group/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

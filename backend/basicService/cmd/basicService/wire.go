//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/LXJ0000/tok/backend/basicService/internal/biz"
	"github.com/LXJ0000/tok/backend/basicService/internal/data"
	"github.com/LXJ0000/tok/backend/basicService/internal/infra"
	"github.com/LXJ0000/tok/backend/basicService/internal/server"
	"github.com/LXJ0000/tok/backend/basicService/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
)

// wireApp init kratos application.
func wireApp(log.Logger, *minio.Core) (*kratos.App, func(), error) {
	panic(wire.Build(
		infra.ProviderSet,
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,

		newApp,
	))
}

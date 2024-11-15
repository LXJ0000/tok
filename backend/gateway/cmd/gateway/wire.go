//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/LXJ0000/tok/backend/gateway/internal/biz"
	"github.com/LXJ0000/tok/backend/gateway/internal/conf"
	"github.com/LXJ0000/tok/backend/gateway/internal/data"
	"github.com/LXJ0000/tok/backend/gateway/internal/server"
	"github.com/LXJ0000/tok/backend/gateway/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

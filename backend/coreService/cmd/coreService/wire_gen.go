// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/LXJ0000/tok/backend/coreService/internal/biz"
	"github.com/LXJ0000/tok/backend/coreService/internal/conf"
	"github.com/LXJ0000/tok/backend/coreService/internal/data"
	"github.com/LXJ0000/tok/backend/coreService/internal/server"
	"github.com/LXJ0000/tok/backend/coreService/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
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
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	userServiceService := service.NewUserServiceService(userUsecase)
	videoRepo := data.NewVideoRepo(dataData, logger)
	videoUsecase := biz.NewVideoUsecase(videoRepo, userRepo, logger)
	videoServiceService := service.NewVideoServiceService(videoUsecase)
	interRepo := data.NewInterRepo(dataData, logger)
	interUsecase := biz.NewInterUsecase(interRepo, logger)
	interServiceService := service.NewInterServiceService(interUsecase)
	grpcServer := server.NewGRPCServer(confServer, greeterService, userServiceService, videoServiceService, interServiceService, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, userServiceService, videoServiceService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}

package server

import (
	"os"

	v1 "github.com/LXJ0000/tok/backend/basicService/api/basicService/v1"
	"github.com/LXJ0000/tok/backend/basicService/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	accountService *service.AccountServiceService,
	fileService *service.FileServiceService,
	logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if os.Getenv("SERVICE_GRPC_ADDR") != "" {
		opts = append(opts, grpc.Address(os.Getenv("SERVICE_GRPC_ADDR")))
	}
	srv := grpc.NewServer(opts...)
	// v1.RegisterGreeterServer(srv, greeter)
	v1.RegisterAccountServiceServer(srv, accountService)
	v1.RegisterFileServiceServer(srv, fileService)
	return srv
}

package service

import (
	"net/http"

	pb "github.com/LXJ0000/tok/backend/coreService/api/coreService/v1"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService, NewUserServiceService, NewVideoServiceService, NewInterServiceService)

var (
	SuccessMeta = &pb.Metadata{
		BizCode: http.StatusOK,
		Message: "success",
	}
)

package service

import (
	"net/http"

	pb "github.com/LXJ0000/tok/backend/basicService/api/basicService/v1"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAccountServiceService)

var (
	SuccessMeta = &pb.Metadata{
		BizCode: http.StatusOK,
		Message: "success",
	}
)

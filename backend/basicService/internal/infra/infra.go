package infra

import (
	"github.com/LXJ0000/tok/backend/basicService/internal/infra/pkg/object_storage/minio"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(minio.NewMinio)

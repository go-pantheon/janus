package push

import (
	"github.com/google/wire"
	v1 "github.com/vulcan-frame/vulcan-gate/app/gate/internal/service/push/v1"
)

var ProviderSet = wire.NewSet(v1.NewPushService)

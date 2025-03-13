package push

import (
	v1 "github.com/go-pantheon/janus/app/gate/internal/service/push/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(v1.NewPushService)

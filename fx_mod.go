package redis_library

import (
	"github.com/neiasit/redis-library/core"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"redis_client",
	fx.Provide(
		core.LoadConfig,
		core.NewClient,
	),
)

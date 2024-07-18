package pkg

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"redis_client",
	fx.Provide(
		LoadConfig,
		NewClient,
	),
)

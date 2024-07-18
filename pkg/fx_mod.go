package pkg

import (
	"github.com/neiasit/redis-library/pkg/core"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"redis_client",
	fx.Provide(
		core.LoadConfig,
		core.NewClient,
	),
)

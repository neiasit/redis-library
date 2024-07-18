package pkg

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

func NewClient(cfg *Config, logger *slog.Logger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address(),
		Password: cfg.Password,
	})
	res := client.Ping(context.Background())
	err := res.Err()
	if err != nil {
		logger.Error("Failed to connect to redis", "address", cfg.Address(), "error", err)
		return nil, err
	}
	return client, nil
}

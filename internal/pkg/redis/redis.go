package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/me2seeks/cola/config"
	"github.com/me2seeks/cola/internal/pkg/logger"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Connect() *redis.Client {
	logger := logger.Logger

	Client = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", config.Cfg.Redis.Host, config.Cfg.Redis.Port),
		Password:    config.Cfg.Redis.Password,
		DB:          config.Cfg.Redis.DB,
		DialTimeout: time.Second * config.Cfg.Redis.DialTimeout,
		MaxRetries:  config.Cfg.Redis.MaxRetries,
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatalf("failed to connect to redis: %v", err)
		return nil
	}
	return Client
}

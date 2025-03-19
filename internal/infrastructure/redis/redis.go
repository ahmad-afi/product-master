package redis

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"product-master/internal/helper"
	"product-master/internal/utils"

	"github.com/go-redis/redis/v8"
)

type RedisConf struct {
	Password          string `env:"redis_password"`
	Port              int    `env:"redis_port"`
	Host              string `env:"redis_host"`
	DB                int    `env:"redis_db"`
	DefaultDB         int    `env:"redis_Defaultdb"`
	RedisMinIdleConns int    `env:"redis_MinIdleConns"`
	RedisPoolSize     int    `env:"redis_PoolSize"`
	RedisPoolTimeout  int    `env:"redis_PoolTimeout"`
	TLSConfig         bool   `env:"redis_TLSConfig"`
}

func NewRedisClient() *redis.Client {
	var redisConfig = RedisConf{
		Password:          utils.EnvString("redis_password"),
		Port:              utils.EnvInt("redis_port"),
		Host:              utils.EnvString("redis_host"),
		DB:                utils.EnvInt("redis_db"),
		RedisMinIdleConns: utils.EnvInt("redis_MinIdleConns"),
		RedisPoolSize:     utils.EnvInt("redis_PoolSize"),
		RedisPoolTimeout:  utils.EnvInt("redis_PoolTimeout"),
		TLSConfig:         utils.EnvBool("redis_TLSConfig"),
	}

	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		MinIdleConns: redisConfig.RedisMinIdleConns,
		PoolSize:     redisConfig.RedisPoolSize,
		PoolTimeout:  time.Duration(redisConfig.RedisPoolTimeout) * time.Second,
		Password:     redisConfig.Password,
		DB:           redisConfig.DB,
	})

	if redisConfig.TLSConfig {
		client.Options().TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		helper.Logger(helper.LoggerLevelPanic, "", fmt.Errorf("Cannot conenct to redis : %s", err.Error()))
		panic(err)
	}
	helper.Logger(helper.LoggerLevelInfo, fmt.Sprintf("Redis ping : %s", pong), nil)

	return client
}

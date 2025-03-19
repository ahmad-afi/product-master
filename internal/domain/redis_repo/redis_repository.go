package redis_repo

import (
	"context"
	"fmt"
	"product-master/internal/helper"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type redisUsersRepoImpl struct {
	redisClient *redis.Client
}

type RedisUsersRepository interface {
	GetData(ctx context.Context, key string) (res *string, err error)
	SetData(ctx context.Context, key, data string, ttl time.Duration) error
	DeleteData(ctx context.Context, key string) error
}

func NewRedisRepoUsers(redisClient *redis.Client) RedisUsersRepository {
	return &redisUsersRepoImpl{
		redisClient: redisClient,
	}
}

func (roi *redisUsersRepoImpl) GetData(ctx context.Context, key string) (res *string, err error) {

	// Get data from Redis
	helper.Logger(helper.LoggerLevelInfo, fmt.Sprintf("Get keys %s from redis\n", key), nil)
	result, err := roi.redisClient.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return
	}
	if result == "" {
		return nil, nil
	}

	res = &result
	return
}

func (roi *redisUsersRepoImpl) SetData(ctx context.Context, key, data string, ttl time.Duration) error {
	if err := roi.redisClient.Set(ctx, key, data, ttl).Err(); err != nil {
		return errors.Wrap(err, "redisRepo.SetUsersCtx.redisClient.set")
	}

	helper.Logger(helper.LoggerLevelInfo, fmt.Sprintf("Set keys %s to redis with TTL %d \n", key, ttl), nil)
	return nil
}

func (roi *redisUsersRepoImpl) DeleteData(ctx context.Context, key string) (err error) {
	if err := roi.redisClient.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "redisRepo.DeleteData.redisClient.Del")
	}
	helper.Logger(helper.LoggerLevelInfo, fmt.Sprintf("Delete keys %s from redis succeed\n", key), nil)
	return nil
}

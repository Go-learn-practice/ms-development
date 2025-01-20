package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"test.com/devCommon/logs"
	"test.com/devUser/config"
	"time"
)

var (
	RedisCacheInstance *RedisCache
)

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	logs.LG.Info("初始化 redis 缓存连接")
	// 读取 redis 配置
	rdb := redis.NewClient(config.Conf.ReadRedisConfig())
	RedisCacheInstance = &RedisCache{
		rdb: rdb,
	}
}

func (redisCache *RedisCache) Put(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := redisCache.rdb.Set(ctx, key, value, expiration).Err()
	return err
}

func (redisCache *RedisCache) Get(ctx context.Context, key string) (interface{}, error) {
	result, err := redisCache.rdb.Get(ctx, key).Result()
	return result, err
}

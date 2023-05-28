package cache

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type IRedisConfig interface {
	RedisAddress(context.Context) (string, error)
	RedisPassword(context.Context) (string, error)
}

type redisCache struct {
	rdb *redis.Client
}

func NewRedis(ctx context.Context, cnf IRedisConfig, db int) (*redisCache, error) {
	addr, err := cnf.RedisAddress(ctx)
	if err != nil {
		return nil, err
	}

	pwd, err := cnf.RedisPassword(ctx)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	return &redisCache{rdb: rdb}, nil
}

func (c *redisCache) Get(ctx context.Context, key string, out protoreflect.ProtoMessage) error {
	value, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	return proto.Unmarshal(data, out)
}

func (c *redisCache) Set(ctx context.Context, key string, val protoreflect.ProtoMessage, expiration time.Duration) error {
	data, err := proto.Marshal(val)
	if err != nil {
		return err
	}
	result := base64.StdEncoding.EncodeToString(data)
	return c.rdb.SetEx(ctx, key, result, expiration).Err()
}

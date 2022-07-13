package common

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Option struct {
	Addr        string
	Password    string
	DB          int
	PoolSize    int
	MinIdleConn int
}

type Redis struct {
	Client *redis.Client
}

var (
	ctx = context.Background()
)

func (opt *Option) NewRedis() *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:         opt.Addr,
		Password:     opt.Password,
		DB:           opt.DB,
		PoolSize:     opt.PoolSize,
		MinIdleConns: opt.MinIdleConn,
	})
	return &Redis{
		Client: rdb,
	}
}

func (c *Redis) Set(key string, value interface{}, ttl time.Duration) {
	err := c.Client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		panic(err)
	}
}

func (c *Redis) Get(key string) string {
	val, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func (c *Redis) Scan() (uint64, []string) {
	var cursor uint64
	keys, cursor, err := c.Client.Scan(ctx, cursor, "*", 0).Result()
	if err != nil {
		panic(err)
	}
	return cursor, keys
}

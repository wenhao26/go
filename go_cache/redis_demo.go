package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:         "192.168.7.209:6379",
		Password:     "",
		DB:           7,
		PoolSize:     200,
		MinIdleConns: 100,
	})

	for {
		go func() {
			val, err := rdb.Get(ctx, "coinsky:diamond_mall:goods_sales:2").Result()
			if err != nil {
				panic(err)
			}
			fmt.Println("coinsky:diamond_mall:goods_sales:2", val)
		}()

		fmt.Println(1)
		time.Sleep(1e9)
	}

}

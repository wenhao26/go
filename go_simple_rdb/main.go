package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
)

var wg sync.WaitGroup
var rdb *redis.Client
var ctx = context.Background()

func init() {
	/*ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()*/

	rdb = redis.NewClient(&redis.Options{
		Addr:         "192.168.7.209:6379",
		Password:     "",
		DB:           7,
		PoolSize:     50,
		MinIdleConns: 10,
	})
	/*_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}*/
}

func main() {
	var cursor uint64
	for {
		keys, cursor, err := rdb.Scan(ctx, cursor, "*", 10).Result()
		fmt.Println("游标值：", cursor)
		if err != nil {
			fmt.Println("扫描键列表失败：", err)
			return
		}
		if cursor == 0 {
			break
		}

		for _, key := range keys {
			keyType, err := rdb.Type(ctx, key).Result()
			if err != nil {
				fmt.Println("获取键类型失败：", err)
				return
			}
			fmt.Printf("key：%v  type：%v\n", key, keyType)

			if keyType == "string" {
				val, err := rdb.Get(ctx, key).Result()
				if err != nil {
					fmt.Println("获取 string 值失败：", err)
					return
				}
				fmt.Printf("key：%v  value：%v\n", key, val)
			} else if keyType == "list" {
				val, err := rdb.LPop(ctx, key).Result()
				if err != nil {
					fmt.Println("获取 list 值失败：", err)
					return
				}
				fmt.Printf("key：%v  value：%v\n", key, val)
			}

			// todo more...
		}
		fmt.Println("=====================================================\n")
		/*wg.Add(1)
		go func() {
			defer wg.Done()

			for _, key := range keys {
				keyType, err := rdb.Type(ctx, key).Result()
				if err != nil {
					fmt.Println("获取键类型失败：", err)
					return
				}
				fmt.Printf("key：%v  type：%v\n", key, keyType)

				if keyType == "string" {
					val, err := rdb.Get(ctx, key).Result()
					if err != nil {
						fmt.Println("获取 string 值失败：", err)
						return
					}
					fmt.Printf("key：%v  value：%v\n", key, val)
				} else if keyType == "list" {
					val, err := rdb.LPop(ctx, key).Result()
					if err != nil {
						fmt.Println("获取 list 值失败：", err)
						return
					}
					fmt.Printf("key：%v  value：%v\n", key, val)
				}

				// todo more...
			}
			fmt.Println("=====================================================\n")

		}()
		wg.Wait()*/
	}

	fmt.Println("执行结束...")
}

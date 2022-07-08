package main

import (
	"fmt"
	"go_project/go/go_learn/common"
)

func main() {
	opt := common.Option{
		Addr:        "192.168.7.209:6379",
		Password:    "",
		DB:          0,
		PoolSize:    200,
		MinIdleConn: 100,
	}
	redis := opt.NewRedis()
	/*redis.Set("test", "123456", 30e9)
	result := redis.Get("test")
	fmt.Println(result)*/

	cursor, keys := redis.Scan()
	fmt.Println("游标=", cursor)
	fmt.Println("Keys=", keys)

}

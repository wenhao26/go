package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"log"
	"reflect"
	"time"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config failed:", err)
		return
	}

	host := reflect.ValueOf(viper.Get("redis.host")).Interface().(string)
	port := reflect.ValueOf(viper.Get("redis.port")).Interface().(int64)
	password := reflect.ValueOf(viper.Get("redis.password")).Interface().(string)
	db := reflect.ValueOf(viper.Get("redis.db")).Interface().(int64)

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       int(db),
	})

	// Set 设置一个key的值
	/*for i := 1; i <= 3; i++ {
		if err := client.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i), 0).Err(); err != nil {
			panic(err)
		}
	}*/

	// SetNx 如果key不存在，则设置这个key的值
	/*if err := client.SetNX("key2", "value2", time.Second*60).Err(); err != nil {
		panic(err)
	}*/

	// MGet 批量查询key的值
	/*values, err := client.MGet("key1", "key2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(values)*/

	// MSet 批量设置key的值
	/*if err := client.MSet("key4", "value4", "key5", "value5").Err(); err != nil {
		panic(err)
	}*/

	// Incr,IncrBy 针对一个key的数值进行递增操作
	/*// Incr函数每次加一
	val, err := client.Incr("key6").Result()
	if err != nil {
		panic(err)
	}*/
	/*// IncrBy函数，可以指定每次递增多少
	val, err := client.IncrBy("key6", 2).Result()
	if err != nil {
		panic(err)
	}
	*/
	// IncrByFloat函数，可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
	/*val, err := client.IncrByFloat("key6", 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)*/

	// GetSet 设置一个key的值，并返回这个key的旧值
	/*val, err := client.GetSet("my_key", "my_value2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)*/

	// Get 查询key的值
	/*val, err = client.Get("my_key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)*/

	// Decr,DecrBy 针对一个key的数值进行递减操作
	/*// Decr函数每次减一
	val, err := client.Decr("key6").Result()
	if err != nil {
		panic(err)
	}*/
	// DecrBy函数，可以指定每次递减多少
	/*val, err := client.DecrBy("key6", 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)*/

	// Del 删除key操作，支持批量删除
	/*if err := client.Del("key1", "key2").Err(); err != nil {
		panic(err)
	}*/

	// Expire 设置已存在key的过期时间,单位秒
	client.Expire("key3", time.Second * 10)

	fmt.Println("main end")
}

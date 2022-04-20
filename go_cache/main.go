package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	// 设置超时时间和清理时间
	c := cache.New(5*time.Minute, 10*time.Minute)

	c.Set("host", "https://example.me", 120)

	value, ok := c.Get("host")
	if ok {
		fmt.Println("host=", value)
	}
}

package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
	"unsafe"
)

type PSubscribeCallback func(pattern, channel, message string)

type PSubscriber struct {
	client redis.PubSubConn
	cbMap  map[string]PSubscribeCallback
}

func (c *PSubscriber) PConnect(address string, password string) {
	conn, err := redis.Dial("tcp", address, redis.DialPassword(password))
	if err != nil {
		log.Println("redis dial failed.", err)
	}

	c.client = redis.PubSubConn{conn}
	c.cbMap = make(map[string]PSubscribeCallback)

	go func() {
		for {
			log.Println("wait...")
			switch res := c.client.Receive().(type) {
			case redis.PMessage:
				pattern := (*string)(unsafe.Pointer(&res.Pattern))
				channel := (*string)(unsafe.Pointer(&res.Channel))
				message := (*string)(unsafe.Pointer(&res.Data))
				c.cbMap[*channel](*pattern, *channel, *message)
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", res.Channel, res.Kind, res.Count)
			case error:
				log.Println("error handle...")
				continue
			}
		}
	}()
}

func (c *PSubscriber) PSubscribe(channel interface{}, cb PSubscribeCallback) {
	err := c.client.PSubscribe(channel)
	if err != nil {
		log.Println("redis Subscribe error.")
	}

	c.cbMap[channel.(string)] = cb
}

func timeoutCallback(patter, channel, device string) {
	log.Println("timeoutCallback patter : "+patter+" channel : ", channel, " offline device : ", device)
}

func main() {
	var psub PSubscriber
	psub.PConnect("11*****.245:6379", "********")
	psub.PSubscribe("__keyevent@0__:expired", timeoutCallback)
	for {
		time.Sleep(1 * time.Second)
	}
}

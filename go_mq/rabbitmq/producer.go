package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"go_project/go/go_mq/rabbitmq/common"
	"log"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 连接
	conn, err := amqp.Dial(common.URL)
	common.ErrorMessage(err, ">>>> 连接服务失败")

	// 建立通道
	ch, err := conn.Channel()
	common.ErrorMessage(err, ">>>> 打开通道失败")

	// 声明队列
	q, err := ch.QueueDeclare(common.QUEUE, true, false, false, false, nil)
	common.ErrorMessage(err, ">>>> 未能声明队列")

	// 发布消息
	wg.Add(1)
	go func() {
		for i := 1; i <= 500000; i++ {
			fmt.Println("goroutine", i)

			msgBody := "test data id-" + strconv.Itoa(i)
			err = ch.Publish("", q.Name, false, false, amqp.Publishing{
				DeliveryMode: amqp.Persistent, // 消息持久化
				ContentType:  "text/plain",
				Body:         []byte(msgBody),
			})
			log.Printf(" [x] Sent %s", msgBody)
			common.ErrorMessage(err, ">>>> 发布消息失败")
		}
		defer wg.Done()
	}()
	wg.Wait()

	fmt.Println("done")
}

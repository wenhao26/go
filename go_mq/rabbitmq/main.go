package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"go_project/go/go_mq/rabbitmq/simple"
	"time"
)

func main() {
	conn := simple.GetClient()
	ch := simple.Channel(conn)
	defer ch.Close()

	err := ch.ExchangeDeclare("app.test.exchange", "direct", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for {
			body := "time=" + time.Now().Format("2006-01-02 15:04:05")
			err = ch.Publish("app.test.exchange", "app.test.routing_key", false, false, amqp.Publishing{
				// DeliveryMode: amqp.Persistent, // 消息持久化
				ContentType: "text/plain",
				Body:        []byte(body),
			})
			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf(" [x] Sent %s\n", body)
		}
	}()
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

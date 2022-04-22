package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"go_project/go/go_mq/rabbitmq/common"
	"log"
)

func main() {
	conn, err := amqp.Dial(common.URL)
	common.ErrorMessage(err, ">>>> 连接服务失败")
	defer conn.Close()

	forever := make(chan bool)
	/*for routine := 1; routine <= common.CONSUMERCNT; routine++ {
		go func(routineNum int) {
			ch, err := conn.Channel()
			common.ErrorMessage(err, ">>>> 打开通道失败")
			defer ch.Close()

			q, err := ch.QueueDeclare(
				common.QUEUE,
				true,
				false,
				false,
				false,
				nil,
			)
			common.ErrorMessage(err, ">>>> 未能声明队列")

			msgData, err := ch.Consume(
				q.Name,
				common.EXCHANGE,
				false, // Auto Ack
				false,
				false,
				false,
				nil,
			)

			if err != nil {
				log.Fatal(err)
			}

			for msg := range msgData {
				log.Printf("In %d consume a message: %s\n", routineNum, msg.Body)
				log.Printf("Done")
				_ = msg.Ack(false) // Ack
			}
		}(routine)
	}*/

	// new
	ch, err := conn.Channel()
	common.ErrorMessage(err, ">>>> 打开通道失败")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		common.QUEUE,
		true,
		false,
		false,
		false,
		nil,
	)
	common.ErrorMessage(err, ">>>> 未能声明队列")

	go func() {
		msgData, err := ch.Consume(
			q.Name,
			common.EXCHANGE,
			false, // Auto Ack
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatal(err)
		}
		for msg := range msgData {
			log.Println("-goroutine1:", string(msg.Body))
			_ = msg.Ack(false) // Ack
		}
	}()

	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

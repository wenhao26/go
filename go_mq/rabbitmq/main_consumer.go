package main

import (
	"fmt"
	"go_project/go/go_mq/rabbitmq/simple"
	"log"
)

func main() {
	conn := simple.GetClient()
	ch := simple.Channel(conn)
	defer ch.Close()

	forever := make(chan bool)
	go func() {
		msgData, err := ch.Consume("app.test.queue", "app.test.consumer", false, false, false, false, nil)
		if err != nil {
			fmt.Println(err)
		}
		for msg := range msgData {
			log.Println("-goroutine1:", string(msg.Body))
			_ = msg.Ack(false) // Ack
		}
	}()
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

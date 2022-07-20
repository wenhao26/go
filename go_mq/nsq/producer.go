package main

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	topicName := "topic1"

	for {
		t := time.Now().String()
		messageBody := []byte("message:" + t)
		err = producer.Publish(topicName, messageBody)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Sent successfully")
		time.Sleep(time.Millisecond * 1)
	}

	producer.Stop()
}

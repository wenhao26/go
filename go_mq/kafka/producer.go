package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	producer, err := sarama.NewSyncProducer([]string{"192.168.7.41:9092", "192.168.7.106:9092", "192.168.7.109:9092"}, config)
	if err != nil {
		panic(err)
		return
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "test_topic",
		Key:   sarama.StringEncoder("test"),
		Value: sarama.StringEncoder("test,true"),
	}
	message, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("Partition:", message, ";Offset:", offset)
}

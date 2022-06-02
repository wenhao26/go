package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_1_1_0
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	client, err := sarama.NewClient([]string{"192.168.7.41:9092", "192.168.7.106:9092", "192.168.7.109:9092"}, config)
	if err != nil {
		panic(err)
		return
	}
	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
		return
	}
	defer consumer.Close()

	topic := "test"
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("无法获取分区列表：", err)
	}
	fmt.Println(partitionList)

	// 循环读取分区
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("无法启动分区[%d]的使用者：%s\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		for msg := range pc.Messages() {
			fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		}
	}
}

package main

import (
	"fmt"
	"strconv"
	"time"
)

func task1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
		time.Sleep(1e9)
	}
}

func task2(ch chan string) {
	for i := 0; ; i++ {
		ch <- strconv.Itoa(i + 4)
		time.Sleep(1e9)
	}
}

func pull(ch1 chan int, ch2 chan string) {
	t := time.Tick(3e9)
	for {
		select {
		case val := <-ch1:
			fmt.Printf("Received on channel-1：%d\n", val)
		case val := <-ch2:
			fmt.Printf("Received on channel-2：%s\n", val)
		case <-t:
			fmt.Println("Log...")
		}
	}
}

func main() {
	// 生产者消费者模型
	ch1 := make(chan int)
	ch2 := make(chan string)

	go task1(ch1)
	go task2(ch2)
	go pull(ch1, ch2)

	time.Sleep(30e9)
	fmt.Println("main end")
}

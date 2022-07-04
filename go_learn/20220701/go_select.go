package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	intChannels := [5]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(5)
	fmt.Printf("The index :%d\n", index)
	intChannels[index] <- index

	select {
	case <-intChannels[0]:
		fmt.Println("1...")
	case <-intChannels[1]:
		fmt.Println("2...")
	case elem := <-intChannels[2]:
		fmt.Println("3...", elem)
	default:
		fmt.Println("other...")

	}

	http.ListenAndServe("0.0.0.0:6060", nil)
}

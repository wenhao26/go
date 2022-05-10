package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg2 sync.WaitGroup

func t1() {
	defer wg2.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A：%d\n", i)
	}
}

func t2() {
	defer wg2.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B：%d\n", i)
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	wg2.Add(2)
	go t1()
	go t2()
	wg2.Wait()
}

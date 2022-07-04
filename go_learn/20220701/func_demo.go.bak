package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
	_ "net/http/pprof"
)

type Community struct {
	Symbol string
}

type MethodUtils struct {
}

var (
	ch = make(chan int, 5)
	wg sync.WaitGroup
)

func (c Community) showName() {
	fmt.Println("货币符号为：", c.Symbol)
}

func asyncLoop(loop int) {
	wg.Add(loop)
	for i := 0; i < loop; i++ {
		cur_i := i
		go func() {
			ch <- 1
			fmt.Println(cur_i)
			<-ch
			wg.Done()
			time.Sleep(5 * time.Second)
		}()
	}
	wg.Wait()
}

func (mu *MethodUtils) matrix(x, y int) {
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			fmt.Print("x")
		}
		fmt.Println()
	}
}

func main() {
	//var c Community
	//c.Symbol = "BTC"
	//c.showName()

	//asyncLoop(15)
	var mu MethodUtils
	mu.matrix(10, 8)

	http.ListenAndServe("0.0.0.0:6060", nil)
}

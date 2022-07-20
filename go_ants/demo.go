package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/panjf2000/ants/v2"
)

var sum2 int32

func job(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum2, n)
	fmt.Printf("run with %d\n", n)
}

func main() {
	defer ants.Release()

	runTimes := 100

	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(100, func(i interface{}) {
		job(i)
		wg.Done()
	})
	defer p.Release()
	// 逐个提交任务
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum2)

}

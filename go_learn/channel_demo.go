package main

import (
	"time"
)

func makeStream(n int) chan bool {
	ch := make(chan bool, n)
	go func() {
		for i := 0; i < n; i++ {
			ch <- true
		}
		close(ch)
	}()
	return ch
}

func doWork() {
	time.Sleep(4e9)
	panic("fk")
}

func main() {
	/*stream := makeStream(5)
	for {
		v, ok := <-stream
		if !ok {
			break
		}
		fmt.Println(v)
	}*/

	/*ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- fmt.Sprintf("A%d", i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- fmt.Sprintf("B%d", i)
		}
	}()

	go func() {
		for {
			select {
			case v := <-ch1:
				fmt.Println(v)
			case v := <-ch2:
				fmt.Println(v)
			}
		}
	}()

	time.Sleep(10e9)*/

	/*go func() {
		for {
			log.Printf("another worker")
			time.Sleep(1e9)
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("出问题了 %s", err)
			}
		}()

		doWork()
	}()

	time.Sleep(10e9)*/






}

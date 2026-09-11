package main

import (
	"fmt"
	"runtime"
	"sync"
)

type count struct {
	num   int
	mutex sync.Mutex
}

func (c *count) increment() {
	c.mutex.Lock()
	c.num += 1
	c.mutex.Unlock()
}

func (c *count) result() {
	fmt.Println(c.num)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}
	done := make(chan bool)

	// 송신(Send) 전용 고루틴 생성
	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // CPU 양보
		}()
	}

	// 수신(Receive) 전용 고루틴 생성
	for i := 1; i <= 10000; i++ {
		<- done
	}

	c.result()
}
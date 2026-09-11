package main

import (
	"fmt"
	"runtime"
)

type count struct {
	num int
}

func (c *count) increment() {
	c.num += 1
}

func (c *count) result() {
	fmt.Println(c.num)
}

// 경쟁 상태(Race Condition)으로 인한 정상적인 값을 유추하지 못함
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

	c.result();
}
package main

import (
	"fmt"
	"time"
)

// 발신 전용
func sendOnly(ch chan <- int, cnt int) {
	for i := 0; i < cnt; i++ {
		ch <- i
	}

	ch <- 777
}

// 수신 전용
func receiveOnly(ch <- chan int) {
	for i := range ch {
		fmt.Println("received :", i)
	}

	fmt.Println(<- ch)
}

func main() {
	ch := make(chan int)

	go sendOnly(ch, 10)
	go receiveOnly(ch)

	time.Sleep(2 * time.Second)
}
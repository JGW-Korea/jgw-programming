package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	// 값 수신 용도
	go func() {
		for {
			num := <- ch1
			fmt.Println("ch1 : ", num)
			time.Sleep(250 * time.Millisecond)
		}
	}()
	
	go func() {
		for {
			ch2 <- "Golang Hi1"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
				case ch1 <- 777: // 값 송신 용도
				case str := <- ch2: {
					fmt.Println("ch2 : ", str)
				}
				// 주의 -> 채널이 수신되지 않으면 모두 default로 빠져버리게 됨
				// default: {
				// 	fmt.Println("default test")
				// }
			}
		}
	}()

	time.Sleep(7 * time.Second)
}
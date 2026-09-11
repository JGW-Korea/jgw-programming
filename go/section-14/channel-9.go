package main

import (
	"fmt"
)

// 채널(Channel)을 반환하는 함수
func returnChannel(cnt int) <- chan int {
	sum := 0
	tot := make(chan int)

	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}

		tot <- sum
		tot <- 777
		tot <- 7777

		close(tot)
	}()

	return tot
}

// 수신 전용 채널(Receive-only Channel) 및 채널 반환(Return Channel)
func total(c <- chan int) <- chan int {
	tot := make(chan int)

	go func() {
		a := 0

		for i := range c {
			a += i
		}

		tot <- a
	}()

	return tot
}

func main() {
	c := returnChannel(10)
	output := total(c)

	fmt.Println("output", <- output)
}
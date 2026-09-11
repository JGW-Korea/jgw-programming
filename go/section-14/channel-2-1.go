package main

import (
	"fmt"
	"time"
)

func rangeSum(rg int, c chan int) {
	fmt.Printf("Range is %d : S ---> %v\n", rg, time.Now())
	
	sum := 0
	
	for i := 1; i <= rg; i++ {
		sum += i
	}
	
	c <- sum

	fmt.Printf("Range is %d : E ---> %v\n", rg, time.Now())
}

func main() {
	c := make(chan int, 3)

	go rangeSum(1000, c)
	go rangeSum(7000, c)
	go rangeSum(5000, c)

	// 순서대로 데이터 수신(동기화) : 채널에서 값 수신 완료 될 떄까지 대기
	result1 := <- c
	// result2 := <- c
	// result3 := <- c

	fmt.Println("Result 1 ---> ", result1)
	// fmt.Println("Result 2 ---> ", result2)
	// fmt.Println("Result 3 ---> ", result3)
}
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
	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(7000, c)
	go rangeSum(5000, c)

	// 순서대로 데이터 수신(동기화) : 채널에서 값 수신 완료 될 떄까지 대기
	result1 := <- c
	result2 := <- c
	result3 := <- c
	result4 := <- c
	// result5 := <- c
	// result6 := <- c
	// result7 := <- c
	// result8 := <- c

	fmt.Println("Result 1 ---> ", result1)
	fmt.Println("Result 2 ---> ", result2)
	fmt.Println("Result 3 ---> ", result3)
	fmt.Println("Result 4 ---> ", result4)
	// fmt.Println("Result 5 ---> ", result5)
	// fmt.Println("Result 6 ---> ", result6)
	// fmt.Println("Result 7 ---> ", result7)
	// fmt.Println("Result 8 ---> ", result8)
	// fmt.Println("Result 2 ---> ", result2)
	// fmt.Println("Result 3 ---> ", result3)
}
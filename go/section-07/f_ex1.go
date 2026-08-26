package main

import "fmt"

// 콜백 함수 #1. 기본 사용 방법
func add(x int, y int) int {
	return x + y
}

func sum(x int, f func(int, int)int) int {
	return f(x, 10)
}

func main() {
	result1 := sum(23, add) // 콜백 함수 #1. 기본 사용 방법
	
	// 콜백 함수 #2. 함수 표현식(Function Expression) 전달
	result2 := sum(23, func(x int, y int)int{
		return x + y
	})

	fmt.Println(result1, result2)
}
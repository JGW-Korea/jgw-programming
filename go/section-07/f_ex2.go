package main

import "fmt"

// 고차 함수 #1. 함수를 매개변수로 전달받는 고차함수
func sum(x int, f func(int, int)int) int {
	return f(x, 10)
}

// 고차 함수 #2. 함수를 반환하는 고차함수
func funcA() func(int, int)int {
	return func(x, y int)int{
		return x + y
	}
}

func main() {
	add := funcA()

	fmt.Println("sum:", add(10, 20))
}
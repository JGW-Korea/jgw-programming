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

func funcC(x *int) {
	*x = 20
	fmt.Printf("%p, %d\n", &x, *x)
}

func main() {
	add := funcA()

	var str string = "string";

	fmt.Println(len(str))

	x := 10
	fmt.Printf("%p, %d\n", &x, x)
	funcC(&x)
	fmt.Printf("%p, %d\n", &x, x)

	fmt.Println("sum:", add(10, 20))
}
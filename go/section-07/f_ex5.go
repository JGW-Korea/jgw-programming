package main

import "fmt"

func pointerParameter(n *int) {
	fmt.Printf("값에 의한 전달(before) #2. [%p] %d\n", &n, *n)
	fmt.Printf("데이터 주소(before) #3. [%p] %d\n", &(*n), *n)
	
	*n = 77
	
	fmt.Printf("값에 의한 전달(after) #4. [%p] %d\n", &n, *n)
	fmt.Printf("데이터 주소(after) #5. [%p] %d\n", &(*n), *n)
}

func main() {
	var x int = 30

	fmt.Printf("데이터 주소 #1. [%p] %d\n", &x, x)
	
	pointerParameter(&x)
	
	fmt.Printf("데이터 주소 #6. [%p] %d\n", &x, x)
}

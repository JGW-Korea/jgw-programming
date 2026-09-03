package main

import "fmt"

type test interface {} // 빈 인터페이스(Empty Interface)

func main() {
	var t test
	fmt.Println("ex 1 : ", t) // Nil 리턴
}
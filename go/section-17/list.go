package main

import (
	"fmt"
)

func funcA(param *int) {
	fmt.Printf("%p => %d\n", &param, *param)
}

func main() {
	// list.New().Back()

	a := 10;

	fmt.Printf("%p => %d\n", &a, a)
	funcA(&a)
}
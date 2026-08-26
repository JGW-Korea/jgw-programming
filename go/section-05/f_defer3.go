package main

import "fmt"

func stack() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("ex1:", i + 1)
	}
}

func main() {
	stack()
}
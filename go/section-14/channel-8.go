package main

import "fmt"

func sum(cnt int) <- chan int {
	sum := 0
	tot := make(chan int)

	go func() {
		for i := 1; i < cnt; i++ {
			sum += i
		}

		tot <- sum
	}()
	
	return tot
}

func main() {
	ch := sum(100)

	fmt.Println("ex1:", <- ch)
}
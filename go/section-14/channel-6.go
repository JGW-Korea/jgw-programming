package main

import "fmt"

func main() {
	ch := make(chan string)
	
	go func() {
		for i := 0; i < 3; i++ {
			ch <- "Good!"
		}
	}()
	
	value1, ok1 := <- ch
	fmt.Println("ex1 : ", value1, ok1)
	
	value2, ok2 := <- ch
	fmt.Println("ex2 : ", value2, ok2)
	
	value3, ok3 := <- ch
	fmt.Println("ex3 : ", value3, ok3)

	close(ch)

	value4, ok4 := <- ch
	fmt.Println("ex4 : ", value4, ok4)
}
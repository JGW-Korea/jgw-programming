package main

import "fmt"

func funcA() {
	var c, d int = 30, 40
	
	// ...
}

func main() {
	var a, b int = 10, 20
	
	funcA()
	
	fmt.Printf("%d / %d / %d / %d", a, b, c, d)
}
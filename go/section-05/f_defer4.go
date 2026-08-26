package main

import "fmt"

func start(str string) string {
	fmt.Println("start:", str)
	return str
}

func end(str string) {
	fmt.Println("end:", str)
}

func a() {
	defer end(start("Hello"))
	fmt.Println("in a")
}

func main() {
	a()
}
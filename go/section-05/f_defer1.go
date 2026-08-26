package main

import "fmt"

func ex_f2() {
	fmt.Println("f2 : called")
}

func ex_f1() {
	fmt.Println("f1 : start!")
	defer ex_f2()
	fmt.Println("f1 : edn")
}

func main() {
	ex_f1()	
}
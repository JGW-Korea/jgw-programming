package main

import "fmt"

var GLOBAL_VARIABLES int = 10

func init() {
	fmt.Println(GLOBAL_VARIABLES)
	fmt.Println(GLOBAL_CONSTANT)
}

func callee() {
	fmt.Println("Callee")
}

const GLOBAL_CONSTANT int = 30
package main

import (
	"fmt"
	"time"
)

func exe(step int) {
	fmt.Printf("exe%d func start -> %v\n", step, time.Now())
	time.Sleep(1 * time.Second)
	fmt.Printf("exe%d func end -> %v\n", step, time.Now())
}

func main() {
	exe(1)
	
	fmt.Println("Main Routine Start", time.Now())
	go exe(2)
	go exe(3)
	time.Sleep(5 * time.Second)
	fmt.Println("Main Routine End", time.Now())
}
package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Printf("exe%d func start -> %v\n", 1, time.Now())
	time.Sleep(1 * time.Second)
	fmt.Printf("exe%d func end -> %v\n", 1, time.Now())
}

func exe2() {
	fmt.Printf("exe%d func start -> %v\n", 2, time.Now())
	time.Sleep(1 * time.Second)
	fmt.Printf("exe%d func end -> %v\n", 2, time.Now())
}

func exe3() {
	fmt.Printf("exe%d func start -> %v\n", 3, time.Now())
	time.Sleep(1 * time.Second)
	fmt.Printf("exe%d func end -> %v\n", 3, time.Now())
}

func main() {
	exe1()
	
	fmt.Println("Main Routine Start", time.Now())
	go exe2()
	go exe3()
	time.Sleep(5 * time.Second)
	fmt.Println("Main Routine End", time.Now())
}
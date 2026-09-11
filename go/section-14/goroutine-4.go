package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func exe(name int) {
	var random int = rand.Intn(100)

	fmt.Println(name, "start -> ", time.Now())

	for i := 0; i < 10; i++ {
		fmt.Println(name, ">>>>>>", random, i)
	}

	fmt.Println(name, "end -> ", time.Now())
}

func main() {
	runtime.GOMAXPROCS(1)

	fmt.Println("Main Routine Start:", time.Now())
	for i := 0; i < 5; i++ {
		go exe(i) // 고루틴 100개 생성
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Main Routine end:", time.Now())
}
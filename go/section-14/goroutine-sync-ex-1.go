package main

import (
	"fmt"
	"sync"
	"time"
)

func onceTest() {
	// 이 부분에 한 번 실행할 코드 작성
	fmt.Println("Once Test Excute!")
}

func main() {
	// 고루틴 동기화 고급
	// Once -> 한 번만 실행 (주로 초기화에 사용)
	// DO로 실행

	once := new(sync.Once)

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println("Goroutine : ", n)
			once.Do(onceTest)
		}(i)
	}

	time.Sleep(2 * time.Second)
}
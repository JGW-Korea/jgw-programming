package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5)  // 버퍼 채널(Buffered Channel)

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Wating :", n)
			condition.Wait()                  // 고루틴 대기 (하위 코드가 실행되지 않음)
			fmt.Println("Wating End : ", n)   // 고루틴 대기가 풀린 이후 실행됨
			mutex.Unlock()
		}(i)
	}
	
	for i := 0; i < 5; i++ {
		<- c
		// fmt.Println("received : ", <- c)
	}
	
	// 대기 중인 고루틴을 하나씩 깨우는 방법
	// for i := 0; i < 5; i++ {
	// 	mutex.Lock()
		
	// 	fmt.Println("Wake Goroutine(Signal) : ", i)
	// 	condition.Signal()                    // 모든 고루틴 생성 후, 한 개 씩 깨움

	// 	mutex.Unlock()
	// }

	// ----
	// 대기 중인 고루틴을 한 번에 깨우는 방법
	mutex.Lock()
	
	fmt.Println("Wake Groutine(Brodcast)")
	condition.Broadcast()

	mutex.Unlock()
	// ----

	time.Sleep(2 * time.Second)
}
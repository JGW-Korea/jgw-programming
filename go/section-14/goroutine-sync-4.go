package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// RWMutex -> 쓰기 Lock -> 쓰기 시도 중에는 다른 곳에서 이전 값을 읽으면 X, 읽기 락, 쓰기 락 전부 방어
	// RMutex  -> 읽기 Lock -> 읽기 시도 중에 값이 변경 방지 즉, 쓰기 락 방어

	data := 0
	mutex := new(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.Lock()     // 쓰기 뮤텍스(Write Mutex) 잠금
			
			data += 1
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Millisecond)
			
			mutex.Unlock()   // 쓰기 뮤텍스(Write Mutex) 잠금 해제
		}
	}()
	
	go func() {
		for i := 1; i <= 10; i++ {
			mutex.RLock()    // 읽기 뮤텍스(Read Mutex) 잠금

			fmt.Println("Read1 : ", data)
			time.Sleep(1 * time.Second)

			mutex.RUnlock()  // 읽기 뮤텍스(Read Mutex) 잠금 해제
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			mutex.RLock()    // 읽기 뮤텍스(Read Mutex) 잠금

			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)

			mutex.RUnlock()  // 읽기 뮤텍스(Read Mutex) 잠금 해제
		}
	}()

	time.Sleep(10 * time.Second)
}
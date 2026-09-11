package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 원자성이 보장되지 않는 함수
func notAtomic() {
	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt += 1
			wg.Done()
		}(i)
	}
	
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt -= 1
			wg.Done()
		}(i)
	}

	wg.Wait()                             
	fmt.Println("[Not Atomic] Watit Group End Cnt?  \t >>>>", cnt)  // 대기가 풀린 이후 실행된다.
}

// 원자성이 보장되는 함수
func trueAtomic() {
	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			atomic.AddInt64(&cnt, 1)
			wg.Done()
		}(i)
	}
	
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			atomic.AddInt64(&cnt, -1)
			wg.Done()
		}(i)
	}

	wg.Wait()                             
	finalCnt := atomic.LoadInt64(&cnt)

	fmt.Println("[True Atomic] Watit Group End Cnt? \t >>>>", cnt)             // 대기가 풀린 이후 실행된다.
	fmt.Println("[True Atomic] Watit Group End Final Cnt? >>>>", finalCnt)  // 대기가 풀린 이후 실행된다. (추천)
}

func main() {
	// 고루틴 동기화 고급
	// 원자성 사용 -> 더 이상 쪼갤 수 없는 하나의 완전한 작업 단위로 보고, 전부 성공하거나 아예 실행되지 않아야 하는 성질
	// 모든 조작이 완료 될 때까지 다른 프로세스 개입 불가
	// sync/atomic에서 원자적 연산자 제공


	notAtomic()   // 원자성이 보장되지 않는 함수 호출
	trueAtomic()  // 원자성이 보장되는 함수 호출
}
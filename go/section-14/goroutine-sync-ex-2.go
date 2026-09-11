package main

import (
	"fmt"
	"sync"
)

func main() {
	// 고루틴 동기화 고급
	// 대기 그룹
	// 실행 흐름 변경 (고루틴이 종료될 때까지 대기 기능)
	// WaitGroup : Add(고루틴 추가), Done(작업 종료 알림), Wati(고루틴 종료시 까지 대기)
	// - Add로 추가 된 고루틴 개수와 Done으로 종료되는 알림 횟수가 같아야 정확하게 동작한다. (Add == Done)
	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("Wait Group : ", n)
			wg.Done()
		}(i)
	}

	wg.Wait()                        // 생성된 모든 고루틴이 끝날 때까지 대기한다.
	fmt.Println("Watit Group End!")  // 대기가 풀린 이후 실행된다.
}
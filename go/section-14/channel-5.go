package main

import "fmt"

func main() {
	ch := make(chan bool)
	
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		
		close(ch) // 5회 채널에 값 전송 후 채널 해제
	}()
	
	// 채널 + range 조합은 채널이 해제될 때까지 값을 수신받는다.
	for i := range ch {
		fmt.Println("ex1 :", i)
	}
}
package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool, 2) // 버퍼 사용
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<- ch
		fmt.Println("Main :", i)
	}
}
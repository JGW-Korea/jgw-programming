package main

import "fmt"

// 반복자 함수 기본 정의 방법
func Count(yield func(int) bool) {
	for n := range 5 {
		if !yield(n) {
			return
		}
	}
}

// for range에서 호출
func main() {
	// temp := Count(func(i int) bool {
	// 	fmt.Println(i)
	// 	return true
	// })

	// 반복자 함수를 직접 호출
    Count(func(n int) bool {
        fmt.Println(n)
        return true
    })

	var result int = 0
	
	for cnt := range Count {
		result += cnt
	}
	
	fmt.Println(result)
}
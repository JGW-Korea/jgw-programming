package main

import "fmt"

func main() {
	// append() 사용 방법 #1. 길이와 용량이 지정되지 않은 슬라이스
	var slice1 []int
	
	fmt.Println("before slice1:", slice1)
	for i := 0; i < 5; i++ {
		slice1 = append(slice1, i + 1)
	}
	fmt.Println("after slice1:", slice1) // [1, 2, 3, 4, 5]
	
	// append() 사용 방법 #2. 길이와 용량이 지정되어 있는 슬라이스
	var slice2 []int = make([]int, 5, 10)
	
	fmt.Println("before slice2:", slice2)
	for i := 0; i < 5; i++ {
		slice2 = append(slice2, i + 1)
	}
	fmt.Println("after slice2:", slice2) // [0, 0, 0, 0, 0, 1, 2, 3, 4, 5]
}
package main

import "fmt"

// 배열 사용 방법 #1. 기본 선언 방법
func array() {
	arr1 := [3]int{1, 2, 3};
	arr2 := [3]int{4, 5, 6}

	// arr2[2] *= 10

	fmt.Printf("%v / memory(%p)\n", arr1, &arr1);
	fmt.Printf("%v / memory(%p)\n", arr2, &arr2);

	if arr1 == arr2 {
		fmt.Println("oning")
	}
}
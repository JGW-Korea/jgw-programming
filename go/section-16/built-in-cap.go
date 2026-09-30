package main

import (
	"fmt"
)

func useLenInArray() {
	arr := [5]int{1, 2, 3, 4, 5}
	
	fmt.Println("length of array:", len(arr))
	fmt.Println("capacity of array:", cap(arr))
}

func useLenInSlice() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("length of slice:", len(slice1))
	fmt.Println("capacity of slice:", cap(slice1))
	
	slice2 := make([]int, 5, 10)
	fmt.Println("length of slice:", len(slice2))
	fmt.Println("capacity of slice:", cap(slice2))
}

func useLenPointer() {
	// 💡 배열은 역참조를 하지 않아도 가능 -> Go 자체에서 역참조하여 길이를 접근
	arr := [5]int{1, 2, 3, 4, 5}
	arrPtr := &arr
	fmt.Println("Length of Pointer to Array:", len(arrPtr), " / cap:", cap(arrPtr))
	
	// ❌ 슬라이스는 역참조를 하지 않으면 컴파일 오류 -> Go 자체에서 *[]T는 len 허용 대상이 아니도록 설계
	slice := make([]int, 5, 10)
	slicePtr := &slice	
	fmt.Println("Length of Pointer to Slice:", len(*slicePtr), " / cap:", cap(*slicePtr))
}

func main() {
	useLenInArray()
	useLenInSlice()
	useLenPointer()
}
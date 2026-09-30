package main

import (
	"fmt"
	"unicode/utf8"
)

func useLenInString() {
	// 문자열 len() 사용 예시
	var str1 string = `c:\go_study\src\`
	var str2 string = "\ud55c\uae00"
	var str3 string = "안녕하세요."
		
	fmt.Printf("value: %v / using byte: %d / string length: %d\n", str1, len(str1), utf8.RuneCountInString(str1))
	fmt.Printf("value: %v / using byte: %d / string length: %d\n", str2, len(str2), utf8.RuneCountInString(str2))
	fmt.Printf("value: %v / using byte: %d / string length: %d\n", str3, len(str3), utf8.RuneCountInString(str3))
}

func useLenInArray() {
	// 배열 len() 사용 예시
	arr := [5]int{1, 2, 3, 4, 5}
		
	for i := 0; i < len(arr); i++ {
		fmt.Printf("idx=%d, value=%d\n", i, arr[i])
	}
}

func useLenInSlice() {
	// 슬라이스 len() 사용 예시 #1. nil 슬라이스 선언 시 길이는 0을 반환
	var slice1 []int // nil 슬라이스
	fmt.Printf("length of slice1: %d\n", len(slice1));            // length of slice1: 0
	fmt.Printf("length=%d / cap=%d\n", len(slice1), cap(slice1)); // length=0 / cap=0
	
	slice1 = append(slice1, 1, 2, 3, 4)
	fmt.Printf("length of slice1: %d\n", len(slice1));            // length of slice1: 0
	fmt.Printf("length=%d / cap=%d\n", len(slice1), cap(slice1)); // length=0 / cap=0
	
	// 슬라이스 len() 사용 예시 #2. make()를 통해 선언하면 길이는 전달한 길이를 반환
	slice2 := make([]int, 5, 10)
	
	// 예시를 위한 조건문 -> 권장 방법 for idx, value := range arr { /* ... */ }
	for i := 0; i < len(slice2); i++ {
		slice2[i] = (i + 1) * 10
	}
	
	fmt.Printf("slice2: %v\n", slice2)                           // slice2: [10 20 30 40 50]
	fmt.Printf("length=%d / cap=%d\n", len(slice2), cap(slice2)) // length=5 / cap=10
}

func useLenInMap() {
	// 맵 len() 사용 예시 #1. 초기화 때 전달한 값만큼 길이가 확보
	var map1 map[string]string = map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
		"key5": "value5",
	}
	fmt.Printf("length of map1: %d\n", len(map1)); // length of map1: 5
	
	// 슬라이스 len() 사용 예시 #2. make()를 통해 empty map을 만들었기 때문에 길이는 0이다.
	map2 := make(map[string]string)
	
	// 예시를 위한 조건문 -> 권장 방법 for key, value := range map2
	for i := 0; i < len(map2); i++ {
		key, value := fmt.Sprintf("key%d", i + 1), fmt.Sprintf("value%d", i + 1)
		map2[key] = value
	}
	
	fmt.Println("map2:", map2) // map2: []
}

func useLenPointer() {
	// 💡 배열은 역참조를 하지 않아도 가능 -> Go 자체에서 역참조하여 길이를 접근
	arr := [5]int{1, 2, 3, 4, 5}
	arrPtr := &arr
	fmt.Println("Length of Pointer to Array:", len(arrPtr), " / cap:", cap(arrPtr))
	
	// ❌ 슬라이스는 역참조를 하지 않으면 컴파일 오류 -> Go 자체에서 *[]T는 len 허용 대상이 아니도록 설계
	slice := make([]int, 5, 10)
	slicePtr := &slice	
	fmt.Println("Length of Pointer to Slice:", len(slicePtr), " / cap:", cap(slicePtr))
}

func main() {
	useLenInString()
	useLenInArray()
	useLenInSlice()
	useLenInMap()
	useLenPointer()
}
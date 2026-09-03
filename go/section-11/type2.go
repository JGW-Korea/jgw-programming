package main

func main() {
	// nil 예제 #1. nil을 가지거나, 할당받을 수 있는 경우
	var map1 map[string]int
	var map2 map[string]int = nil
	var ptrValue *int = nil

	// nil 예제 #2. nil을 가질 수 없는 경우 (컴파일 에러)
	var value1 int = nil
	value2 := nil
}
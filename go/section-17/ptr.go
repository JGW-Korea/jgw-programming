package main

import (
	"fmt"
	"unsafe"
)

func main() {
    var a int = 10

	var ptr *int = &a
	var addr uintptr = uintptr(unsafe.Pointer(&a))
	
	fmt.Println(*ptr)
	
	fmt.Println(ptr)
	fmt.Println(addr)

	var targetString string = "Hello, Go Programming"

	byteSlice  := []byte(targetString)
	runeSlice  := []rune(targetString)
	int32Slice := []int32(targetString)

	fmt.Printf("value: %s, len: %d\n", byteSlice,  len(byteSlice))  // 문자열 정상 출력
	fmt.Printf("value: %s, len: %d\n", runeSlice,  len(runeSlice))  // 문자열 정상 출력 X
	fmt.Printf("value: %s, len: %d\n", int32Slice, len(runeSlice))  // 문자열 정상 출력 X

}
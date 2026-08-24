package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str1 string = `c:\go_study\src\`
	var str2 string = "\ud55c\uae00"
	var str3 string = "안녕하세요."
	
	fmt.Printf("value: %v / using byte: %d / string length: %d\n", str1, len(str1), utf8.RuneCountInString(str1))
	fmt.Printf("value: %v / using byte: %d / string length: %d\n", str2, len(str2), utf8.RuneCountInString(str2))
	fmt.Printf("value: %v / using byte: %d / string length: %d\n", str3, len(str3), utf8.RuneCountInString(str3))

	fmt.Printf("%c\n", str1[0])

	// var char1 rune = 'd'
	// var char2 int = 'd'
	// var char3 byte = 'd'
	// var char4 int32 = 'd'
	// var char5 int16 = 'd'
	// var char6 float32 = 'd'
	// var char7 float64 = 'd'

	// fmt.Printf("value: %c\n", char1)
	// fmt.Printf("value: %c\n", char2)
	// fmt.Printf("value: %c\n", char3)
	// fmt.Printf("value: %c\n", char4)
	// fmt.Printf("value: %c\n", char5)
	// fmt.Printf("value: %c\n", char6)
	// fmt.Printf("value: %c\n", char7)

	var char1 int = 'h'
	var char2 byte = 'e'
	var char3 rune = 'l'
	var char4 int32 = 'l'
	var char5 int64 = 'o'

	fmt.Printf("%c %c %c %c %c\n", char1, char2, char3, char4, char5);

	var str string = "world"
	var strIndex0 byte = str[0]
	var strIndex1 byte = str[1]
	var strIndex2 byte = str[2]
	var strIndex3 byte = str[3]
	var strIndex4 byte = str[4]

	fmt.Printf("%c %c %c %c %c\n", strIndex0, strIndex1, strIndex2, strIndex3, strIndex4);

	var targetString string = "Hello, Go Programming"
	
	byteSlice  := []byte(targetString)
	runeSlice  := []rune(targetString)
	int32BitSlice := []int32(targetString)

	fmt.Printf("value: %s, len: %d\n", byteSlice, len(byteSlice))  // 문자열 정상 출력
	fmt.Printf("value: %v, len: %d\n", runeSlice, len(runeSlice))  // 문자열 정상 출력 X
	fmt.Printf("value: %v, len: %d\n", int32BitSlice, len(runeSlice)) // 문자열 정상 출력 X

	var byte1 byte = 'a'                    // 단일 문자 표현 가능
	var byte2 []byte = []byte("Hello")      // 문자열 표현 가능
	var byte3 []byte = []byte("가")          // 3Byte 차지하는 한글 단일 문자 표현 가능
	var byte4 []byte = []byte("안녕하세요~")   // 3Byte 차지하는 한글 단일 문자열 표현 가능
	var byte5 []byte = []byte("🤔")         // 4Byte 차지하는 이모지 표현 가능

	fmt.Printf("%v, %c\n", byte1, byte1) // a, len(1)
	fmt.Printf("%v, %s, len(%d)\n", byte2, byte2, len(byte2)) // Hello, len(5)
	fmt.Printf("%v, %s, len(%d)\n", byte3, byte3, len(byte3)) // 가, len(1)
	fmt.Printf("%v, %s, len(%d)\n", byte4, byte4, len(byte4)) // 안녕하세요, len(5)
	fmt.Printf("%v, %s, len(%d)\n", byte5, byte5, len(byte5)) // 🤔, len(1)

	var int32Byte []int32 = []int32("Hi")
	var runeSlic2e []rune = []rune("Nice")

	fmt.Printf("%v, len(%d)\n", int32Byte, len(int32Byte))
	fmt.Printf("%v, len(%d)\n", runeSlic2e, len(runeSlic2e))
}


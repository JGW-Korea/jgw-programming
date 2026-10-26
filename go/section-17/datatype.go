package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

// any #2. 매개변수
func add(x int, y any) (int, error) {
	switch value := y.(type) {
		case int: {
			return x + value, nil
		}
		default: {
			return -1, fmt.Errorf("잘못된 자료형...")
		}
	}
}

func main() {
	// any
	var (
		integer  any = 10
		floating any = 10.4
		boolean  any = false
		str      any = "Hello"
	)

	fmt.Println(integer, reflect.TypeOf(integer))    // 10 int
	fmt.Println(floating, reflect.TypeOf(floating))  // 10.4 float64
	fmt.Println(boolean, reflect.TypeOf(boolean))    // false bool
	fmt.Println(str, reflect.TypeOf(str))            // Hello string

	value, err := add(10, 40);

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)

	// rune
	// ⚠️ []byte 길이 출력
	var str1 []byte = []byte("Hello, Go Programming")
	var str2 string = "만나서 반가워요!!"

	fmt.Printf("value=%s, len=%d\n", str1, len(str1)) // value=..., len=21
	fmt.Printf("value=%s, len=%d\n", str2, len(str2)) // value=..., len=24

	// 정상적으로 길이를 출력하는 방법 #1. rune으로 형변환
	var formatRuneStr1 []rune = []rune(string(str1))
	var formatRuneStr2 []rune = []rune(str2)

	fmt.Println("value=%s, len=%d\n", formatRuneStr1, len(formatRuneStr1)) // value=..., len=21
	fmt.Println("value=%s, len=%d\n", formatRuneStr2, len(formatRuneStr2)) // value=..., len=10

	// 정상적으로 길이를 출력하는 방법 #2. utf8.RuneCountInString()
	fmt.Println("value=%s, len=%d\n", formatRuneStr1, utf8.RuneCountInString(string(str1))) // value=..., len=21
	fmt.Println("value=%s, len=%d\n", formatRuneStr2, utf8.RuneCountInString(str2))         // value=..., len=10

	// 💡 문자열 길이 출력
	var str3 []rune = []rune("What is the length of this string?")
	var str4 []rune = []rune("한글도 상관없지요~~")

	fmt.Printf("value=%s, len=%d\n", str3, len(str3))
	fmt.Printf("value=%s, len=%d\n", str4, len(str4))

	// complex
	var c1 complex64 = 1 + 2i
	var c2 complex128 = complex(3, 4)
}
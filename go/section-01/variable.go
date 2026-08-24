package main

import (
	"fmt"
	"reflect"
)

type MyInt = int

func main() {
	var myInt MyInt = 10
	// 단축 변수 선언 시 사용 방법 -> myInt := MyInt(10)

	var compareIntVariable = 20

	if reflect.TypeOf(myInt) == reflect.TypeOf(compareIntVariable) {
		fmt.Println("동일한 자료형")
	} else {
		fmt.Println("서로 다른 자료형")
	}
}
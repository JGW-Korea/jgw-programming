package main

import (
	"fmt"
)

type Union interface {
	int | string
}

// 타입 스위치(Type Switch) #1. 기본 사용 방법
func narrowing(value any) {
	switch value := value.(type) {
		case int: {     // value가 int 자료형으로 좁혀짐
			fmt.Println(value, "type is int")
		}
		case string: {  // value가 string 자료형으로 좁혀짐
			fmt.Println(value, "type is string")
		}
		default: {
			fmt.Println("?")
		}
	}
}

// 타입 스위치(Type Switch) #2. 제네릭
func narrowingGeneric[T Union](value T) {
	switch value := any(value).(type) {
		case int: {     // value가 int 자료형으로 좁혀짐
				fmt.Println(value, "type is int")
		}
		case string: {  // value가 string 자료형으로 좁혀짐
				fmt.Println(value, "type is string")
		}
		default: {
				fmt.Println("?")
		}
	}
}

func main() {
	narrowing(10)
	narrowing("10")
}

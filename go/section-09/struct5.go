package main

import "fmt"

// 중첩 구조체(Nested Struct)
type Car struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct {
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {
	car := Car{
		"520d",
		"silver",
		"bmw",
		spec{
			length: 4000,
			height: 1000,
			width:  2000,
		},
		// 또는 -> spec{4000, 1000, 2000}
	}
	
	// 중첩 구조체 접근 #1. 상위 구조체
	fmt.Println("Parent Struct:", car.name)
	fmt.Println("Parent Struct:", car.color)
	fmt.Println("Parent Struct:", car.company)
	fmt.Printf("Parent Struct: %#v\n", car.detail)
	
	// 중첩 구조체 접근 #2. 하위 구조체
	fmt.Println("length", car.detail.length)
	fmt.Println("height", car.detail.height)
	fmt.Println("width", car.detail.width)
}
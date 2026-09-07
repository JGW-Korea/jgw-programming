package main

import "fmt"

func main() {
	var x int32 = 4921
	var y byte  = byte(x)

	fmt.Println("x:", x)
	fmt.Println("y:", y)  // ⚠️ 데이터 유실
}
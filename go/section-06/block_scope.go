// 블록 스코프(Block Scope) 예시
package main

import "fmt"

func main() {
	{
		x := 10
		y := 20
		
		fmt.Printf("%d", x + y);
	}
}
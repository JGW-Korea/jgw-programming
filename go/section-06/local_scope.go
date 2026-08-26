// 지역 스코프(Local Scope) 예시
package main

import "fmt"

var a, b, c = 1, 2, 3

func main() {
	{
		var d int = 4
		
		fmt.Printf("%d / %d / %d / %d", a, b, c, d);
	}
	
	fmt.Printf("%d / %d / %d / %d", a, b, c, d);
}
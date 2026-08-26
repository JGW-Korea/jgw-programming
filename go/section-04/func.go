package main

import "fmt"

func multiply(x, y int) int {
	return x * y
}

func sum(x, y int) int {
	return x + y
}

func main() {
	functions := []func(int, int) int {
		multiply,
		sum,
	}

	var variableFunc func(int, int) int = multiply


	fmt.Println("value:", functions[0](10, 20))
	fmt.Println("value:", functions[1](10, 20))
	fmt.Println("value:", variableFunc(10, 20))

	// 함수 타입 #3. 맵의 요소로 할당
	mapFuncs := map[string]func(int, int) int {
		"mul_func": multiply,
		"sum_func": sum,
	}

	for _, function := range mapFuncs {
		fmt.Println("map func value:", function(10, 20))
	}

	callee()
}
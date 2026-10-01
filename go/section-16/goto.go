package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var result string = "values: ";
	
	if random := rand.Intn(10) + 1; random % 2 == 0 {
		goto isEven
	} else {
		goto isOdd
	}
	
	// random이 짝수인 경우, isOdd 레이블 코드는 실행되지 않음
	isOdd:
	for i := range 10 {
		if i % 2 == 1 {
			result += strconv.Itoa(i) + " "
		}
	}
	
	fmt.Println(result)
	return
	
	// random이 홀수인 경우, isEven 레이블 코드는 실행되지 않음
	isEven:
	for i := range 10 {
		if i % 2 == 0 {
			result += strconv.Itoa(i) + " "
		}
	}
	
	fmt.Println(result)
	return
}
package main

import "fmt"

var slice []int = []int{1, 2, 3, 4, 5}
var user map[string]string = map[string]string{
	"name": "userA",
	"id":   "userId",
	"pwd":  "userPwd",
}

func main() {
	// range 사용 예시 #1. 일반 반복
	for i := range 10 {
		fmt.Println(i)
	}
	
	// range 사용 예시 #2. 반복문 또는 슬라이스
	for idx, value := range slice {
		fmt.Printf(
			"idx: %d, value: %d -> slice[%d]=%d\n",
			idx,
			value,
			idx,
			value,
		)
	}
	
	// range 사용 예시 #3. 맵(Map)
	for key, value := range user {
		fmt.Printf(
			"key: %s, value: %s -> slice[%s]=%s\n",
			key,
			value,
			key,
			value,
		)
	}

	for count := 5; range count {
		fmt.Println(count)
	}
}
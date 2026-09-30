package main

import "fmt"

func main() {
	targetMap := map[string]string{
		"name": "Vito",
		"location": "New York",
	}

	delete(targetMap, "name")       // targetMap의 name 프로퍼티를 제거 
	delete(targetMap, "undefined")  // 존재하지 않는 키를 제거해도 오류가 발생 X

	fmt.Println(targetMap)
}
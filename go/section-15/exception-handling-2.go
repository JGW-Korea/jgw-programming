package main

import (
	"fmt"
	"log"
)

func notZero(num int) (string, error) {
	var s string
	
	if num != 0 {
		s = fmt.Sprint("Hello Golang : ", num)
		return s, nil
	}

	return "", fmt.Errorf("%d를 입력했습니다. 에러 발생!", num)
}

func main() {
	a, err1 := notZero(1)

	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("ex1 : ", a)

	b, err2 := notZero(0)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("ex1 : ", b)

	fmt.Println("End Error Handling!")
}
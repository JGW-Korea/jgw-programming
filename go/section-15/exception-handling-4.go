package main

import (
	"errors"
	"fmt"
	"log"
)

func notZero(num int) (string, error) {
	var s string
	
	if num != 0 {
		s = fmt.Sprint("Hello Golang : ", num)
		return s, nil
	}

	return "", errors.New("0를 입력했습니다. 에러 발생!")
}

func main() {
	a, err1 := notZero(1)

	if err1 != nil {
		// log.Fatal(err1)
		log.Fatal(err1.Error())
	}

	fmt.Println("ex1 : ", a)

	b, err2 := notZero(0)

	if err2 != nil {
		// log.Fatal(err2)
		log.Fatal(err2.Error())
	}

	fmt.Println("ex1 : ", b)

	fmt.Println("End Error Handling!")
}
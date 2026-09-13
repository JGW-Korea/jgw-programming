package main

import (
	"errors"
	"fmt"
)

func main() {
	var err1 error = errors.New("Error occurred - 1")
	err2 := errors.New("Error occurred - 2")

	fmt.Println("error1 : ", err1)
	fmt.Println("error1 : ", err1.Error())

	fmt.Println("error2 : ", err2)
	fmt.Println("error2 : ", err2.Error())
}
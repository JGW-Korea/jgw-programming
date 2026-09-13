package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("unnamedfile") // 강제 예외 발생

	if err != nil {
		// log.Fatal(err.Error()) // 예외 처리 방법 #1
		log.Fatal(err) // 예외 처리 방법 #2
	}

	fmt.Println("ex1 : ", file.Name())
}
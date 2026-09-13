// Panic & Recover

package main

import (
	"fmt"
	"os"
)

func fileOpen(fileName string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error :", r)
		}
	}()

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("ex 1 : ", file.Name())
	}

	defer file.Close()
}

func main() {
	fileOpen("undefined.txt")
	fmt.Println("End Main")
}
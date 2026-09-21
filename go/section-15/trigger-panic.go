package main

import (
	"fmt"
	"log"
)

func step3() {
	panic("Trigger Exception!!")
	fmt.Println("Hahahah!!")

}

func step2() {
	step3()
}

func step1() {
	step2()
}

func main() {
	// recove()가 없으면 프로그램 중단
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	
	log.Panic("sddddd")

	step1()

	fmt.Println("Hahahah")
}
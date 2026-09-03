package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

func (dog Dog) bite() {
	fmt.Println(dog.name, "bites!")
}

// 동물의 행동을 정의한 인터페이스
type Behavivor interface {
	bite()
}

func main() {
	dog1 := Dog{"poll", 10}

	var inter1 Behavivor
	inter1 = dog1
	inter1.bite()

	dog1.bite()
}
package main

import "fmt"

// Go 인터페이스 - 암시적 인터페이스 구현(Implicit Interface Imlementation)
type Behavior interface {
	bite()
}

// 구조체 선언 #1. 인터페이스 구조를 만족하는 구조체
type Dog struct {
	name   string
	weight int
}

func (dog Dog) bite() {
	fmt.Println(dog.name, "bites!")
}

// ---

// 구조체 선언 #2. 인터페이스 구조를 만족하지 않는 구조체
type Cat struct {
	name   string
	weight int
}

func main() {
	var dog1 Dog = Dog{"poll", 10}
	var inter1 Behavior = dog1
	inter1.bite()
	
	dog2 := Dog{"marry", 12}
	inter2 := Behavior(dog2)
	inter2.bite()
	
	inters := []Behavior{dog1, dog2}
	for idx, inter := range inters {
		inters[idx].bite()
		inter.bite()
	}
	
	// Cat은 Behavior 인터페이스 구조를 만족하지 못하여 컴파일 에러가 발생
	// inter2 := Behavior(Cat{"gogo", 5})
}
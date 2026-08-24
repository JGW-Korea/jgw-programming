package main

import "fmt"

type User struct {
	Name string
	Age int
}

func pointer() {
	userA := User{ Name: "A", Age: 28 }
	userB := userA

	userB.Name = "B"

	fmt.Println(userA)  // {A 28} 출력
	fmt.Println(userB)  // {B 28} 출력

	var userC *User = &userB

	userC.Name = "C"

	fmt.Println(userB)  // {B 28} 출력
	fmt.Println(userC)  // {B 28} 출력

	// var targetVariable int = 89
	// var pointerVariable2 *int = &targetVariable
	// var pointerVariable3 = &targetVariable
	// poiinterVariable4 := &targetVariable

	targetVariableddd := 10

	pointerVariableddd := &targetVariableddd

	fmt.Println(pointerVariableddd)
	fmt.Println(&pointerVariableddd)
	fmt.Println(*pointerVariableddd)

	*pointerVariableddd = 102



}